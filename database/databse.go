package database

import (
	"fmt"
	"sync"

	"main/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(config *config.Config) *sqlx.DB {
	var databaseInstance *sqlx.DB
	var once sync.Once
	var err error

	once.Do(func() {
		databaseInstance, err = sqlx.Connect("postgres", connectionStringFormat(config))
		if err != nil {
			panic(err.Error())
		}
	})

	return databaseInstance
}

func connectionStringFormat(config *config.Config) string {
	if config.DB_PASSWORD != "" {
		return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
			config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_SSLMODE)
	}
	return fmt.Sprintf("user=%s dbname=%s sslmode=%s", config.DB_USER, config.DB_NAME, config.DB_SSLMODE)
}

func InitGORM_DB(config *config.Config) *gorm.DB {
	gormDB, err := gorm.Open(postgres.Open(connectionStringFormat(config)))
	if err != nil {
		panic(fmt.Sprintf("GORM init error: &s", err.Error()))
	}

	return gormDB
}
