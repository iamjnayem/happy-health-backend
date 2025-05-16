package main

import (
    "github.com/gofiber/fiber/v2"
   	"github.com/iamjnayem/happy-health-backend/internal/config"
    "github.com/iamjnayem/happy-health-backend/internal/database"
    "github.com/iamjnayem/happy-health-backend/internal/modules/patient"
)

func main() {
    config.LoadConfig()
    database.Connect()

    app := fiber.New()

    api := app.Group("/api")
    patient.RegisterRoutes(api)

    app.Listen(":3000")
}
