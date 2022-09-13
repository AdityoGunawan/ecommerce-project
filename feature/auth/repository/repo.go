package repository

import (
	"ecommerce-project/feature/auth/entities"
	"ecommerce-project/feature/auth/service"
	"ecommerce-project/middlewares"
	"ecommerce-project/models"
	"errors"

	"gorm.io/gorm"
)

type Storage struct {
	query *gorm.DB
}

func New(db *gorm.DB) entities.RepositoryInterface {
	return &Storage{
		query: db,
	}
}

func (storage *Storage) InsertInto(core entities.CoreRegister) (int, error) {
	model := models.RegCoreToModel(core)
	tx := storage.query.Create(&model)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (Storage *Storage) Generate(email, password string) (string, error) {
	var datauser models.User
	tx := Storage.query.Where("email = ?", email).First(&datauser)
	pass := service.CheckPasswordHash(password, datauser.Password)
	if !pass {
		return "", errors.New("Password Tidak Sesuai")
	}

	if tx.Error != nil {
		return "", tx.Error
	}
	if tx.RowsAffected != 1 {
		return "", errors.New("Gagal Terdapat dua data")
	}

	token, err := middlewares.CreateToken(int(datauser.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}
