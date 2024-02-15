package services

import (
	"book-explorer-es/internal/database"
	"book-explorer-es/internal/models"

	"gorm.io/gorm"
)

type IProducerService interface {
	CreateProducer(producer *models.Producer) error
	GetProducer(id uint) (*models.Producer, error)
	GetAllProducers() ([]models.Producer, error)
	UpdateProducer(producer *models.Producer) error
	DeleteProducer(id uint) error
}

type ProducerService struct {
	db *gorm.DB
}

func NewProducerService() IProducerService {
	return &ProducerService{
		db: database.PostgresDB,
	}
}

func (ps *ProducerService) CreateProducer(producer *models.Producer) error {
	return ps.db.Create(producer).Error
}

func (ps *ProducerService) GetProducer(id uint) (*models.Producer, error) {
	producer := &models.Producer{}
	if err := ps.db.First(producer, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return producer, nil
}

func (ps *ProducerService) GetAllProducers() ([]models.Producer, error) {
	producers := []models.Producer{}
	if err := ps.db.Find(&producers).Error; err != nil {
		return nil, err
	}
	return producers, nil
}

func (ps *ProducerService) UpdateProducer(producer *models.Producer) error {
	return ps.db.Model(producer).Updates(producer).Error
}

func (ps *ProducerService) DeleteProducer(id uint) error {
	producer := &models.Producer{
		Model: gorm.Model{
			ID: id,
		},
	}
	return ps.db.Delete(producer).Error
}
