package entity

import "time"

type Balance struct {
	ID        uint      `gorm:"primaryKey,autoIncrement"`
	Amount    float64   `gorm:"not null;type:float"`
	UserId    uint      `gorm:"not null;type:integer"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time `gorm:"<-:update"`
}
