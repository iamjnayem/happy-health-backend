package patient

import "github.com/iamjnayem/happy-health-backend/internal/database"

func CreatePatient(patient *Patient) error {
    return database.DB.Create(patient).Error
}
