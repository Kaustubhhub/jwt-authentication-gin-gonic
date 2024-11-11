package controllers

import (
	"context"
	"fmt"
	"jwt_authentication_gin_gonic/database"

	helper "jwt_authentication_gin_gonic/helpers"
	"log"
	"net/http"

	// "strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	// "go.mongodb.org/mongo-driver/mongo"
	"jwt_authentication_gin_gonic/models"
	// "golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword() {

}

func VerifyPassword() {

}

func Signup(c *gin.Context) {
	var user models.User
	if c.ShouldBind(&user) == nil {
		fmt.Println("insode if")
		log.Println(*user.First_name)
		log.Println(*user.Last_name)
		log.Println(*user.Password)
	}
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

func Getuser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
