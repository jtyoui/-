// Package play
// @Time  : 2022/3/4 下午2:07
// @Author: Jtyoui@qq.com
// @note  :
package play

import (
	"github.com/gin-gonic/gin"
	"github.com/jtyoui/gam"
	"github.com/jtyoui/gam/cdb"
	"github.com/jtyoui/my-date-with-a-vampire/internal/common"
)

type Play struct {
	Title string       `json:"title"`
	Name  string       `json:"name"`
	URL   string       `json:"url"`
	List  []VideoModel `json:"list"`
}

func (*Play) GetId(c *gin.Context, videoId int) {
	operate := gam.NewOperate(cdb.FIRST, VideoModel{}, cdb.DefaultField("id", videoId))
	if err := operate.Join(common.GDb).Error; err != nil {
		common.Err(common.DataCode, c)
		return
	}
	video := operate.Data.([]VideoModel)[0]

	operate = gam.NewOperate(cdb.FIND, VideoModel{}, cdb.DefaultField("type", video.Type))
	if err := operate.Join(common.GDb).Error; err != nil {
		common.Err(common.DataCode, c)
		return
	}

	videos := operate.Data.([]VideoModel)

	p := &Play{
		Title: video.Title,
		Name:  video.Name,
		URL:   video.URL,
		List:  videos,
	}

	common.Success(p, c)
}
