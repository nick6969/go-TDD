package mysql

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	*gorm.DB
}

// ConnectDB is Connect database
func ConnectDB() (*Database, error) {

	database, err := gorm.Open("mysql", os.Getenv("MYSQL_DATABASE")+"?parseTime=true")

	if err != nil {
		return nil, err
	}

	database.LogMode(os.Getenv("DEBUG_MODE") == "1")

	err = databaseMigrate(database)

	if err != nil {
		return nil, err
	}

	return &Database{database}, nil
}

func databaseMigrate(db *gorm.DB) (err error) {

	err = TableCustomerMigrate(db)
	if err != nil {
		return
	}

	return
}
