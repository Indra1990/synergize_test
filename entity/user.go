package entity

import "time"

type User struct {
	ID          uint         `gorm:"primaryKey,autoIncrement"`
	Username    string       `gorm:"not null;type:varchar(255)"`
	Email       string       `gorm:"not null;unique;type:varchar(255)"`
	PhoneNumber string       `gorm:"not null;unique;type:varchar(255)"`
	Password    string       `gorm:"not null;type:varchar(255)"`
	CreatedAt   time.Time    `gorm:"<-:create"`
	UpdatedAt   time.Time    `gorm:"<-:update"`
	BankAccount *BankAccount `gorm:"foreignkey:UserId;" json:"user"`
	Balance     *Balance     `gorm:"foreignkey:UserId;" json:"balance"`
}

type UserListWithBankAccount struct {
	ID                   uint
	Username             string
	Email                string
	PhoneNumber          string
	CreatedAtUser        string
	UpdatedAtUser        string
	BankName             string
	AccountName          string
	AccountBankNumber    string
	CreatedAtBankAccount string
	UpdatedAtBankAccount string
	Balance              string
}
