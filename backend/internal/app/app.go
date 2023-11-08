package app

import (
	"chat/config"
	"chat/internal/api/http/v1/routes"
	"chat/internal/repository"
	"chat/internal/service"
	"chat/internal/usecase"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	// Init repository
	messageRepo, err := repository.New(cfg.StoragePath)
	if err != nil {
		panic(err)
	}

	// Init services
	messageSvc := service.NewMessageService(messageRepo)

	// Init usecases
	messageUc := usecase.NewMessageUseCase(messageSvc)

	// HTTP server
	router := gin.Default()
	routes.New(router, messageUc)
	router.Run(cfg.Address)
}
