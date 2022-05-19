package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func get_method() {

	app := fiber.New()

	app.Get("/:cloud.api?", func(ctx *fiber.Ctx) error {
		fmt.Println(ctx.Params("name"))
		result := ctx.Params("name")
		buffer := make([]byte, len(result))
		copy(buffer, result)
		resultCp := string(buffer)
		fmt.Println(resultCp)
		return ctx.SendString("hello")
	})
	app.Listen(":3000")
}
