package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nuttikung/fiber-vote-api/models"
	"gorm.io/gorm"
)

func CandidateList(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		candidates, err := models.GetCandidates(db)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"data":  candidates,
					"error": err.Error(),
				})
		}

		return c.JSON(fiber.Map{
			"data":  candidates,
			"error": nil,
		})
	}
}

func CandidateSingle(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		candidate, err := models.GetCandidate(db, uint(id))

		if err != nil {
			return c.JSON(
				fiber.Map{
					"data":  nil,
					"error": err.Error(),
				})
		}

		return c.JSON(
			fiber.Map{
				"data":  candidate,
				"error": nil,
			})
	}
}
