package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VooDooStack/FitStackAPI/internal/comment"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

//Response - returns a pointer to a Response struct
type Response struct {
	Message string
	Error   error
}

func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func LoggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("endpoint hit")
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes...")
	h.Router = mux.NewRouter()
	h.Router.Use(LoggingMiddleWare)

	h.Router.HandleFunc("/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/comment", h.GetAllComment).Methods("GET")
	h.Router.HandleFunc("/comment/{id}", h.DeleteComment).Methods("DELETE")
	h.Router.HandleFunc("/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/comment/{id}", h.UpdateComment).Methods("PUT")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "OK"}); err != nil {
			panic(err)
		}
	})
}
