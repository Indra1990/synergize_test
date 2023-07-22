package entity

import "time"

type Transaction struct {
	ID            uint      `gorm:"primaryKey,autoIncrement"`
	Type          string    `gorm:"not null;type:varchar(255)"`
	Status        string    `gorm:"not null;type:varchar(255)"`
	Amount        float64   `gorm:"not null;type:float"`
	Notes         string    `gorm:"type:text"`
	UserId        uint      `gorm:"not null;type:integer"`
	AccountBankId uint      `gorm:"not null;type:integer"`
	CreatedAt     time.Time `gorm:"<-:create"`
	UpdatedAt     time.Time `gorm:"<-:update"`
}
