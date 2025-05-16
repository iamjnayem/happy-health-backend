package patient

import "github.com/gofiber/fiber/v2"

func RegisterPatientHandler(c *fiber.Ctx) error {
    var req RegisterPatientRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
    }

    // Ideally you'd validate req fields here

    if err := RegisterPatient(req.Name, req.Phone, req.Email); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Could not register patient"})
    }

    resp := RegisterPatientResponse{
        Message: "Patient registered. OTP sent.",
        OTP:     "123456", // Dummy OTP, in real app generate dynamically
    }

    return c.JSON(resp)
}
