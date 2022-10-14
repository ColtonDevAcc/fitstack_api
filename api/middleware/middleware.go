package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/VooDooStack/FitStackAPI/infrastructure/secrets"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	apiKeyHeader        = "X-API-Key"
	cronExecutedHeader  = "X-Appengine-Cron"
	valName             = "FIREBASE_ID_TOKEN"
)

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

// Gin middleware for JWT auth
func AuthJWT(client *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get(authorizationHeader)
		token := strings.Replace(authHeader, "Bearer ", "", 1)
		idToken, err := client.VerifyIDToken(c, token) // usually hits a local cache
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}
		uuid := idToken.UID
		c.Set("token", token)
		c.Set("uuid", uuid)
		c.Next()
	}
}

// API key auth middleware
func AuthAPIKey(secretId string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.Header.Get(apiKeyHeader)

		secret, err := secrets.GetSecret(secretId)
		if err != nil {
			log.Println("failed to get secret")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		log.Println("comparing secret with provided key", secret, key)

		if secret != key {
			log.Println("key does not match!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		log.Println("no error during check")
		c.Next()
	}
}

//? What happens here is, We extract the Authorization header and then split the Authorization header value to get just the token.
//? We then, validate the token with firebaseAuth.verifyIDToken() method.
//? If the token is not valid we abort any further actions.
//! https://medium.com/wesionary-team/authenticate-rest-api-in-go-with-firebase-authentication-36cdf7c254c
