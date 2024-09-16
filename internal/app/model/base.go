package model

import "time"

type BaseModel struct {
	ID        uint      `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type User struct {
	BaseModel
	UserName string `gorm:"size:100;unique" json:"username"`
	Password string `gorm:"size:100" json:"password"`
	Email    string `gorm:"size:100;" json:"email"`
}
