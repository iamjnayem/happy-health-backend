package dto 

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