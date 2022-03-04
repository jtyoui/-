// Package login
// @Time  : 2022/3/3 下午5:49
// @Author: Jtyoui@qq.com
// @note  :
package login

import (
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	gorm.Model
	UUid      string    `json:"uuid"        gorm:"comment:僵尸码;column:uuid"`
	Name      string    `json:"name"        gorm:"comment:昵称;column:name"           validate:"min=2,max=10"`
	BirthDate time.Time `json:"birth_date"  gorm:"comment:出生年月;column:birth_date"  validate:"required"`
	Gender    string    `json:"gender"      gorm:"comment:性别;column:gender"         validate:"oneof=男 女"`
}

func (UserModel) TableName() string {
	return "sys_user"
}
