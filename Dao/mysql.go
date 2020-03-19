package Dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
	"github.com/shao1f/app/model"
)

func Search(city string) []model.CityToId {
	ci := make([]model.CityToId, 0)
	db, err := gorm.Open("mysql", "root:Syf3344521jsy@tcp(rm-bp1kr713l8l4px1814o.mysql.rds.aliyuncs.com:3306)/city_id?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Errorf("db error", err)
		return ci
	}
	db.SingularTable(true) // 避免gorm使用表的复数形式
	err = db.Where("city like ?", fmt.Sprintf("%%%v%%", city)).Find(&ci).Error
	if err != nil {
		log.Errorf("db select err:%v", err)
		return ci
	}
	if db.RecordNotFound() {
		log.Error("dn not found")
		return ci
	}
	return ci
}
