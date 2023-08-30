package main

import (
	"gin-gorm-rest/config"
	"gin-gorm-rest/controller"
	"gin-gorm-rest/middleware"
	"gin-gorm-rest/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// router := gin.New()
	router := gin.Default() //attaches Logger and Recovery middleware to the router

	config.Connect()

	// router.GET("/",func(c *gin.Context) {
	// 	c.String(200,"Hello World")
	// })

	//Middleware
	// router.Use(middleware.Authenticate)

	//Method -1
	routes.UserRoute(router)

	//Method -2
	userRouter := router.Group("/users", middleware.Authenticate(), middleware.AddHeader())
	{
		userRouter.GET("/", controller.GetUsers)
		userRouter.POST("/", controller.CreateUser)
		userRouter.PATCH("/:id", controller.UpdateUser)
		userRouter.DELETE("/:id", controller.DeleteUser)
	}

	router.Run(":8080")
}
