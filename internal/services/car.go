package services

import (
	"book-explorer-es/internal/database"
	"book-explorer-es/internal/models"

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
	db *gorm.DB
}

func NewCarService() ICarService {
	return &CarService{
		db: database.PostgresDB,
	}
}

func (cs *CarService) CreateCar(car *models.Car) error {
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
