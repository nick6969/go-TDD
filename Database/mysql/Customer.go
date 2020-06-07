package mysql

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Customer struct {
	ID        uint
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func TableCustomerMigrate(db *gorm.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS  customers  (
			id int(10) unsigned NOT NULL AUTO_INCREMENT,
			username varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
			password varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
			created_at datetime NOT NULL,
			updated_at datetime NOT NULL,
			deleted_at datetime DEFAULT NULL,
			PRIMARY KEY (id),
			UNIQUE KEY uix_customers_username (username),
			KEY idx_customers_deleted_at (deleted_at)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
		`

	return db.Exec(query).Error
}

func (db *Database) CheckUserNameCanUse(name string) bool {
	user := Customer{}
	err := db.Where(&Customer{Username: name}).First(&user).Error
	return err != nil
}

func (db *Database) CreateCustomer(username, password string) (id uint, err error) {

	customer := Customer{Username: username, Password: password}
	err = db.Create(&customer).Error

	if err != nil {
		return
	}

	id = customer.ID
	return
}
