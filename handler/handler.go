package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shao1f/app/manager"
	"github.com/shao1f/app/model"
	"github.com/shao1f/app/service"
	"log"
)

func StartPage(c *gin.Context) {
	var person model.Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.JSON(200, person)
}

func GetWeather(c *gin.Context) {
	var locaInfo model.LocationInfo
	if err := c.ShouldBind(&locaInfo); err != nil {
		c.JSON(200, model.Response{
			ErrCode:  499,
			ErrorMsg: "参数错误",
		})
		return
	}
	log.Println(locaInfo)
	loca, err := manager.GetGaodeLoca(locaInfo)
	if err != nil {
		c.JSON(200, model.Response{
			ErrCode:  500,
			ErrorMsg: "内部错误",
		})
		return
	}
	cityCode, err := manager.GetGaodeCity(loca)
	if err != nil {
		c.JSON(200, model.Response{
			ErrCode:  500,
			ErrorMsg: "内部错误",
		})
		return
	}
	weather, err := manager.GetGodeWeather(cityCode)
	if err != nil {
		c.JSON(200, model.Response{
			ErrCode:  500,
			ErrorMsg: "内部错误",
		})
		return
	}
	resp := model.WeatherResult{
		Response: model.Response{
			ErrCode:  0,
			ErrorMsg: "操作成功",
		},
		Data: weather,
	}
	c.JSON(200, resp)
}

func Search(c *gin.Context) {
	var search model.SearchReq
	if err := c.ShouldBindQuery(&search); err != nil {
		c.JSON(200, model.Response{
			ErrCode:  499,
			ErrorMsg: "参数错误",
		})
		return
	}
	data := service.Search(search.City)
	c.JSON(200, model.SearchResp{
		Response: model.Response{
			ErrCode:  0,
			ErrorMsg: "操作成功",
		},
		Data: data,
	})
}
