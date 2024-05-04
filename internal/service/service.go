package service

import (
	"gorm.io/gorm"
	"main/models"
)

type Service struct {
	Db *gorm.DB
}

func InitialSrv(db *gorm.DB) *Service {
	return &Service{Db: db}
}

func (s *Service) GetEvents() ([]models.Events, error) {
	var events []models.Events
	err := s.Db.Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *Service) DeleteEvent(id int64) error {
	return s.Db.Where("id = ?", id).Delete(&models.Events{}).Error
}

func (s *Service) UpdateEvents(events []models.Events) ([]models.Events, error) {
	err := s.Db.Save(events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}
