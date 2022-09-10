package database

import (
	"ecommerce-project/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB(data *config.DataBase) *gorm.DB {
	alamat := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", data.DB_USERNAME, data.DB_PASSWORD, data.DB_HOST, data.DB_PORT, data.DB_NAME)

	db, err := gorm.Open(mysql.Open(alamat), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal Menghubungkan Ke DataBase", err)
	}

	return db
}
