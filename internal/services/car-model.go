package services

import (
	"book-explorer-es/internal/database"
	"book-explorer-es/internal/models"

	"gorm.io/gorm"
)

type ICarModelService interface {
	CreateCarModel(carModel *models.CarModel) error
	GetCarModel(id uint) (*models.CarModel, error)
	GetAllCarModels() ([]models.CarModel, error)
	UpdateCarModel(carModel *models.CarModel) error
	DeleteCarModel(id uint) error
}

type CarModelService struct {
	db *gorm.DB
}

func NewCarModelService() ICarModelService {
	return &CarModelService{
		db: database.PostgresDB,
	}
}

func (cs *CarModelService) CreateCarModel(carModel *models.CarModel) error {
	return cs.db.Create(carModel).Error
}

func (cs *CarModelService) GetCarModel(id uint) (*models.CarModel, error) {
	carModel := &models.CarModel{}
	if err := cs.db.First(carModel, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return carModel, nil
}

func (cs *CarModelService) GetAllCarModels() ([]models.CarModel, error) {
	carModels := []models.CarModel{}
	if err := cs.db.Find(&carModels).Error; err != nil {
		return nil, err
	}
	return carModels, nil
}

func (cs *CarModelService) UpdateCarModel(carModel *models.CarModel) error {
	return cs.db.Model(carModel).Updates(carModel).Error
}

func (cs *CarModelService) DeleteCarModel(id uint) error {
	carModel := &models.CarModel{
		Model: gorm.Model{
			ID: id,
		},
	}
	return cs.db.Delete(carModel).Error
}
