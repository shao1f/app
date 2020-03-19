package service

import (
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/shao1f/app/Dao"
	"github.com/shao1f/app/manager"
	"github.com/shao1f/app/model"
)

func Search(city string) *model.WeatherInfo {
	// result := make([]model.City, 0)
	res := Dao.Search(city)
	if len(res) == 0 {
		log.Warn("search nil")
		return nil
	}
	// tempMap := make(map[string]int)
	// // 会搜索出来一些相同数据，例如西安市->3005，西安市->3001 相同名称但是城市id不同，进行一次过滤
	// if len(res) > 1 {
	// 	for _, v := range res {
	// 		_, ok := tempMap[v.City]
	// 		if !ok {
	// 			tempMap[v.City] = v.CityId
	// 			city := model.City{
	// 				City:   v.City,
	// 				CityId: v.CityId,
	// 			}
	// 			result = append(result, city)
	// 		}
	// 	}
	// } else {
	// 	city := model.City{
	// 		City:   res[0].City,
	// 		CityId: res[0].CityId,
	// 	}
	// 	result = append(result, city)
	// }
	cityIdStr := fmt.Sprintf("%v", res[0].CityId)
	resp, err := manager.GetGodeWeather(cityIdStr)
	if err != nil {
		log.Errorf("get gaode weather err,city=%v", res[0])
		return nil
	}

	return resp
}
