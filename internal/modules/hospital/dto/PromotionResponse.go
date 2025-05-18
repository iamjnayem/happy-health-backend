package dto 
import "time"

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