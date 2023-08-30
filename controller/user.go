package controller

import (
	"fmt"
	"gin-gorm-rest/api"
	"gin-gorm-rest/config"
	"gin-gorm-rest/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    users,
	})
}

func CreateUser(c *gin.Context) {
	//Request Validation
	var CreateUserReq api.CreateUserRequest

	if err := c.ShouldBindJSON(&CreateUserReq); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//Create User
	newUser := models.User{
		Name:     CreateUserReq.Name,
		Email:    CreateUserReq.Email,
		Password: CreateUserReq.Password,
	}

	config.DB.Create(&newUser) //You cannot pass a struct to ‘create’, so you should pass a pointer to the data.
	fmt.Println(newUser)

	c.JSON(200, gin.H{
		"message": "Success",
		"data":    newUser,
	})

}

func UpdateUser(c *gin.Context) {
	//Request Validation
	var UpdateUserReq api.UpdateUserRequest

	if err := c.ShouldBindJSON(&UpdateUserReq); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//Check if user exists
	var user models.User

	if err := config.DB.Where("id = ?", c.Params.ByName("id")).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "User not found!"})
		return
	}

	//Update User

	// config.DB.Model(&models.User{}).Where("id = ?", c.Params.ByName("id")).Updates(&UpdateUserReq)	//with and without &

	config.DB.Model(&user).Updates(&UpdateUserReq)

	c.JSON(200, gin.H{
		"message": "Success",
		"data":    user,
	})

}

func DeleteUser(c *gin.Context) {
	//Check if user exists
	var user models.User

	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "User not found!"})
		return
	}

	//Delete User

	config.DB.Delete(&user) //deleted_at column ----> soft delete
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    user,
	})

}
