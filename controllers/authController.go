package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/mongo"
)

func Signup(c *gin.Context) {
	fmt.Println("reached here at signup")

	c.JSON(200, gin.H{"message": "Signup successfull"})
}

func Signin(c *gin.Context) {
	fmt.Println("reached here at signin")
	c.JSON(200, gin.H{"message": "Signin successfull"})
}

func Getusers(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Signin successfull"})
}

func Getuser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Signin successfull"})
}
