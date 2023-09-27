package router

import (
	"project-api/controllers"
	"project-api/middleware"

	"github.com/gin-gonic/gin"
)

// Route to call all method 
// call the method in controller by route
func Route(){
	r:= gin.Default()

	// For user get all method from userControllers
	user := r.Group("/users") 
	{
		user.POST("/register", controllers.Signup) // Create User
		user.POST("login", controllers.Login) // Login as user & get the token
		user.GET("/showAll", controllers.ShowAllUser) // Show all successfully created users
		user.GET("/validate", middleware.RequireAuth, controllers.Validate) // Validate if user already login
		user.PUT("/update/:id", middleware.RequireAuth, controllers.UpdateUser) // Update  user & get id by param id
		user.DELETE("/delete/:id",middleware.RequireAuth, controllers.DeleteUser) // Update  user & get id by param id
	}

	// For photo get all method from photoControllers
	photo := r.Group("/photo") 
	{
		photo.POST("/post", middleware.RequireAuth, controllers.PostPhotos) // Create/post photo for user
		photo.GET("/show", middleware.RequireAuth, controllers.ShowPhotos) // Show all photo 
		photo.PUT("/update/:id", middleware.RequireAuth, controllers.UpdatePhotos) // Update photo for user & get id by param id
		photo.DELETE("/delete/:id",middleware.RequireAuth, controllers.DeletePhotos) // Delete photo for user & get id by param id
	}

	r.Run()
}