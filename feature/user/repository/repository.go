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

func (storage *Storage) UpdateImage(alamat string, id int) (string, error) {
	var user models.User
	user.Foto = alamat
	tx := storage.query.Model(&models.User{}).Where("id = ?", id).Updates(user)
	if tx.Error != nil {
		return "Gagal Upload Foto", tx.Error
	}

	return "Sukses Upload Foto", nil
}

func (storage *Storage) UpdateUserName(core entities.CoreUser, id int) (string, error) {
	var user models.User
	user.Username = core.Username
	tx := storage.query.Model(&models.User{}).Where("id = ?", id).Updates(user)
	if tx.Error != nil {
		return "Gagal Mengupdate", tx.Error
	}

	return "Sukses Update Username", nil
}

func (storage *Storage) UpdatePassword(core entities.CoreUser, id int) (string, error) {
	var user models.User
	hash, _ := service.HashPassword(core.Password)
	user.Password = hash
	tx := storage.query.Model(&models.User{}).Where("id = ?", id).Updates(user)
	if tx.Error != nil {
		return "Gagal Mengupdate", tx.Error
	}
	return "Sukses Update Password", nil
}

func (storage *Storage) UpdateEmail(core entities.CoreUser, id int) (string, error) {
	var user models.User
	user.Email = core.Email
	tx := storage.query.Model(&models.User{}).Where("id = ?", id).Updates(user)
	if tx.Error != nil {
		return "Gagal Mengupdate", tx.Error
	}
	return "Sukses Update Email", nil
}

func (storage *Storage) UpdateName(core entities.CoreUser, id int) (string, error) {
	var user models.User
	user.Name = core.Name
	tx := storage.query.Model(&models.User{}).Where("id = ?", id).Updates(user)
	if tx.Error != nil {
		return "Gagal Mengupdate", tx.Error
	}
	return "Sukses Update Name", nil
}

func (storage *Storage) Delete(userId int) (string, error) {
	tx := storage.query.Unscoped().Delete(&models.User{}, "id = ?", userId)
	if tx.Error != nil || tx.RowsAffected != 1 {
		return "Gagal Menggapus", tx.Error
	}

	return "Data Dihapus", nil
}
