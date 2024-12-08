package middleware

import (
	"log"

	"github.com/funukonta/management_toko/pkg/common"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next() // Proceed to the next middleware or handler
		if err != nil {
			statusCode, message := common.ParseError(err)

			log.Printf("Error: %v, Status Code: %d", err, statusCode)

			return c.Status(statusCode).JSON(common.RespError{
				Status:  statusCode,
				Message: message,
			})
		}
		return nil
	}
}
