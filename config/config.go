package config

import "os"

type DataBase struct {
	DB_SERVERPORT string
	DB_USERNAME   string
	DB_PASSWORD   string
	DB_HOST       string
	DB_PORT       string
	DB_NAME       string
}

var dbAddress *DataBase

func GetDataBase() *DataBase {
	if dbAddress == nil {
		dbAddress = InitDB()
	}
	return dbAddress
}

func InitDB() *DataBase {
	var data DataBase
	data.DB_SERVERPORT = os.Getenv("DB_SERVERPORT")
	data.DB_USERNAME = os.Getenv("DB_USERNAME")
	data.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	data.DB_HOST = os.Getenv("DB_HOST")
	data.DB_PORT = os.Getenv("DB_PORT")
	data.DB_NAME = os.Getenv("DB_NAME")

	return &data
}
