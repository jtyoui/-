// Package vampire
// @Time  : 2021/12/8 上午10:27
// @Author: Jtyoui@qq.com
// @note  : 登陆
package vampire

import "github.com/gin-gonic/gin"

type User struct {
	Token string `json:"token" binding:"required"`
}

func (l User) PostLogin(c *gin.Context) {

}
