package dao

import (
	"log"
	"lottery/models"
	"xorm.io/xorm"
)

type CodeDao struct {
	engine *xorm.Engine
}

func NewCodeDao(engine *xorm.Engine) *CodeDao {
	return &CodeDao{
		engine: engine,
	}
}

func (c *CodeDao) Get(id int) *models.LtCode {
	data := &models.LtCode{Id: id}
	ok, err := c.engine.Get(data)

	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (c *CodeDao) GetAll(page, size int) []models.LtCode {
	offset := (page - 1) * size
	datalist := make([]models.LtCode, 0)
	err := c.engine.
		Desc("id").
		Limit(size, offset).
		Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (c *CodeDao) CountAll() int64 {
	num, err := c.engine.
		Count(&models.LtCode{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (c *CodeDao) Delete(id int) error {
	data := &models.LtCode{Id: id, SysStatus: 1}
	_, err := c.engine.ID(data.Id).Update(data)

	return err
}

func (c *CodeDao) Update(data *models.LtCode, columns []string) error {
	_, err := c.engine.ID(data.Id).MustCols(columns...).Update(data)
	return err
}

func (c *CodeDao) Create(data *models.LtCode) error {
	_, err := c.engine.Insert(data)
	return err
}

//找到下一个可用的最小优惠券
func (c *CodeDao) NextUsingCode(giftId, codeId int) *models.LtCode {
	datalist := make([]models.LtCode, 0)
	err := c.engine.Where("gift_id = ?", giftId).
		Where("sys_status = ?", 0).
		Where("id > ?", codeId).
		Asc("id").Limit(1).
		Find(&datalist)
	if err != nil || len(datalist) < 1 {
		return nil
	} else {
		return &datalist[0]
	}
}

// 根据唯一的code来更新
func (c *CodeDao) UpdateByCode(data *models.LtCode, columns []string) error {
	_, err := c.engine.Where("code=?", data.Code).
		MustCols(columns...).Update(data)
	return err
}

func (c *CodeDao) CountByGift(id int) int64 {
	num, err := c.engine.Where("GiftId=?", id).Count(&models.LtCode{})

	if err != nil {
		return 0
	} else {
		return num
	}
}

func (c *CodeDao) Search(giftId int) []models.LtCode {
	datalist := make([]models.LtCode, 0)
	err := c.engine.Where("GiftId=?", giftId).Find(&datalist)

	if err != nil {
		log.Printf("code_dao Search() have error : %s\n", err)
		return datalist
	} else {
		return datalist
	}

}
