package patient

func RegisterPatient(name, phone, email string) error {
    patient := &Patient{
        Name:  name,
        Phone: phone,
        Email: email,
        OTP:   "123456", // Dummy OTP for now
    }
    return CreatePatient(patient)
}
