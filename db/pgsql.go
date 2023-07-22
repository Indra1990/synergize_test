package db

import (
	"fmt"
	"log"
	"os"
	"synergize/entity"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDatabasePgSQLConnection() (*gorm.DB, error) {
	err := initToYmal()
	if err != nil {
		fmt.Println("error", err.Error())
		return nil, err
	}

	dbUser := viper.GetString("database.pgsql.user")
	dbPass := viper.GetString("database.pgsql.password")
	dbHost := viper.GetString("database.pgsql.host")
	dbName := viper.GetString("database.pgsql.database")
	dbPort := viper.GetString("database.pgsql.port")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			IgnoreRecordNotFoundError: false,         // Ignore ErrRecordNotFound error for logger
			LogLevel:                  logger.Silent, // Log level
			Colorful:                  true,          // Disable color
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
	})

	if dbErr != nil {
		fmt.Println(dbErr)
		return nil, dbErr
	}

	if !db.Migrator().HasTable("users") {
		if err := db.Migrator().CreateTable(&entity.User{}); err != nil {
			fmt.Println(err.Error())
			panic(err.Error())
		}
	}

	if !db.Migrator().HasTable("bank_accounts") {
		if err := db.Migrator().CreateTable(&entity.BankAccount{}); err != nil {
			fmt.Println(err.Error())
			panic(err.Error())
		}
	}

	if !db.Migrator().HasTable("transactions") {
		if err := db.Migrator().CreateTable(&entity.Transaction{}); err != nil {
			fmt.Println(err.Error())
			panic(err.Error())
		}
	}

	if !db.Migrator().HasTable("balances") {
		if err := db.Migrator().CreateTable(&entity.Balance{}); err != nil {
			fmt.Println(err.Error())
			panic(err.Error())
		}
	}

	fmt.Println("connected to db ")
	return db, nil

}

func initToYmal() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("synergize")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
