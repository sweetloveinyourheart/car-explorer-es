package services

import (
	"book-explorer-es/internal/database"
	"book-explorer-es/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type ICarService interface {
	CreateCar(car *models.Car) error
	GetCar(id uint) (*models.Car, error)
	GetAllCars() ([]models.Car, error)
	UpdateCar(car *models.Car) error
	DeleteCar(id uint) error
}

type CarService struct {
	db              *gorm.DB
	carModelService ICarModelService
	featureService  IFeatureService
}

func NewCarService() ICarService {
	return &CarService{
		db:              database.PostgresDB,
		carModelService: NewCarModelService(),
		featureService:  NewFeatureService(),
	}
}

func (cs *CarService) CreateCar(car *models.Car) error {
	// Get car model
	carModel, _ := cs.carModelService.GetCarModel(car.CarModel.ID)
	if carModel == nil {
		err := fmt.Errorf("car model is not exist")
		return err
	}

	// Get car feature
	carFeatures := make([]uint, len(car.Features))
	for i, feature := range car.Features {
		carFeatures[i] = feature.ID
	}
	features, _ := cs.featureService.GetFeatures(carFeatures)
	if len(features) == 0 {
		err := fmt.Errorf("no features found")
		return err
	}

	car.CarModelID = carModel.ID
	car.Features = features

	return cs.db.Create(car).Error
}

func (cs *CarService) GetCar(id uint) (*models.Car, error) {
	car := &models.Car{}
	if err := cs.db.First(car, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return car, nil
}

func (cs *CarService) GetAllCars() ([]models.Car, error) {
	cars := []models.Car{}
	if err := cs.db.Find(&cars).Error; err != nil {
		return nil, err
	}
	return cars, nil
}

func (cs *CarService) UpdateCar(car *models.Car) error {
	return cs.db.Model(car).Updates(car).Error
}

func (cs *CarService) DeleteCar(id uint) error {
	car := &models.Car{
		Model: gorm.Model{
			ID: id,
		},
	}
	return cs.db.Delete(car).Error
}
