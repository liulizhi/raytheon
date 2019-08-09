package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// InitDB init db information
func InitDB(connURL string) (db *gorm.DB, err error) {
	db, err = gorm.Open(DBTYPE, connURL)
	if err != nil {
		return
	}

	if db.Error != nil {
		err = db.Error
		return
	}
	return
}
