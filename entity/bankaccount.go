package entity

import "time"

type BankAccount struct {
	ID                uint      `gorm:"primaryKey,autoIncrement"`
	UserId            uint      `gorm:"not null;type:integer"`
	BankName          string    `gorm:"not null;type:varchar(255)"`
	AccountName       string    `gorm:"not null;type:varchar(255)"`
	AccountBankNumber string    `gorm:"not null;type:varchar(255)"`
	CreatedAt         time.Time `gorm:"<-:create"`
	UpdatedAt         time.Time `gorm:"<-:update"`
	User              User      `gorm:"foreignkey:UserId;"  json:"user"`
}
