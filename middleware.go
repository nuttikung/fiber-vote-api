package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func logMiddleware(c *fiber.Ctx) error {
	fmt.Printf("Method: %s, URL: %s \n", c.Method(), c.OriginalURL())
	return c.Next()
}
