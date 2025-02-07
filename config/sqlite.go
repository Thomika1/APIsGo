package config

import (
	"os"

	"github.com/Thomika1/APIsGo.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.go"

	// Checks if db file exists
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("->database file not found...")
		// Create database  file\
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("->sqlite opening error: %v", err)
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("->sqlite automigration error: %v", err)
	}

	// Return db
	return db, nil
}
