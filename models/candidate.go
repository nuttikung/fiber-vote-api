package models

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Candidate struct {
	gorm.Model
	// ID      uint 	`gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
	Vote    int32  `json:"vote"`
}

func GetCandidates(db *gorm.DB) ([]Candidate, error) {
	var candidates []Candidate

	result := db.Find(&candidates)

	if result.Error != nil {
		log.Fatalf("ERROR when execute: %v", result.Error)
		return nil, result.Error
	}

	if len(candidates) == 0 {
		return []Candidate{}, nil
	}

	return candidates, nil
}

func GetCandidate(db *gorm.DB, id uint) (Candidate, error) {
	var candidate Candidate

	result := db.Find(&candidate, id)

	if result.Error != nil {
		log.Fatalf("ERROR when execute: %v", result.Error)
		return candidate, result.Error
	}

	if result.RowsAffected == 0 {
		fmt.Printf("Vannot find candidate id: %d", id)
		// log.Fatalf("ERROR when execute: %v", result.Error)
		return candidate, result.Error
	}

	return candidate, nil
}
