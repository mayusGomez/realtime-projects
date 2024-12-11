package web

import (
	"github.com/gin-gonic/gin"
	"livecomments/dispatcher/domain"
	"net/http"
)

type NewComment struct {
	Video        string `json:"video" binding:"required"`
	ConnectionId string `json:"connection_id" binding:"required"`
	Comment      string `json:"comment" binding:"required"`
}

type CommentHandler struct {
	postCommentCmd domain.PostCommentCmd
}

func NewCommentHandler(postCommentCmd domain.PostCommentCmd) *CommentHandler {
	return &CommentHandler{
		postCommentCmd: postCommentCmd,
	}
}

func (handler *CommentHandler) Handle(c *gin.Context) {
	var req NewComment
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := handler.postCommentCmd.PostComment(req.ConnectionId, req.Video, req.Comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})

}
