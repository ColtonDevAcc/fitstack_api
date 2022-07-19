package handler

import (
	"fmt"

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
	var user models.User

	db := c.MustGet("db").(*gorm.DB)
	firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)
	u, err := firebaseAuth.GetUserByEmail(c, "cbristow99@gmail.com")
	if err != nil {
		fmt.Print(err)
	}

	userCheck := db.Where(models.User{DisplayName: u.DisplayName}).Find(&user)
	if userCheck != nil {
		c.JSON(200, gin.H{
			"error":  "user already exists",
			"result": user,
		})

		return
	} else {
		err = db.Create(&models.User{
			Uuid:        u.UID,
			DisplayName: u.DisplayName,
			Email:       u.Email,
		}).Error
		if err != nil {
			c.JSON(200, gin.H{
				"error": err.Error,
			})
		}

		c.JSON(200, gin.H{
			"user": u,
		})
	}
}
