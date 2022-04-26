package dao

import (
	"lottery/models"
	"xorm.io/xorm"
)

type ResultDao struct {
	engine *xorm.Engine
}

func NewResultDao(engine *xorm.Engine) *ResultDao {
	return &ResultDao{
		engine: engine,
	}
}

func (r *ResultDao) Get(id int) *models.LtResult {
	data := &models.LtResult{Id: id}
	ok, err := r.engine.Get(data)

	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (r *ResultDao) GetAll(page, size int) []models.LtResult {
	offset := (page - 1) * size
	datalist := make([]models.LtResult, 0)
	r.engine.Desc("id").Limit(size, offset).Find(&datalist)

	return datalist
}

func (r *ResultDao) CountAll() int64 {
	num, err := r.engine.Count(&models.LtResult{})

	if err != nil {
		return 0
	} else {
		return num
	}
}

func (r *ResultDao) Delete(id int) error {
	data := &models.LtResult{Id: id, SysStatus: 1}
	_, err := r.engine.ID(data.Id).Update(data)

	return err
}

func (r *ResultDao) Update(data *models.LtResult, columns []string) error {
	_, err := r.engine.ID(data.Id).MustCols(columns...).Update(data)
	return err
}

func (r *ResultDao) Create(data *models.LtResult) error {
	_, err := r.engine.Insert(data)
	return err
}

func (r *ResultDao) GetNewPrice(size int, ids []int) []models.LtResult {
	datalist := make([]models.LtResult, 0)
	err := r.engine.
		In("gift_id", ids).
		Desc("id").
		Limit(size).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (r *ResultDao) SearchByGift(id int, page int, size int) []models.LtResult {
	offset := (page - 1) * size
	datalist := make([]models.LtResult, 0)
	err := r.engine.
		Where("gift_id=?", id).
		Desc("id").
		Limit(size, offset).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (r *ResultDao) SearchByUser(uid int, page int, size int) []models.LtResult {
	offset := (page - 1) * size
	datalist := make([]models.LtResult, 0)
	err := r.engine.
		Where("uid=?", uid).
		Desc("id").
		Limit(size, offset).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (r *ResultDao) CountByGift(id int) int64 {
	num, err := r.engine.
		Where("gift_id = ?", id).
		Count(&models.LtResult{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (r *ResultDao) CountByUser(uid int) int64 {
	num, err := r.engine.Where("uid=?", uid).Count(&models.LtResult{})
	if err != nil {
		return 0
	} else {
		return num
	}
}
