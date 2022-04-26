package dao

import (
	"lottery/models"
	"xorm.io/xorm"
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{
		engine: engine,
	}
}

func (r *UserDao) Get(id int) *models.LtUser {
	data := &models.LtUser{Id: id}
	ok, err := r.engine.Get(data)

	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (r *UserDao) GetAll(page, size int) []models.LtUser {
	offset := (page - 1) * size
	datalist := make([]models.LtUser, 0)
	r.engine.Desc("id").Limit(size, offset).Find(&datalist)

	return datalist
}

func (r *UserDao) CountAll() int {
	num, err := r.engine.Count(&models.LtUser{})

	if err != nil {
		return 0
	} else {
		return int(num)
	}
}

func (r *UserDao) Delete(id int) error {
	data := &models.LtUser{Id: id}
	_, err := r.engine.ID(data.Id).Update(data)

	return err
}

func (r *UserDao) Update(data *models.LtUser, columns []string) error {
	_, err := r.engine.ID(data.Id).MustCols(columns...).Update(data)
	return err
}

func (r *UserDao) Create(data *models.LtUser) error {
	_, err := r.engine.Insert(data)
	return err
}
