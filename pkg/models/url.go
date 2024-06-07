package models

type URL struct {
	ID        uint   `gorm:"primaryKey"`
	LongURL   string `gorm:"not null"`
	ShortCode string `gorm:"uniqueIndex;not null"`
}
