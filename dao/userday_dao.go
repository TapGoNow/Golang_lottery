package dao

import (
	"lottery/models"
	"xorm.io/xorm"
)

type UserDayDao struct {
	engine *xorm.Engine
}

func NewUserDayDao(engine *xorm.Engine) *UserDayDao {
	return &UserDayDao{
		engine: engine,
	}
}

func (r *UserDayDao) Get(id int) *models.LtUserday {
	data := &models.LtUserday{Id: id}
	ok, err := r.engine.Get(data)

	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (r *UserDayDao) GetAll(page, size int) []models.LtUserday {
	offset := (page - 1) * size
	datalist := make([]models.LtUserday, 0)
	r.engine.Desc("id").Limit(size, offset).Find(&datalist)

	return datalist
}

func (r *UserDayDao) CountAll() int64 {
	num, err := r.engine.Count(&models.LtUserday{})

	if err != nil {
		return 0
	} else {
		return num
	}
}

func (r *UserDayDao) Delete(id int) error {
	data := &models.LtUserday{Id: id}
	_, err := r.engine.ID(data.Id).Update(data)

	return err
}

func (r *UserDayDao) Update(data *models.LtUserday, columns []string) error {
	_, err := r.engine.ID(data.Id).MustCols(columns...).Update(data)
	return err
}

func (r *UserDayDao) Create(data *models.LtUserday) error {
	_, err := r.engine.Insert(data)
	return err
}

func (r *UserDayDao) Search(uid int, day int) []models.LtUserday {
	datalist := make([]models.LtUserday, 0)
	err := r.engine.
		Where("uid=?", uid).
		Where("day=?", day).
		Desc("id").
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}

}

func (r *UserDayDao) Count(uid int, day int) int {
	info := &models.LtUserday{}
	ok, err := r.engine.
		Where("uid=?", uid).
		Where("day=?", day).
		Get(info)
	if !ok || err != nil {
		return 0
	} else {
		return info.Num
	}
}
