package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(user, password, host, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		user,
		password,
		host,
		database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// err = db.AutoMigrate(&dao.User{})
	// if err != nil {
	// 	return nil, err
	// }

	return db, nil
}
