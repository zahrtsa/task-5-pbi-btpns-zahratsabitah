package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Token & Validate
func RequireAuth(c *gin.Context){
	// Get the cookie off req
	tokenString, err := c.Cookie("Authorization")
	if err != nil{
		c.AbortWithStatus(http.StatusUnauthorized)
		c.Redirect(http.StatusFound, "/users/login") //if method is GET, for validate 
	}

	// Decode/validate it
	// Parse takes the token string and a function for looking up the key. The latter is especially
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil // GET secret key from env file
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64){
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user ID with token sub
		userID := uint(claims["sub"].(float64)) 
		// Set  variable userID as userID, so in the controller we can get userID who is logged in
		c.Set("userID", userID) 

		// Continue
		c.Next()
		
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}