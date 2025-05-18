package hospital 

import (
	"github.com/iamjnayem/happy-health-backend/internal/database"
)

func GetHospitalByUid(uid string) (*Hospital, error) {
	var hospital Hospital
	result := database.DB.Where("uid = ?", uid).First(&hospital)
	if result.Error != nil {
		return nil, result.Error
	}
	return &hospital, nil
}
