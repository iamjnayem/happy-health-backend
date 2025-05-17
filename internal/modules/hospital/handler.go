package hospital


import (
	"github.com/gofiber/fiber/v2"
)

func GetHospitalHandler(c *fiber.Ctx) error {
	uid := c.Params("uid")
	hospital, err := GetHospitalDetails(uid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch hospital details",
		})
	}

	return c.JSON(hospital) 
}