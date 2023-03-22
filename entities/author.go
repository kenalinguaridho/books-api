package entities

import "time"

type Author struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	Gender    string    `gorm:"type:char(1)" json:"gender"`
	Email     string    `gorm:"type:varchar(100)" json:"email"`
	Age       int       `gorm:"type:int" json:"age"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoCreateTime" json:"updated_at"`
}

type AuthorBookResponse struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Email     string    `json:"email"`
}