package repository

import (
	"ecommerce-project/feature/auth/service"
	"ecommerce-project/feature/user/entities"
	"ecommerce-project/models"

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

func (storage *Storage) Select(id int) (entities.CoreUser, error) {
	var datauser models.User
	tx := storage.query.Where("id", id).First(&datauser)
	if tx.Error != nil {
		return entities.CoreUser{}, tx.Error
	}
	core := models.ModelToCore(datauser)
	return core, nil
}

func (storage *Storage) Update(data entities.CoreUser, id int) (string, error) {
	var user models.User
	user = models.CoreUserToModel(data)
	if user.Password != "" {
		hash, _ := service.HashPassword(user.Password)
		user.Password = hash
	}
	tx := storage.query.Model(&models.User{}).Where("id = ?", id).Updates(user)
	if tx.Error != nil || tx.RowsAffected != 1 {
		return "Melakukan Update", tx.Error
	}

	return "Sukses Update", nil
}

func (storage *Storage) Delete(userId int) (string, error) {
	tx := storage.query.Where("id = ?", userId).Delete(&models.User{})
	if tx.Error != nil || tx.RowsAffected != 1 {
		return "Gagal Menggapus", tx.Error
	}

	return "Data Dihapus", nil
}
