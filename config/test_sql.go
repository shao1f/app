package main

import (
	"bufio"
	"github.com/jinzhu/gorm"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CityToId struct {
	City      string    `gorm:"city"`
	CityId    int       `gorm:"city_id"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

func main() {
	f, err := os.Open("/Users/shao/gopath/src/github.com/shao1f/app/config/config.toml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	db, err := gorm.Open("mysql", "root:Syf3344521jsy@tcp(rm-bp1kr713l8l4px1814o.mysql.rds.aliyuncs.com:3306)/city_id?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	if err != nil {
		log.Fatal("db error", err)
	}
	// ci := CityToId{}
	// if err := db.First(&ci).Error; err != nil {
	// 	log.Fatalf("db select err:%v", err)
	// }
	// if db.RecordNotFound() {
	// 	fmt.Println("find nil")
	// }
	// fmt.Println(ci)
	defer db.Close()
	for {
		line, err := rd.ReadString('\n') // 以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		line = strings.Replace(line, "\n", "", -1)
		res := strings.Split(line, "\t")
		// fmt.Println(res[0], res[1])
		city := res[0]
		cityId, _ := strconv.Atoi(res[1])
		ti := time.Now()
		cityInfo := CityToId{
			City:      city,
			CityId:    cityId,
			CreatedAt: ti,
			UpdatedAt: ti,
		}
		//_ = fmt.Sprintf("insert into city_to_id (`city`,`city_id`,`created_at`) values(`%s`,`%v`,`%v`);", city, cityId, ti)
		db.Create(cityInfo)
		//fmt.Println(sqlStr)

	}

}
