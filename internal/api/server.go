package api

import (
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/handlers"
	"log"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	log.Printf("config DSN %v", config.Dsn)

	rh := &rest.RestHandler{
		App: app,
	}
	setupRoutes(rh)
	app.Listen(config.ServerPort)

}

func setupRoutes(rh *rest.RestHandler) {
	// user handlers
	handlers.SetupUserRoutes(rh) //rh input because this func require this arguments

	//transactions
	//catalog
}
