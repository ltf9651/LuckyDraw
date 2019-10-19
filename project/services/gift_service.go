package services

import (
	"LuckyDraw/project/dao"
	"LuckyDraw/project/models"
)

type GiftService interface {
	GetAll() []models.Gift
	CountAll() int64
	Get(id int) *models.Gift
	Delete(id int) error
	Update(data *models.Gift, columns []string) error
	Create(data *models.Gift) error
}

type giftService struct {
	dao *dao.GiftDao
}

func NewGiftService() GiftService {
	return &giftService{
		dao: dao.NewGiftDao(nil),
	}
}

func (s *giftService) GetAll() []models.Gift {
	return s.dao.GetAll()
}
func (s *giftService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *giftService) Get(id int) *models.Gift {
	return s.dao.Get(id)

}
func (s *giftService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *giftService) Update(data *models.Gift, columns []string) error {
	return s.dao.Update(data, columns)

}
func (s *giftService) Create(data *models.Gift) error {
	return s.dao.Create(data)
}
