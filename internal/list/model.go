// Package list
// @Time  : 2022/3/4 下午2:51
// @Author: Jtyoui@qq.com
// @note  :
package list

import "gorm.io/gorm"

type VideoModel struct {
	gorm.Model
	Title string `gorm:"comment:电影标题;column:title" json:"title"`
	Name  string `gorm:"comment:电影名称;column:name"  json:"name"`
	Type  int    `gorm:"comment:电影类型;column:type"  json:"-"`
	URL   string `gorm:"comment:电影URL;column:url"   json:"url"`
	Time  string `gorm:"comment:电影时长;column:time"  json:"time"`
}

func (VideoModel) TableName() string {
	return "sys_video"
}
