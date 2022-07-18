package handler

import (
	"fmt"
	"time"

	"firebase.google.com/go/auth"
	"github.com/VooDooStack/FitStackAPI/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func userSignIn(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)
	user, err := firebaseAuth.GetUserByEmail(c, "cbristow99@gmail.com")
	if err != nil {
		fmt.Print(err)
	}

	db.Get("")

	c.JSON(200, gin.H{
		"message": user.UID,
	})
}

func signUp(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)
	user, err := firebaseAuth.GetUserByEmail(c, "cbristow99@gmail.com")
	if err != nil {
		fmt.Print(err)
	}

	newUser := models.User{
		DisplayName: user.DisplayName,
		Uuid:        user.UID,
		Email:       user.Email,
		UpdatedAt:   time.Now().GoString(),
		CreatedAt:   user.UserMetadata.CreationTimestamp,
	}

	db.Create(&newUser)

	c.JSON(200, gin.H{
		"message": newUser,
	})
}
