package patient

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(api fiber.Router) {
    group := api.Group("/patients")
    group.Get("/register", RegisterPatientHandler)
}
