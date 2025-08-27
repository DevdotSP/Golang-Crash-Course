package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func SampleFiber() {
	app := fiber.New()

	app.Post("/customer", func(c fiber.Ctx) error {
		var customer Customer
		if err := c.Bind().Body(&customer); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		fmt.Printf("Received: %+v\n", customer)
		return c.JSON(customer)
	})

	log.Fatal(app.Listen(":3000"))
}
