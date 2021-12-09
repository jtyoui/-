// Package vampire
// @Time  : 2021/12/8 上午10:16
// @Author: Jtyoui@qq.com
// @note  : model数据库表结构
package vampire

import "time"

type People struct {
	UUid      string    `json:"-"           gorm:"comment:僵尸码;column:uuid;primaryKey"`
	CreatedAt time.Time `json:"create_time" gorm:"comment:创建日期"                      binding:"isdefault"`
	Name      string    `json:"name"        gorm:"comment:昵称;column:name"             binding:"min=2,max=10"`
	BirthDate time.Time `json:"birth_date"  gorm:"comment:出生年月;column:birth_date"    binding:"required"`
	Gender    string    `json:"gender"      gorm:"comment:性别;column:gender"           binding:"oneof=男 女"`
}

func (People) TableName() string {
	return "SysPeople"
}
