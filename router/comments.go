package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/VooDooStack/FitStackAPI/internal/comment"
	"github.com/gin-gonic/gin"
)

// GetComment - retrieve a comment by ID
func (h *Handler) GetComment(c *gin.Context) {
	id := c.Param("id")

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to parse json body")
	}
	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to get comment by ID")
	}

	c.JSON(http.StatusOK, gin.H{"comment": comment})
}

// GetAllComments - retrieves all comments from the comment service
func (h *Handler) GetAllComments(c *gin.Context) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to retrieve all comments")
	}

	c.JSON(http.StatusOK, comments)
}

// PostComment - adds a new comment
func (h *Handler) PostComment(c *gin.Context) {

	var comment comment.Comment
	if err := json.NewDecoder(c.Request.Body).Decode(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to decode JSON Body")
	}

	comment, err := h.Service.PostComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to post new comment")
	}

	c.JSON(http.StatusOK, comment)
}

// UpdateComment - updates a comment by ID
func (h *Handler) UpdateComment(c *gin.Context) {
	id := c.Param("id")

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to parse uint from ID")
	}

	var comment comment.Comment
	if err := json.NewDecoder(c.Request.Body).Decode(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to decode JSON Body")

	}

	comment, err = h.Service.UpdateComment(uint(commentID), comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to update comment")
	}

	c.JSON(http.StatusOK, comment)
}

// DeleteComment - deletes a comment by ID
func (h *Handler) DeleteComment(c *gin.Context) {
	id := c.Param("id")
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to parse uint from ID")

	}

	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to delete comment by comment ID")
	}

	//TODO: return a 204 status code
}
