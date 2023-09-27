package controllers

import (
	"net/http"
	"os"
	"project-api/database"
	"project-api/models"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

// Register
func Signup(c *gin.Context){
	// Get the data off req body
	var body struct{
		Username string `binding:"required"`
		Email string 	`binding:"required"`
		Password string `binding:"required,min=6"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}

	// check email with govalidator
	if govalidator.IsEmail(body.Email) {
		// check if users have been created with the same email
		var user models.User
		database.DB.First(&user, "email = ?", body.Email)
		if user.Email == body.Email {
			c.JSON(http.StatusBadRequest, gin.H{"error":"The user already exists"})
		} else {
			// Create the user
			user := models.User{
				Username: body.Username,
				Email: body.Email, 
				Password: string(hash),
			}
			result := database.DB.Create(&user)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error":"Failed to create user"})
				return
			}

			// Respond
			c.JSON(http.StatusOK, gin.H{"message" : "User created successfully"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email tidak valid"})
	}
}

// Login
func Login(c *gin.Context){
	// Get the  email and pass off req body
	var body struct{
		Email string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	// Look up requested user
	var user models.User
	database.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error":"User not found, create a user"})
		return
	}

	// Compare sent in pass with save user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid Email or Password"})
		return
	}

	//generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Failed to create token"})
		return
	}

	//send it back 
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message" : "Login successful"})
}

// Validate 
func Validate(c *gin.Context){
	// Retrieves the logged in user ID
	userID := c.MustGet("userID").(uint)
	// Validate who the logged in user is
	c.JSON(http.StatusOK, gin.H{
		"User who logged is " :  userID})
}

// Show All User
func ShowAllUser(c *gin.Context){
	// Get  all data from model
	var user []models.User
	database.DB.Preload("Photo").Find(&user)

	// Respomd
	c.JSON(200, gin.H{
		"message" : user,
	})
}

// Update
func UpdateUser(c *gin.Context){
	// Retrieves the logged in user ID
	userID := c.MustGet("userID").(uint)
	userIDStr := c.Param("id") 
	userIDFromParam, err := strconv.ParseUint(userIDStr, 10,32) // get param id and change the type
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid user ID"})
	}

	// Check if userID and param id have the same value
	if userID != uint(userIDFromParam) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : "Unauthorized"})
			return
	}

	// Check user data that matches the ID
	var user models.User
	if err := database.DB.First(&user, userIDFromParam).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error" : "User Not Found"})
		return
	}

	// Get the data off req body
	type body struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var  updateUser body 
	if err := c.BindJSON(&updateUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Check the data condition of the req body
	if updateUser.Email != "" { 
        user.Email = updateUser.Email} // Check email req
    if updateUser.Username != "" {
        user.Username = updateUser.Username} // Check username req
	if len(updateUser.Password) > 0 && len(updateUser.Password) >= 6 { // Check pass req
		hash, err := bcrypt.GenerateFromPassword([]byte(updateUser.Password), 10)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to hash password"})
			return
		}
		user.Password = string(hash)
	}

	// Save new data to database
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Failed to update user"})
		return
	}

	// Respond and show new user data
	userShow := database.DB.Where("id = ?", userID).Preload("Photo").Find(&user)
	c.JSON(http.StatusOK, gin.H{
		"message" : "User Update successfully",
		"value" : userShow,
	})
}

// Delete
func DeleteUser(c *gin.Context){
	// Retrieves the logged in user ID
	userID := c.MustGet("userID").(uint)
	userIDStr := c.Param("id")
	userIDFromParam, err := strconv.ParseUint(userIDStr, 10,32) // get param id and change the type
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid user ID"})
	}

	// Check if userID and param id have the same value
	if userID != uint(userIDFromParam) {
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "Unauthorized"})
		return
	}
	
	// Delete User
	database.DB.Select(clause.Associations).Delete(&models.User{}, userID)
	database.DB.Where("user_id = ?", userID).Delete(&models.Photo{})

	// Respond With It
	c.JSON(200, gin.H{
		"message" : "User Already Delete",
	})
}