package routes

import (
	"com/beeda/search/app/http/controllers"
	"com/beeda/search/app/repositories"
	"com/beeda/search/app/services"
	"com/beeda/search/bootstrap"

	"github.com/gofiber/fiber/v2"
)

func SetUpApiRoutes(api fiber.Router) {
	testRepo := repositories.NewTestRepository(bootstrap.App().DB)
	testService := services.NewTestService(testRepo)
	testController := controllers.NewTestController(testService)
	api.Get("/", testController.Test)
	api.Get("/test", testController.GetAllHandler)
}
