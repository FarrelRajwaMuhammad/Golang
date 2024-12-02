package models

type User struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	UserID uint   `gorm:"not null" json:"userId"`
	Title  string `gorm:"type:text;not null" json:"title"`
	Body   string `gorm:"type:text;not null" json:"body"`
}
