package router

import (
	"github.com/Thomika1/APIsGo.git/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	// Router declaration
	v1 := router.Group("api/v1")

	// Openings logic
	v1.GET("opening", handler.CreateOpeningHandler)

	v1.POST("opening", handler.ShowOpeningHandler)

	v1.DELETE("opening", handler.DeleteOpeningHandler)

	v1.PUT("opening", handler.UpdateOpeningHandler)

	v1.GET("openings", handler.ListOpeningHandler)

} // func initialize routes
