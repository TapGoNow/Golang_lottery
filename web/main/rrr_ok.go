// 这里做一些简单的验证
package main

import (
	"fmt"
)

func main() {
	k := []int{1}

	//fmt.Println(k[0])
	fmt.Sprintf("hhh-%d", k[0])
}

type data struct {
	Id         int64
	Ip         string
	Blacktime  string
	SysCreated string
	SysUpdated string
}

func NewData() *data {
	return &data{
		Id:         123,
		Ip:         "123456789",
		Blacktime:  "bbb",
		SysCreated: "ssss",
		SysUpdated: "yyyyy",
	}
}

func setByCache() {
	data := NewData()
	// 集群模式，redis缓存
	key := fmt.Sprintf("info_blackip_%s", data.Ip)
	//rds := datasource.InstanceCache()
	// 数据更新到redis缓存
	params := []interface{}{key}
	params = append(params, "Ip", data.Ip)
	if data.Id > 0 {
		params = append(params, "Blacktime", data.Blacktime)
		params = append(params, "SysCreated", data.SysCreated)
		params = append(params, "SysUpdated", data.SysUpdated)
	}
}
