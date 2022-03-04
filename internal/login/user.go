// Package login
// @Time  : 2022/3/3 下午5:49
// @Author: Jtyoui@qq.com
// @note  :
package login

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jtyoui/gam"
	"github.com/jtyoui/gam/cdb"
	"github.com/jtyoui/my-date-with-a-vampire/internal/common"
	"strings"
)

type User struct{}

func (*User) GetLogin(c *gin.Context, token string) {
	operate := gam.NewOperate(cdb.FIRST, UserModel{}, gam.DefaultField("uuid", token))
	if err := operate.Join(common.GDb).Error; err != nil {
		common.Err(common.LoginCode, c)
		return
	}
	users := operate.Data.([]UserModel)
	common.Success(users[0], c)
}

func (*User) PostRegister(c *gin.Context, user UserModel) {
	token := strings.ReplaceAll(uuid.New().String(), "-", "")
	user.UUid = token

	operate := gam.NewOperate(cdb.CREATE, &user)
	if err := operate.Join(common.GDb).Error; err != nil {
		common.Err(common.RegisterCode, c)
		return
	}
	common.Success(token, c)
}
