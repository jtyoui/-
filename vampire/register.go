// Package vampire
// @Time  : 2021/12/8 上午10:28
// @Author: Jtyoui@qq.com
// @note  : 注册
package vampire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

type Register struct{}

func (r *Register) PostLogon(c *gin.Context) {
	var people People
	if err := c.ShouldBindJSON(&people); err != nil {
		Err(RegisterCode, c)
		return
	}
	token := people.submit()
	if token == "" {
		Err(RegisterCode, c)
		return
	}
	Success(token, c)
}

func (p *People) submit() (token string) {
	u := uuid.New().String()
	token = strings.ReplaceAll(u, "-", "")
	p.UUid = token
	if err := GDb.Create(p).Error; err != nil {
		return ""
	}
	return
}
