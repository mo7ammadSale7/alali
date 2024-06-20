package Models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `json:"username" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Token    string `json:"token" gorm:"type:varchar(255)"`
}
