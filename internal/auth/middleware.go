package auth

import (
	"log"
	"net/http"
	"strings"
	"time"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

//OAuth
//APIKey
//Cron

const (
	authorizationHeader  = "Authorization"
	apiKeyHeader         = "X-API-Key"
	cronExecuteJobHeader = "X-Cron-Execute-Job"
	valName              = "firebaseIdToken"
)

//! gin middleware for JWT Auth
func AuthJWT(client *auth.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		authHeader := ctx.Request.Header.Get(authorizationHeader)
		log.Print("authHeader: ", authHeader)
		token := strings.Replace(authHeader, "Bearer ", "", 1)
		IdToken, err := client.VerifyIDToken(ctx, token) // usually hits local cache
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
		}

		log.Println("Auth Time: ", time.Since(startTime))

		ctx.Set(valName, IdToken)
		ctx.Next()
	}
}
