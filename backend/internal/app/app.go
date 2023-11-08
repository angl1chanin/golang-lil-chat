package app

import (
	"chat/config"
	"chat/internal/api/http/v1/routes"
	"chat/internal/repository"
	"chat/internal/service"
	"chat/internal/usecase"
	"github.com/gin-contrib/cors"
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

	// Настройка CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}                                                                                                               // Разрешенный домен
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}                                                                         // Разрешенные методы
	config.AllowHeaders = []string{"Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With", "ngrok-skip-browser-warning"} // Разрешенные заголовки

	// Добавление CORS middleware в роутер
	router.Use(cors.New(config))

	routes.New(router, messageUc)
	router.Run(cfg.Address)
}
