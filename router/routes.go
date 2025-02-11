package router

import (
	"github.com/Thomika1/APIsGo.git/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()

	// Router declaration
	v1 := router.Group("api/v1")

	// Openings logic
	v1.GET("/opening", handler.ShowOpeningHandler)

	v1.POST("/opening", handler.CreateOpeningHandler)

	v1.DELETE("/opening", handler.DeleteOpeningHandler)

	v1.PUT("/opening", handler.UpdateOpeningHandler)

	v1.GET("/openings", handler.ListOpeningHandler)

} // func initialize routes
