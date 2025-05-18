package dto
import "time"

type DepartmentResponse struct {
	Id          int        `json:"id"`
	NameEn      string     `json:"name_en,omitempty"`
	NameBn      string     `json:"name_bn,omitempty"`
	Description string     `json:"description,omitempty"`
	Status      int8       `json:"status"`
	Meta        *string    `json:"meta,omitempty"`
	Remarks     *string    `json:"remarks,omitempty"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty"`
}