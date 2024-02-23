package services

import (
	"book-explorer-es/internal/database"
	"book-explorer-es/internal/models"

	"gorm.io/gorm"
)

type IFeatureService interface {
	CreateFeature(feature *models.Feature) error
	GetFeature(id uint) (*models.Feature, error)
	GetFeatures(ids []uint) ([]models.Feature, error)
	GetAllFeatures() ([]models.Feature, error)
	UpdateFeature(feature *models.Feature) error
	DeleteFeature(id uint) error
}

type FeatureService struct {
	db *gorm.DB
}

func NewFeatureService() IFeatureService {
	return &FeatureService{
		db: database.PostgresDB,
	}
}

func (ps *FeatureService) CreateFeature(feature *models.Feature) error {
	return ps.db.Create(feature).Error
}

func (ps *FeatureService) GetFeature(id uint) (*models.Feature, error) {
	feature := &models.Feature{}
	if err := ps.db.First(feature, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return feature, nil
}

func (ps *FeatureService) GetAllFeatures() ([]models.Feature, error) {
	features := []models.Feature{}
	if err := ps.db.Find(&features).Error; err != nil {
		return nil, err
	}
	return features, nil
}

func (ps *FeatureService) GetFeatures(ids []uint) ([]models.Feature, error) {
	features := []models.Feature{}
	if err := ps.db.Find(&features, "id IN (?)", ids).Error; err != nil {
		return nil, err
	}
	return features, nil
}

func (ps *FeatureService) UpdateFeature(feature *models.Feature) error {
	return ps.db.Model(feature).Updates(feature).Error
}

func (ps *FeatureService) DeleteFeature(id uint) error {
	feature := &models.Feature{
		Model: gorm.Model{
			ID: id,
		},
	}
	return ps.db.Delete(feature).Error
}
