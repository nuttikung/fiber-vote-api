package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nuttikung/fiber-vote-api/config"
	_ "github.com/nuttikung/fiber-vote-api/docs"
	"github.com/nuttikung/fiber-vote-api/router"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// var candidates []models.Candidate

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
	config.InitializeEnv()

	db := config.InitializeDatabase()
	// defer db.Close()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${method} | ${status} | ${path}",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Health Check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	// app.Post("/api/v1/vote", postVote)
	app.Post("/api/v1/upload", postUploadFile)

	router.InitializeRoutes(app, db)

	// running on port
	port := os.Getenv("PORT")
	app.Listen(fmt.Sprintf(":%v", port))
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
