package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/markDoesany/QuickyKanban/internal/config"
	"github.com/markDoesany/QuickyKanban/internal/models"
	"gorm.io/gorm"
)

func CreateProject(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project.OwnerID = user.ID

	if err := config.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Project created successfully",
		"project": project,
	})
}

func GetProjects(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	var projects []models.Project

	if err := config.DB.Preload("Users").Where("owner_id = ?", user.ID).Or("id IN (?)",
		config.DB.Table("project_users").Select("project_id").Where("user_id = ?", user.ID)).Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

func GetProject(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unathorized access"})
	}

	idParam := c.Param("id")
	projectID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid project ID"})
		return
	}

	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	var project models.Project

	if err := config.DB.Preload("Users").Preload("Tasks").
		Where("id = ? AND (owner_id = ? OR id IN (SELECT project_id  FROM proect_users WHERE user_id = ?))",
			projectID, user.ID, user.ID).First(&project).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"project": project})
}

func UpdateProject(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	idParam := c.Param("id")
	projectID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid project ID"})
		return
	}

	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	var project models.Project

	if err := config.DB.Where("id = ? and owner_id", projectID, user.ID).Find(&project).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Project not found or unauthorized"})
		return
	}

	var updateProjectData struct {
		Name        string `json:"title"`
		Status      string `json:"status"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&updateProjectData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	project.Name = updateProjectData.Name
	project.Status = updateProjectData.Status
	project.Description = updateProjectData.Description

	if err := config.DB.Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Project updated",
		"project": project,
	})
}

func DeleteProject(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	idParam := c.Param("id")
	projectID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid project ID"})
		return
	}

	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	var project models.Project

	if err := config.DB.Where("id = ? and owner_id", projectID, user.ID).Find(&project).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Project not found or unauthorized"})
		return
	}

	if err := config.DB.Delete(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}
