package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	"database/sql"

	_ "github.com/lib/pq"
	_ "github.com/nuttikung/fiber-vote-api/docs"
)

var db *sql.DB

var candidates []Candidate

// @title Most popular Mascot in Thailand API
// @version 1.0
// @description This is a sample go api for vote Mascot in Thailand.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host example.swagger.io
// @BasePath /api/v1
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("LOAD ENV error")
	}

	connStr := os.Getenv("DATABASE_URL")

	_db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	db = _db

	defer db.Close()

	println("Connect Database Successful")

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logMiddleware)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Health Check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/api/v1/candidate", getCandidates)
	app.Get("/api/v1/candidate/:id", getCandidate)

	app.Post("/api/v1/upload", postUploadFile)

	app.Post("/api/v1/vote", postVote)

	app.Listen(":3000")
}

func postUploadFile(c *fiber.Ctx) error {
	var uploadPath string = "./upload/"
	file, err := c.FormFile("image")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = c.SaveFile(file, uploadPath+file.Filename)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("File upload complete.")
}
