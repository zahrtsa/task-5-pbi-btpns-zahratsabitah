package controllers

import (
	"net/http"
	"project-api/database"
	"project-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Add Photos
func PostPhotos(c *gin.Context){
	// Retrieves the logged in user ID
	userID := c.MustGet("userID").(uint)

	// Get data from req 
	var photo models.Photo
    if err := c.BindJSON(&photo); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

	// Get user data by userID
	var user models.User
    if err := database.DB.Where("id = ?", userID).Preload("Photo").First(&user).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to get user data"})
        return
    }
	
	// Check if the user already have photo or not
    if user.Photo.ID == 0 {
        photo.UserID = userID
        database.DB.Create(&photo) // create photo for user
        user.Photo = photo

		// Respond and see user photo 
		photoShow := database.DB.First(&photo, userID)
		c.JSON(http.StatusOK, gin.H{
			"message" : "User Photo Create successfully",
			"value" : photoShow})
	} else {
		// If user already have photo
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Already have photo, update to change",
		})
	}
}

// See  All Photos
func ShowPhotos(c *gin.Context){
	// Get All Show  Photos
	var photo []models.Photo
	database.DB.Find(&photo)

	// Resond
	c.JSON(http.StatusOK, gin.H{
		"message" : photo,
	})
}

// Update Photos
func UpdatePhotos(c *gin.Context){
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

	// Get photo req
	var requestData map[string]interface{}
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Get photo by user id who logged in
	var user models.User
	if err := database.DB.Where("id = ?", userIDFromParam).Preload("Photo").First(&user).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error" : "User Not Found"})
		return
	}

	// Check id from photo who relation with user
	if  user.Photo.ID != 0 {
		if err := database.DB.Model(&user.Photo).Updates(requestData).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error" : "Failed to update user photo"})
			return
		}

		// Respond & show photo data
		var photo models.Photo
		photoShow := database.DB.First(&photo, userID)
			c.JSON(http.StatusOK, gin.H{
				"message" : "User Photo Update successfully",
				"value" : photoShow,
			})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "ID does not match",
		})
	}
}

// Delete Photos
func DeletePhotos(c *gin.Context){
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

	// Delete User
	database.DB.Delete(&models.Photo{}, userID)

	// Respond With It
	c.JSON(200, gin.H{
		"message" : "Photo has been successfully deleted, post a new photo"})
}