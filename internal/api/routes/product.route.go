package routes

import (
	"github.com/funukonta/management_toko/internal/api/handlers"
	"github.com/funukonta/management_toko/internal/api/repositories"
	"github.com/funukonta/management_toko/internal/api/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProductRoutes(app fiber.Router, db *gorm.DB) {
	repoProduct := repositories.NewRepoProduct(db)
	ucProduct := usecases.NewUsecaseProduct(repoProduct)
	handlers := handlers.NewHandlerProduct(ucProduct)

	product := app.Group("/products")
	{
		product.Post("", handlers.AddProduct)
		product.Get("", handlers.GetAllProduct)
		product.Get("/:id", handlers.GetDetailProduct)
		product.Put("/:id", handlers.UpdateProduct)
		product.Delete("/:id", handlers.DeleteProduct)
	}

}
