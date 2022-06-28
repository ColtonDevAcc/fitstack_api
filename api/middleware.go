package api

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware : to verify all authorized operations
// AuthMiddleware : to verify all authorized operations
func AuthMiddleware(c *gin.Context) {
	firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)
	authorizationToken := c.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
	if idToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id token not available"})
		c.Abort()
		return
	}

	//verify token
	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}
	c.Set("UUID", token.UID)
	c.Next()
}

//? What happens here is, We extract the Authorization header and then split the Authorization header value to get just the token.
//? We then, validate the token with firebaseAuth.verifyIDToken() method.
//? If the token is not valid we abort any further actions.
//! https://medium.com/wesionary-team/authenticate-rest-api-in-go-with-firebase-authentication-36cdf7c254c
