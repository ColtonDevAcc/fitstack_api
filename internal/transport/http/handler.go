package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/VooDooStack/FitStackAPI/internal/comment"
	"github.com/gorilla/mux"
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

func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes...")
	h.Router = mux.NewRouter()

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

//get comment
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("error getting comment error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		SendErrorResponse(w, "error getting comment:", err)
		return
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%+v\n", comment)
}

//get all comments
func (h *Handler) GetAllComment(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Println("error getting all comments:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(comments); err != nil {
		panic(err)
	}
}

//delete comment
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("error deleting comment error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.Service.DeleteComment(uint(i))
	if err != nil {
		fmt.Println("error deleting comment:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "comment deleted\n")
}

//post comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		fmt.Println("error decoding comment:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	comment, err := h.Service.PostComment(comment)
	if err != nil {
		SendErrorResponse(w, "error posting comment:", err)
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

//update comment
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("error updating comment error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	comment, err := h.Service.UpdateComment(uint(i), comment.Comment{
		Slug: "/new",
	})
	if err != nil {
		fmt.Println("error updating comment:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//! updated function to send responses
	if err = sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

func sendOkResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func SendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err}); err != nil {
		panic(err)
	}
}
