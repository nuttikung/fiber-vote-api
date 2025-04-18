package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitializeRoutes(app *fiber.App, db *gorm.DB) {
	v1 := app.Group("/api/v1")

	InitialCandidateRoutes(v1, db)
}
