package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/wellminozzo/mvcgo/commands"
	"github.com/wellminozzo/mvcgo/controllers"
	"github.com/wellminozzo/mvcgo/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.Connection()
}

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		commands.Run()
	} else {
		//app fiber
		app := fiber.New()

		//rotas
		app.Get("/", controllers.HelloIndex)
		app.Get("/json", controllers.HelloJson)

		//init
		app.Listen(":" + os.Getenv("WEB_PORT"))
	}
}
