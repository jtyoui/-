// Package vampire
// @Time  : 2021/12/8 上午10:27
// @Author: Jtyoui@qq.com
// @note  : 登陆
package vampire

import "github.com/gin-gonic/gin"

type Index struct {
	Token string `json:"token" form:"token" binding:"required"`
}

func (l *Index) GetLogin(c *gin.Context) {
	if err := c.ShouldBindQuery(l); err != nil {
		Err(LoginCode, c)
		return
	}
	people, ok := login[People]("uuid", l.Token)
	if !ok {
		Err(LoginCode, c)
		return
	}
	Success(people, c)
}

func (l *Index) login() (people *People, ok bool) {
	if err := GDb.Where("uuid = ?", l.Token).First(&people).Error; err != nil {
		return
	}
	ok = true
	return
}
