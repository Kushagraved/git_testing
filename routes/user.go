package routes

import (
	"gin-gorm-rest/controller"
	"gin-gorm-rest/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	// router.GET("/",func(c *gin.Context) {
	// 	c.String(200,"Hello World")
	// })

	router.GET("/", middleware.Authenticate(), controller.CreateUser)
	router.POST("/", controller.CreateUser)
	router.PATCH("/:id", controller.UpdateUser)
	router.DELETE("/:id", controller.DeleteUser)

}
