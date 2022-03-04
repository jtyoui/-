// Package list
// @Time  : 2022/3/4 下午2:06
// @Author: Jtyoui@qq.com
// @note  :
package list

import (
	"github.com/gin-gonic/gin"
	"github.com/jtyoui/gam"
	"github.com/jtyoui/gam/cdb"
	"github.com/jtyoui/my-date-with-a-vampire/internal/common"
)

type List struct {
	Type int `json:"type"`
	Page int `json:"page"`
	Size int `json:"size"`
}

func (*List) PostVideo(c *gin.Context, ls List) {
	fields := ls.getFields()
	operate := gam.NewOperate(cdb.FIND, VideoModel{}, fields...)
	if err := operate.Join(common.GDb).Error; err != nil {
		common.Err(common.DataCode, c)
		return
	}
	common.Success(operate.Data, c)
}

func (l *List) getFields() []cdb.Fielder {
	field := []cdb.Fielder{
		cdb.DefaultField("type", l.Type),
		cdb.NewField(l.Page, l.Size, cdb.AccurateMatch, cdb.LIMIT),
	}
	return field
}
