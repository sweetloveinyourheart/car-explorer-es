package services

import (
	"book-explorer-es/internal/database"
	"book-explorer-es/internal/models"
	"fmt"

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
	db              *gorm.DB
	producerService IProducerService
}

func NewCarModelService() ICarModelService {
	return &CarModelService{
		db:              database.PostgresDB,
		producerService: NewProducerService(),
	}
}

func (cs *CarModelService) CreateCarModel(carModel *models.CarModel) error {
	producer, _ := cs.producerService.GetProducer(carModel.Producer.ID)
	if producer == nil {
		err := fmt.Errorf("producer is not exist")
		return err
	}

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
	producer, err := cs.producerService.GetProducer(carModel.Producer.ID)
	if producer == nil && err == nil {
		err = fmt.Errorf("producer is not exist")
		return err
	}

	if err != nil {
		err = fmt.Errorf("update existed car failed with err %e", err)
		return err
	}

	err = cs.db.Model(carModel).Updates(carModel).Error
	if err != nil {
		err = fmt.Errorf("update existed car failed with err %e", err)
		return err
	}

	return err
}

func (cs *CarModelService) DeleteCarModel(id uint) error {
	carModel := &models.CarModel{
		Model: gorm.Model{
			ID: id,
		},
	}

	err := cs.db.Delete(carModel).Error
	if err != nil {
		err = fmt.Errorf("delete car failed")
	}

	return err
}
