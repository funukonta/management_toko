package routes

import (
	"github.com/funukonta/management_toko/internal/api/handlers"
	"github.com/funukonta/management_toko/internal/api/repositories"
	"github.com/funukonta/management_toko/internal/api/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(app fiber.Router, db *gorm.DB) {
	repoUser := repositories.NewRepoUser(db)
	ucUser := usecases.NewUsecaseUser(repoUser)
	handlers := handlers.NewHandlerUser(ucUser)

	app.Get("/users", handlers.GetUser)

	app.Post("/login", handlers.Login)
}
