package models

type URL struct {
	ID        uint `gorm:"primaryKey"`
	LongURL   uint `gorm:"not null"`
	ShortCode uint `gorm:"uniqueIndex;not null"`
}
