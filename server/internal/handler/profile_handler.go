package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/markDoesany/QuickyKanban/internal/config"
	"github.com/markDoesany/QuickyKanban/internal/models"
)

type ProfileUpdateRequest struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	ImageURL string `json:"profile_img"`
}

func GetProfile(c *gin.Context) {
	idParam := c.Param("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":          user.ID,
		"username":    user.Username,
		"profile_img": user.ImageURL,
	})
}

func UpdateProfile(c *gin.Context) {
	idParam := c.Param("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var updateData ProfileUpdateRequest
	if err := c.ShouldBindJSON(&updateData).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if updateData.Username != "" {
		user.Username = updateData.Username
	}

	if updateData.Role != "" {
		user.Role = updateData.Role
	}

	if updateData.ImageURL != "" {
		user.ImageURL = updateData.ImageURL
	}

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Profile updated successfully",
		"id":          user.ID,
		"role":        user.Role,
		"profile_img": user.ImageURL,
	})
}

func UploadProfileImage(c *gin.Context) {
	idParam := c.Param("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	uploadDir := "/static/profile_images"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	filename := fmt.Sprintf("user_%d%s", userID, filepath.Ext(file.Filename))
	filePath := filepath.Join(uploadDir, filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file in the server"})
		return
	}

	user.ImageURL = "/" + filePath
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Profile image uploaded successfully",
		"id":          user.ID,
		"role":        user.Role,
		"profile_img": user.ImageURL,
	})
}
