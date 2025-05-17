package main

import (
    "github.com/gofiber/fiber/v2"
   	"github.com/iamjnayem/happy-health-backend/internal/config"
    "github.com/iamjnayem/happy-health-backend/internal/database"
    "github.com/iamjnayem/happy-health-backend/internal/modules/patient"
    "github.com/iamjnayem/happy-health-backend/internal/modules/hospital"
)

func main() {
    config.LoadConfig()
    database.Connect()

    app := fiber.New()

    api := app.Group("/api")
    v1 := api.Group("/v1")

    patient.RegisterRoutes(v1)
    hospital.RegisterRoutes(v1)

    app.Listen(":3000")
}
