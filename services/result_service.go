package services

import (
	"lottery/dao"
	"lottery/datasource"
	"lottery/models"
)

type ResultService interface {
	GetAll(page, size int) []models.LtResult
	CountAll() int64
	GetNewPrize(size int, giftIds []int) []models.LtResult
	SearchByGift(giftId, page, size int) []models.LtResult
	SearchByUser(uid, page, size int) []models.LtResult
	CountByGift(giftId int) int64
	CountByUser(uid int) int64
	Get(id int) *models.LtResult
	Delete(id int) error
	Update(user *models.LtResult, columns []string) error
	Create(user *models.LtResult) error
}

type resultService struct {
	dao *dao.ResultDao
}

func (r *resultService) GetAll(page, size int) []models.LtResult {
	return r.GetAll(page, size)
}

func (r *resultService) CountAll() int64 {
	return r.dao.CountAll()
}

func (r *resultService) GetNewPrize(size int, giftIds []int) []models.LtResult {
	return r.dao.GetNewPrice(size, giftIds)
}

func (r *resultService) SearchByGift(giftId, page, size int) []models.LtResult {
	return r.dao.SearchByGift(giftId, page, size)
}

func (r *resultService) SearchByUser(uid, page, size int) []models.LtResult {
	return r.dao.SearchByUser(uid, page, size)
}

func (r *resultService) CountByGift(giftId int) int64 {
	return r.dao.CountByGift(giftId)
}

func (r *resultService) CountByUser(uid int) int64 {
	return r.dao.CountByUser(uid)
}

func (r *resultService) Get(id int) *models.LtResult {
	return r.dao.Get(id)
}

func (r *resultService) Delete(id int) error {
	return r.dao.Delete(id)
}

func (r *resultService) Update(user *models.LtResult, columns []string) error {
	return r.dao.Update(user, columns)
}

func (r *resultService) Create(user *models.LtResult) error {
	return r.dao.Create(user)
}

func NewResultService() ResultService {
	return &resultService{
		dao: dao.NewResultDao(datasource.InstanceDbMaster()),
	}
}
