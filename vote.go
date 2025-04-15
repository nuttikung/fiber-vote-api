package main

import "github.com/gofiber/fiber/v2"

type VoteRequest struct {
	ID int `json:"id"`
}

func postVote(c *fiber.Ctx) error {
	voteRequest := new(VoteRequest)

	if err := c.BodyParser(voteRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, candidate := range candidates {
		if candidate.ID == voteRequest.ID {
			candidates[i].Vote += 1
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": candidates[i]})
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}
