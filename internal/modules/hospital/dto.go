package hospital

import "time"

type HospitalResponse struct {
	Id        int       `json:"id"`
	Uid       string    `json:"uid,omitempty"`
	NameEn    string    `json:"name_en,omitempty"`
	NameBn    string    `json:"name_bn,omitempty"`
	Address   string    `json:"address,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Email     string    `json:"email,omitempty"`
	Status    int8      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Meta      *string   `json:"meta,omitempty"`
	Remarks   *string   `json:"remarks,omitempty"`
}

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

type PromotionResponse struct {
	Id          int        `json:"id"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Type        string     `json:"type,omitempty"`
	ExpireAt    *time.Time `json:"expire_at,omitempty"`
	Status      int8       `json:"status"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty"`
}

type HospitalFullResponse struct {
	Hospital         HospitalResponse          `json:"hospital"`
	Branches         []BranchResponse          `json:"branches"`
	Departments      []DepartmentResponse      `json:"departments"`
	HospitalServices []HospitalServiceResponse `json:"service_categories"`
	Promotions       []PromotionResponse       `json:"offers"`
}
