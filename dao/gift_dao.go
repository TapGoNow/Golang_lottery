package dao

import (
	"lottery/comm"
	"lottery/models"
	"xorm.io/xorm"
)

type GiftDao struct {
	engine *xorm.Engine
}

func NewGiftDao(engine *xorm.Engine) *GiftDao {
	return &GiftDao{
		engine: engine,
	}
}

func (d *GiftDao) Get(id int) *models.LtGift {
	data := &models.LtGift{Id: id}
	ok, err := d.engine.Get(data)

	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *GiftDao) GetAll() []models.LtGift {
	datalist := make([]models.LtGift, 0)
	err := d.engine.Asc("sys_status").Asc("displayorder").Find(&datalist)

	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *GiftDao) GetAllUse() []models.LtGift {
	now := comm.NowUnix()
	datalist := make([]models.LtGift, 0)

	err := d.engine.Cols("id", "title", "prize_num", "left_num", "prize_code",
		"prize_time", "img", "displayorder", "gtype", "gdata").
		Desc("gtype").Asc("displayorder").
		Where("prize_num >= ?", 0).
		Where("sys_status = ?", 0).
		Where("timee_begin <= ?", now).
		Where("time_end >= ?", now).
		Find(&datalist)

	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *GiftDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtGift{})

	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *GiftDao) Delete(id int) error {
	data := &models.LtGift{Id: id, SysStatus: 1}
	_, err := d.engine.ID(data.Id).Update(data)

	return err
}

func (d *GiftDao) Update(data *models.LtGift, columns []string) error {
	_, err := d.engine.ID(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *GiftDao) Create(data *models.LtGift) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *GiftDao) IncrLeftNum(id, num int) (int64, error) {
	r, err := d.engine.ID(id).
		Incr("left_num", num).
		Update(&models.LtGift{Id: id})
	return r, err
}

func (d *GiftDao) DecrLeftNum(id, num int) (int64, error) {
	r, err := d.engine.ID(id).Incr("left_num", num).Update(&models.LtGift{Id: id})

	return r, err
}
