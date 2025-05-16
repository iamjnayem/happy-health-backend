package patient

// RegisterPatientRequest is the request payload for registering a patient
type RegisterPatientRequest struct {
    Name  string `json:"name" validate:"required"`
    Phone string `json:"phone" validate:"required,e164"`  // You can add custom validators
    Email string `json:"email" validate:"required,email"`
}

// RegisterPatientResponse is the response payload after successful registration
type RegisterPatientResponse struct {
    Message string `json:"message"`
    OTP     string `json:"otp"`
}
