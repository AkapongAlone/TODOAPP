package databases

import (
	"fmt"
	"os"
	"sync"

	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	sqlDBOnce     sync.Once
	sqlDBInstance *gorm.DB
)

func NewPostgres() *gorm.DB {
	sqlDBOnce.Do(func() {
	AGIAN:
		DB_HOST := os.Getenv("DB_HOST")
		DB_USERNAME := os.Getenv("DB_USERNAME")
		DB_PASSWORD := os.Getenv("DB_PASSWORD")
		DB_DATABASE := os.Getenv("DB_DATABASE")
		DB_PORT := os.Getenv("DB_PORT")

		var err error
		// PostgreSQL DSN format
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
			DB_HOST, DB_USERNAME, DB_PASSWORD, DB_DATABASE, DB_PORT)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			goto AGIAN // Consider adding a sleep or a maximum retry count to avoid an infinite loop
		}

		sqlDBInstance = db
	})
	return sqlDBInstance
}

