package router

import (
	docs "github.com/Thomika1/APIsGo.git/docs"

	"github.com/Thomika1/APIsGo.git/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath "api/v1"

// PingExample godoc
// @Summary
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /router

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()
	basePath := "api/v1"
	docs.SwaggerInfo.BasePath = basePath

	// Router declaration
	v1 := router.Group(basePath)

	// Openings logic
	v1.GET("/opening", handler.ShowOpeningHandler)

	v1.POST("/opening", handler.CreateOpeningHandler)

	v1.DELETE("/opening", handler.DeleteOpeningHandler)

	v1.PUT("/opening", handler.UpdateOpeningHandler)

	v1.GET("/openings", handler.ListOpeningHandler)

	// Initialize swagger

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

} // func initialize routes
