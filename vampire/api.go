// Package vampire
// @Time  : 2021/12/22 下午8:35
// @Author: Jtyoui@qq.com
// @note  : api接口
package vampire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jtyoui/my-date-with-a-vampire/vampire/model"
	"strings"
)

// Login 登陆
func Login(c *gin.Context) {
	token := c.Query("token")
	field := FindByField{
		Key:   "uuid",
		Value: token,
		Ship:  AND,
	}
	people, ok := FindCond[model.User]([]FindByField{field})
	if !ok {
		Err(LoginCode, c)
		return
	}
	Success(people[0], c)
}

// Logon 注册
func Logon(c *gin.Context) {
	var people model.User
	if err := c.ShouldBindJSON(&people); err != nil {
		Err(ParamCode, c)
		return
	}

	token := strings.ReplaceAll(uuid.New().String(), "-", "")
	people.UUid = token

	if Crate[model.User](people) {
		Err(RegisterCode, c)
		return
	}

	Success(token, c)
}

// VideoByType 根据电影类型来获取电影的列表
func VideoByType(c *gin.Context) {
	var vt VideoType
	if err := c.ShouldBindJSON(&vt); err != nil {
		Err(ParamCode, c)
		return
	}

	video, ok := FindCond[model.Video](vt.getFields())
	if !ok {
		Err(DataCode, c)
		return
	}
	Success(video, c)
}

// VideoPlay 获取播放列表
func VideoPlay(c *gin.Context) {
	videoId := c.Query("video_id")
	m := PlayList(videoId)
	if m == nil {
		Err(DataCode, c)
		return
	}
	Success(m, c)
}
