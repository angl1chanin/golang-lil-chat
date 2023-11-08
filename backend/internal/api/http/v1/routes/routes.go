package routes

import (
	"chat/internal/api/http/v1/handlers"
	"chat/internal/usecase"
	"github.com/gin-gonic/gin"
)

func New(router *gin.Engine, messages usecase.MessageUseCase) {
	// init router group
	v1 := router.Group("/v1")
	chat := v1.Group("/chat")

	handlers := handlers.Init(messages)

	chat.GET("/messages", handlers.Get)
	chat.POST("/messages", handlers.Create)
	chat.DELETE("/messages/:id", handlers.Delete)
}
