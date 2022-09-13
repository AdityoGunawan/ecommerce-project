package repository

import (
	"ecommerce-project/feature/product/entities"
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

func (storage *Storage) InsertI(core entities.CoreProduct) (int, error) {
	model := models.CoreToModel(core)
	tx := storage.query.Create(&model)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (storage *Storage) SelectAll(page int) ([]entities.CoreProduct, error) {
	var data []models.Product
	var count = 4 * (page - 1)

	tx := storage.query.Limit(4).Offset(count).Find(&data)
	if tx.Error != nil {
		return []entities.CoreProduct{}, tx.Error
	}
	corelist := models.ProductToCoreList(data)
	return corelist, nil

}

func (storage *Storage) Delete(userid, deleteid int) (string, error) {
	tx := storage.query.Unscoped().Delete(&models.Product{}, "user_id = ? and id = ?", userid, deleteid)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return "Gagal Menghapus", tx.Error
	}

	return "Sukses Menghapus Barang", nil
}

func (storage *Storage) SelectMyProduct(id int) ([]entities.CoreProduct, error) {
	var data []models.Product
	tx := storage.query.Find(&data, "user_id = ?", id)
	if tx.Error != nil {
		return []entities.CoreProduct{}, tx.Error
	}
	corelist := models.ProductToCoreList(data)
	return corelist, nil

}
