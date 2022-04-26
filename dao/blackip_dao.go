package dao

import (
	"lottery/models"
	"xorm.io/xorm"
)

type BlackIpDao struct {
	engine *xorm.Engine
}

func NewBlackIpDao(engine *xorm.Engine) *BlackIpDao {
	return &BlackIpDao{
		engine: engine,
	}
}

func (d *BlackIpDao) Get(id int) *models.LtBlackip {
	data := &models.LtBlackip{Id: id}
	ok, err := d.engine.Get(data)

	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *BlackIpDao) GetByIp(Ip string) *models.LtBlackip {
	data := &models.LtBlackip{Ip: Ip}
	ok, err := d.engine.Get(data)

	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *BlackIpDao) GetAll(page, size int) []models.LtBlackip {
	offset := (page - 1) * size
	datalist := make([]models.LtBlackip, 0)
	d.engine.Desc("id").Limit(size, offset).Find(&datalist)

	return datalist
}

func (d *BlackIpDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtBlackip{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *BlackIpDao) Search(ip string) []models.LtBlackip {
	datalist := make([]models.LtBlackip, 0)
	err := d.engine.
		Where("ip=?", ip).
		Desc("id").
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *BlackIpDao) Update(data *models.LtBlackip, columns []string) error {
	_, err := d.engine.ID(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *BlackIpDao) Create(data *models.LtBlackip) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *BlackIpDao) Delete(id int) error {
	data := &models.LtBlackip{Id: id}
	_, err := d.engine.ID(data.Id).Delete(data)

	return err
}
