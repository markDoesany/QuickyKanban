package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markDoesany/QuickyKanban/internal/config"
	"github.com/markDoesany/QuickyKanban/internal/models"
)

func GetComments(c *gin.Context) {
	taskID := c.Query("task_id")

	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is missing"})
		return
	}

	var comments []models.Comment
	if err := config.DB.Preload("User").Where("task_id = ?", taskID).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

func PostComment(c *gin.Context) {
	var comment models.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if comment.TaskID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is missing"})
	}

	if err := config.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Faied to post comment"})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Commend added successfully",
		"comment": comment,
	})
}
