package services

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"lottery/comm"
	"lottery/dao"
	"lottery/datasource"
	"lottery/models"
	"sync"
)

// IP信息，可以缓存(本地或者redis)，有更新的时候，再根据具体情况更新缓存
var cachedBlackIpList = make(map[string]*models.LtBlackip)
var cachedBlackIpLock = sync.Mutex{}

type BlackIpService interface {
	GetAll(page, size int) []models.LtBlackip
	CountAll() int64
	Search(ip string) []models.LtBlackip
	Get(id int) *models.LtBlackip
	//Delete(id int) error
	Update(user *models.LtBlackip, columns []string) error
	Create(user *models.LtBlackip) error
	GetByIp(ip string) *models.LtBlackip
}

type blackIpService struct {
	dao *dao.BlackIpDao
}

func (b *blackIpService) GetAll(page, size int) []models.LtBlackip {
	return b.dao.GetAll(page, size)
}

func (b *blackIpService) CountAll() int64 {
	return b.dao.CountAll()
}

func (b *blackIpService) Search(ip string) []models.LtBlackip {
	return b.dao.Search(ip)
}

func (b *blackIpService) Get(id int) *models.LtBlackip {
	return b.dao.Get(id)
}

func (b *blackIpService) Update(data *models.LtBlackip, columns []string) error {
	b.updateByCache(data, columns)
	return b.dao.Update(data, columns)
}

func (b *blackIpService) Create(data *models.LtBlackip) error {
	return b.dao.Create(data)
}

func (b *blackIpService) GetByIp(ip string) *models.LtBlackip {
	//先从缓存读取数据
	data := b.getByCache(ip)
	if data == nil || data.Ip == "" {
		//再从数据库中读取数据
		data = b.dao.GetByIp(ip)
		if data == nil || data.Ip == "" {
			data = &models.LtBlackip{Ip: ip}
		}
		b.setByCache(data)
	}
	return data
}

func (b *blackIpService) getByCache(ip string) *models.LtBlackip {
	// 集群模式，redis缓存
	key := fmt.Sprintf("info_blackip_%s", ip)
	rds := datasource.InstanceCache()
	dataMap, err := redis.StringMap(rds.Do("HGETALL", key))
	if err != nil {
		log.Println("blackip_service.getByCache HGETALL key=", key, ", error=", err)
		return nil
	}
	dataIp := comm.GetStringFromStringMap(dataMap, "Ip", "")
	if dataIp == "" {
		return nil
	}
	data := &models.LtBlackip{
		Id:         int(comm.GetInt64FromStringMap(dataMap, "Id", 0)),
		Ip:         dataIp,
		Blacktime:  int(comm.GetInt64FromStringMap(dataMap, "Blacktime", 0)),
		SysCreated: int(comm.GetInt64FromStringMap(dataMap, "SysCreated", 0)),
		SysUpdated: int(comm.GetInt64FromStringMap(dataMap, "SysUpdated", 0)),
	}
	return data
}

func (b *blackIpService) setByCache(data *models.LtBlackip) {
	if data == nil || data.Ip == "" {
		return
	}
	// 集群模式，redis缓存
	key := fmt.Sprintf("info_blackip_%s", data.Ip)
	rds := datasource.InstanceCache()
	// 数据更新到redis缓存
	params := []interface{}{key}
	params = append(params, "Ip", data.Ip)
	if data.Id > 0 {
		params = append(params, "Blacktime", data.Blacktime)
		params = append(params, "SysCreated", data.SysCreated)
		params = append(params, "SysUpdated", data.SysUpdated)
	}
	_, err := rds.Do("HMSET", params...)
	if err != nil {
		log.Println("blackip_service.setByCache HMSET params=", params, ", error=", err)
	}
}

//数据更新，直接清空缓存数据
func (b *blackIpService) updateByCache(data *models.LtBlackip, columns []string) {
	if data == nil || data.Ip == "" {
		return
	}
	// 集群模式，redis缓存
	key := fmt.Sprintf("info_blackip_%s", data.Ip)
	rds := datasource.InstanceCache()
	// 删除redis中的缓存
	rds.Do("DEL", key)
}

func NewBlackIpService() BlackIpService {
	return &blackIpService{
		dao: dao.NewBlackIpDao(datasource.InstanceDbMaster()),
	}
}
