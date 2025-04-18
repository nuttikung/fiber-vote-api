package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nuttikung/fiber-vote-api/controller"
	"gorm.io/gorm"
)

func InitialCandidateRoutes(router fiber.Router, db *gorm.DB) {
	candidateRoute := router.Group("/candidate")

	candidateRoute.Get("/", controller.CandidateList(db))
	candidateRoute.Get("/:id", controller.CandidateSingle(db))

}
