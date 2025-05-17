package hospital 

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	hospitalGroup := router.Group("/hospital")
	hospitalGroup.Get("/:uid", GetHospitalHandler)

}
