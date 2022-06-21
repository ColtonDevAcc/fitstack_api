package router

import (
	"net/http"

	"github.com/VooDooStack/FitStackAPI/internal/auth"
	"github.com/VooDooStack/FitStackAPI/internal/comment"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router  *gin.Engine
	Service *comment.Service
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() *gin.Engine {
	client, err := auth.InitAuth()
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Setting Up Routes")
	// initialize new gin engine (for server)
	h.Router = gin.Default()

	// setup our routes
	versioned := h.Router.Group("/v1/")
	h.Router.Use(auth.AuthJWT(client))

	versioned.GET("/api/comment", h.GetAllComments)
	versioned.POST("/api/comment", h.PostComment)
	versioned.GET("/api/comment/{id}", h.GetComment)
	versioned.PUT("/api/comment/{id}", h.UpdateComment)
	versioned.DELETE("/api/comment/{id}", h.DeleteComment)
	versioned.GET("/api/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return h.Router
}
