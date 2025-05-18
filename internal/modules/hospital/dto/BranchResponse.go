package dto
import "time"

type BranchResponse struct {
	Id            int        `json:"id"`
	HospitalId    int        `json:"hospital_id"`
	NameEn        string     `json:"name_en,omitempty"`
	NameBn        string     `json:"name_bn,omitempty"`
	Address       string     `json:"address,omitempty"`
	Phone         string     `json:"phone,omitempty"`
	Email         string     `json:"email,omitempty"`
	GoogleMapLink string     `json:"google_map_link,omitempty"`
	Status        int8       `json:"status"`
	Meta          *string    `json:"meta,omitempty"`
	Remarks       *string    `json:"remarks,omitempty"`
	CreatedAt     time.Time  `json:"created_at,omitempty"`
	UpdatedAt     time.Time  `json:"updated_at,omitempty"`
}