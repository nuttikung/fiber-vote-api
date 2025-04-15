package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Candidate struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
	Vote    int    `json:"vote"`
}

func getCandidates(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT * FROM candidate")

	if err != nil {
		log.Fatal(err)
	}

	var candidates []Candidate
	for rows.Next() {
		var c Candidate
		err := rows.Scan(&c.ID, &c.Name, &c.Profile, &c.Vote)
		if err != nil {
			break
			// return nil, err
		}
		candidates = append(candidates, c)
	}

	if err = rows.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":  candidates,
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":  candidates,
		"error": nil,
	})
}

func getCandidate(c *fiber.Ctx) error {
	var id int
	var err error

	if id, err = strconv.Atoi(c.Params("id")); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var candidate Candidate
	row := db.QueryRow(
		"SELECT * FROM candidate WHERE id=$1;",
		id,
	)

	err = row.Scan(&candidate.ID, &candidate.Name, &candidate.Profile, &candidate.Vote)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":  candidate,
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":  candidate,
		"error": nil,
	})

}
