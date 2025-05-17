package hospital 


import (
	"gorm.io/gorm"
	"time"
)


type Hospital struct {
	gorm.Model
	Id 	   uint   `json:"id" gorm:"primaryKey"`
	UID string `json:"uid" gorm:"type:varchar(255);not null;unique"`
	NameEn string `json:"name_en" gorm:"type:varchar(255);not null;unique"`
	NameBn string `json:"name_bn" gorm:"type:varchar(255);not null;unique"`
	Address string `json:"address" gorm:"type:text;not null;unique"`
	Phone  string `json:"phone" gorm:"type:varchar(20);not null;unique"`
	Email  string `json:"email" gorm:"type:varchar(30);not null;unique"`
	Status int8  `json:"status" gorm:"type:tinyint(1);not null;default:1"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}