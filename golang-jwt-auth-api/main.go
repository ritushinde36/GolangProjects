package main

import (
	"jwt-auth-sql/controllers"
	"jwt-auth-sql/middleware"
	"jwt-auth-sql/pkg"

	"github.com/gin-gonic/gin"
)

func init() {
	pkg.LoadEnvVariables()
	pkg.ConnectToDB()
	pkg.CreateTable()
}

func main() {
	//set up the router
	router := gin.Default()
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.GET("/users", middleware.AuthenticateUser, controllers.GetAllUsers)
	router.GET("/user", middleware.AuthenticateUser, controllers.GetUserDataByID)
	router.PUT("/user", middleware.AuthenticateUser, controllers.UpdateUser)
	router.DELETE("/user", middleware.AuthenticateUser, controllers.DeleteUser)
	router.Run()

}
