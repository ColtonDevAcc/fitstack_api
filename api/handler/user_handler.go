package handler

import (
	"fmt"

	"firebase.google.com/go/auth"
	"github.com/VooDooStack/FitStackAPI/data/usecases"
	"github.com/VooDooStack/FitStackAPI/domain"
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
	var user domain.User
	db := c.MustGet("db").(*gorm.DB)
	firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)

	u, err := firebaseAuth.GetUserByEmail(c, "cbristow99@gmail.com")
	if err != nil {
		fmt.Print(err)
	}

	user, err = usecases.GetUserByEmail(db, u.Email)
	if err != nil {
		fmt.Print(err)
		c.JSON(500, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}
