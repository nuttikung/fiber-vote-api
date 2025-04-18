package main

import (
	"fmt"
	"log"

	"github.com/nuttikung/fiber-vote-api/models"
	"gorm.io/gorm"
)

type VoteRequest struct {
	ID int `json:"id"`
}

func voteCandidate(db *gorm.DB, id uint) {
	var candidate models.Candidate

	db.First(&candidate, id)

	candidate.Vote = candidate.Vote + 1

	result := db.Model(&candidate).Updates(candidate)

	if result.Error != nil {
		log.Fatalf("Update Candidate Failed: %v", result.Error)
	}

	fmt.Println("Update Candidate Success")
}

// func postVote(c *fiber.Ctx) error {
// 	voteRequest := new(VoteRequest)

// 	if err := c.BodyParser(voteRequest); err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 	}

// 	for i, candidate := range candidates {
// 		if candidate.ID == voteRequest.ID {
// 			candidates[i].Vote += 1
// 			return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": candidates[i]})
// 		}
// 	}

// 	return c.SendStatus(fiber.StatusNotFound)
// }
