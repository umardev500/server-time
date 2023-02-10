package main

import (
	"fmt"
	"log"
	"os"
	"server-time/injector"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	app := fiber.New()
	injector.NewAppInjector(app)

	fmt.Printf("⚡️[server]: Server is running on porting %s\n", port)

	log.Fatal(app.Listen(port))
}
