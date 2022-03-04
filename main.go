// Package My_date_with_a_vampire
// @Time  : 2021/12/5 下午9:13
// @Author: Jtyoui@qq.com
// @note  : 主函数
package main

import (
	"embed"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jtyoui/gam"
	"github.com/jtyoui/gam/web"
	"github.com/jtyoui/my-date-with-a-vampire/config"
	"github.com/jtyoui/my-date-with-a-vampire/internal/common"
	"github.com/jtyoui/my-date-with-a-vampire/internal/list"
	"github.com/jtyoui/my-date-with-a-vampire/internal/login"
	"github.com/jtyoui/my-date-with-a-vampire/internal/play"
	"github.com/jtyoui/my-date-with-a-vampire/pkg"
	"gorm.io/gorm"
	"path"
	"strings"
	"time"
)

//go:embed web/dist config.toml
var efs embed.FS

func InitDatabase() {
	db, err := gorm.Open(common.InitMysql(), &gorm.Config{})
	if err != nil {
		panic("链接数据库失败")
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("开辟线程池失败")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err = db.AutoMigrate(&login.UserModel{}, &list.VideoModel{}); err != nil {
		panic(err)
	}
	common.GDb = db
}

func InitRun() {
	if config.GConfig.System.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	_ = router.SetTrustedProxies(nil)

	api := config.GConfig.System.PrefixRouter

	dist := gam.NewDist()
	dist.CustomRuleFunc(func(url string) string {
		suffix := path.Ext(url)
		if !strings.HasPrefix(url, "/"+api) && suffix == "" {
			return web.DefaultRouterHtml
		}
		return url
	})
	router.Use(cors.Default())   // 跨越
	router.Use(dist.LoadFs(efs)) // 中间间

	public := router.Group(api)
	g := gam.NewGinRouter(public, nil)

	g.AutoRouter(&login.User{}, &list.List{}, &play.Play{})

	fmt.Printf("############### 请访问: http://%s:%d #########", pkg.GetLocalBoundIP(), config.GConfig.System.Port)
	if err := router.Run(fmt.Sprintf("0.0.0.0:%d", config.GConfig.System.Port)); err != nil {
		panic(err)
	}
}

func main() {
	config.InitConfig(efs)
	InitDatabase()
	InitRun()
}
