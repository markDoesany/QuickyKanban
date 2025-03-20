package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/markDoesany/QuickyKanban/internal/handler"
	"github.com/markDoesany/QuickyKanban/internal/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	router := r.Group("/api")
	{
		router.POST("/register", handler.Register)
		router.POST("/login", handler.Login)
	}

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/profile/:id", handler.GetProfile)
		protected.PUT("/profile/:id", handler.UpdateProfile)
		protected.POST("/profile/upload/:id", handler.UploadProfileImage)

		protected.POST("/projects", handler.CreateProject)
		protected.GET("/projects", handler.GetProjects)
		protected.GET("/projects/:id", handler.GetProject)
		protected.PUT("/projects/:id", handler.UpdateProject)
		protected.DELETE("/projects/:id", handler.DeleteProject)

		protected.GET("/comments", handler.GetComments)
		protected.POST("/comments", handler.PostComment)
	}
	// 	protected.POST("/tasks", handler.CreateTask)
	// 	protected.GET("/tasks", handler.GetTasks)
	// 	protected.GET("/tasks/:id", handler.GetTask)
	// 	protected.PUT("/tasks/:id", handler.UpdateTask)
	// 	protected.DELETE("/tasks/:id", handler.DeleteTask)
	// 	protected.POST("/tasks/:id/assign", handler.AssignUsersToTask)

	// }

}
