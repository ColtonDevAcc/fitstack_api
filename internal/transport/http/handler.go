package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/VooDooStack/FitStackAPI/internal/auth"
	"github.com/VooDooStack/FitStackAPI/internal/comment"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router  *gin.Engine
	Service *comment.Service
}

// Response object
type Response struct {
	Message string
	Error   string
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// LoggingMiddleware - a handy middleware function that logs out incoming requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(
			log.Fields{
				"Method": r.Method,
				"Path":   r.URL.Path,
			}).
			Info("handled request")
		next.ServeHTTP(w, r)
	})
}

// BasicAuth - a handy middleware function that will provide basic auth around specific endpoints
func BasicAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("basic auth endpoint hit")
		user, pass, ok := r.BasicAuth()
		if user == "admin" && pass == "password" && ok {
			original(w, r)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			sendErrorResponse(w, "not authorized", errors.New("not authorized"))
		}
	}
}

// validateToken - validates an incoming jwt token
func validateToken(accessToken string) bool {
	var mySigningKey = []byte("missionimpossible")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return mySigningKey, nil
	})

	if err != nil {
		return false
	}

	return token.Valid
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

	versioned.GET("/api/comment", h.GetAllComments).Methods("GET")
	versioned.POST("/api/comment", JWTAuth(h.PostComment)).Methods("POST")
	versioned.GET("/api/comment/{id}", h.GetComment).Methods("GET")
	versioned.PUT("/api/comment/{id}", JWTAuth(h.UpdateComment)).Methods("PUT")
	versioned.DELETE("/api/comment/{id}", JWTAuth(h.DeleteComment)).Methods("DELETE")
	versioned.GET("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I am Alive!"}); err != nil {
			panic(err)
		}
	})

	return h.Router
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
