package dto
import "time"
// HospitalResponse represents the response structure for a hospital.
type HospitalServiceResponse struct {
	Id        int        `json:"id"`
	NameEn    string     `json:"name_en,omitempty"`
	NameBn    string     `json:"name_bn,omitempty"`
	Price     float64    `json:"price,omitempty"`
	Status    int8       `json:"status"`
	Meta      *string    `json:"meta,omitempty"`
	Remarks   *string    `json:"remarks,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
}