package main

import (
	"log"

	"github.com/nuttikung/fiber-vote-api/models"
	"gorm.io/gorm"
)

func getCandidates(db *gorm.DB) []models.Candidate {
	var candidates []models.Candidate

	result := db.Find(&candidates)

	if result.Error != nil {
		log.Fatalf("ERROR when execute: %v", result.Error)
	}

	return candidates
}

// func getCandidates(c *fiber.Ctx) error {
// 	rows, err := db.Query("SELECT * FROM candidate")
// 		log.Fatal(err)
// 	}

// 	var candidates []models.Candidate

// 	for rows.Next() {
// 		var c models.Candidate
// 		err := rows.Scan(&c.ID, &c.Name, &c.Profile, &c.Vote)
// 		if err != nil {
// 			break
// 			// return nil, err
// 		}
// 		candidates = append(candidates, c)
// 	}

// 	if err = rows.Err(); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"data":  candidates,
// 			"error": err.Error(),
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"data":  candidates,
// 		"error": nil,
// 	})
// }

// func getCandidate(c *fiber.Ctx) error {
// 	var id int
// 	var err error

// 	if id, err = strconv.Atoi(c.Params("id")); err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 	}

// 	var candidate models.Candidate
// 	row := db.QueryRow(
// 		"SELECT * FROM candidate WHERE id=$1;",
// 		id,
// 	)

// 	err = row.Scan(&candidate.ID, &candidate.Name, &candidate.Profile, &candidate.Vote)

// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"data":  candidate,
// 			"error": err.Error(),
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"data":  candidate,
// 		"error": nil,
// 	})

// }
