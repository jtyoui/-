// Package My_date_with_a_vampire
// @Time  : 2021/12/5 下午9:13
// @Author: Jtyoui@qq.com
// @note  : 主函数
package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/golang-middleware/ginDist"
)

//go:embed web/dist
var efs embed.FS

func main() {
	router := gin.Default()
	_ = router.SetTrustedProxies(nil)
	router.Use(ginDist.Static(efs, map[string]string{
		"/":         "index.html",
		"/register": "index.html",
	},
	))
	if err := router.Run(":25461"); err != nil {
		panic(err)
	}
}
