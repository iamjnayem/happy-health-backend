package patient

import "gorm.io/gorm"

type Patient struct {
    gorm.Model
    Name      string
    Phone     string
    Email     string
    OTP       string
    Verified  bool
}
