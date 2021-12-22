// Package vampire
// @Time  : 2021/12/9 下午3:53
// @Author: Jtyoui@qq.com
// @note  : 数据库操作
package vampire

import (
	"github.com/jtyoui/my-date-with-a-vampire/vampire/model"
	"gorm.io/gorm/schema"
	"strconv"
)

type FieldShip uint

const (
	AND FieldShip = iota
	OR
	LIMIT
	SELECT
)

type FindByField struct {
	Key   string
	Value any
	Ship  FieldShip
}

type VideoType struct {
	Type int `json:"type"`
	Page int `json:"page"`
	Size int `json:"size"`
}

func (v *VideoType) getFields() []FindByField {
	field := []FindByField{
		{Key: "type", Value: v.Type, Ship: AND},
		{Key: strconv.Itoa(v.Page), Value: v.Size, Ship: LIMIT},
	}
	return field
}

func FindCond[T schema.Tabler](fields []FindByField) (t []T, ok bool) {
	db := GDb
	for _, field := range fields {
		switch field.Ship {
		case AND:
			db = db.Where(field.Key+" = ?", field.Value)
		case OR:
			db = db.Or(field.Key+" = ?", field.Value)
		case LIMIT:
			pre, _ := strconv.Atoi(field.Key)
			num := field.Value.(int)
			db = db.Offset(pre).Limit(num)
		case SELECT:
			db = db.Select(field.Key)
		}
	}
	if err := db.Find(&t).Error; err != nil {
		return
	}
	ok = true
	return
}

func Crate[T schema.Tabler](t T) (ok bool) {
	if err := GDb.Create(&t).Error; err != nil {
		return
	}
	ok = true
	return
}

func PlayList(videoId string) (m map[string]any) {
	var video model.Video
	if GDb.Where("id = ?", videoId).First(&video).Error != nil {
		return
	}
	var videos []model.Video
	if GDb.Where("type = ?", video.Type).Find(&videos).Error != nil {
		return
	}
	m = make(map[string]any)
	m["url"] = video.URL
	m["name"] = video.Name
	m["title"] = video.Title
	m["list"] = videos
	return
}
