package entities

import "time"

type Book struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(200)" json:"title"`
	AuthorId    uint      `json:"author_id"`
	Author      Author    `gorm:"foreignKey:AuthorId" json:"author"`
	Description string    `json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type BookResponse struct {
	Id          uint               `json:"id"`
	Title       string             `json:"title"`
	AuthorId    uint               `json:"-"`
	Author      AuthorBookResponse `json:"author"`
	Description string             `json:"description"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}
