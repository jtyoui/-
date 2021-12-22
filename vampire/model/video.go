// Package model
// @Time  : 2021/12/22 下午9:00
// @Author: Jtyoui@qq.com
// @note  : 电影资源表
package model

type Video struct {
	ID    uint   `gorm:"primaryKey"                   json:"id"`
	Title string `gorm:"comment:电影标题;column:title" json:"title"`
	Name  string `gorm:"comment:电影名称;column:name"  json:"name"`
	Type  int    `gorm:"comment:电影类型;column:type"  json:"-"`
	URL   string `gorm:"comment:电影URL;column:url"   json:"url"`
	Time  string `gorm:"comment:电影时长;column:time"   json:"time"`
}

func (Video) TableName() string {
	return "SysVideo"
}
