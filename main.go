// Package My_date_with_a_vampire
// @Time  : 2021/12/5 下午9:13
// @Author: Jtyoui@qq.com
// @note  : 主函数
package main

import (
	"embed"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-middleware/ginDist"
	"github.com/golang-middleware/ginRouter"
	"github.com/jtyoui/my-date-with-a-vampire/vampire"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"path"
	"strings"
	"time"
)

//go:embed web/dist
var efs embed.FS

//go:embed config.toml
var Config string

// InitConfig 初始化配置文件
func InitConfig() {
	paths := flag.String("f", "", "配置文件的路径")
	flag.Parse()
	var err error
	if *paths != "" {
		_, err = toml.DecodeFile(*paths, &vampire.GConfig)
	} else {
		_, err = toml.Decode(Config, &vampire.GConfig)
	}
	if err != nil {
		panic("配置文件加载失败")
	}
}

func InitDatabase() {
	db, err := gorm.Open(vampire.InitMysql(), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名单数
		},
	})
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

	if err = db.AutoMigrate(&vampire.People{}); err != nil {
		panic(err)
	}
	vampire.GDb = db
}

func InitRun() {
	if vampire.GConfig.System.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.Use(cors.Default())
	_ = router.SetTrustedProxies(nil)

	api := vampire.GConfig.System.PrefixRouter

	dist := ginDist.AutoReplace{
		F: func(url string) string {
			suffix := path.Ext(url)
			if !strings.HasPrefix(url, "/"+api) && suffix == "" {
				return ginDist.DefaultRouterHtml
			}
			return url
		},
	}

	// 中间间
	router.Use(dist.Default(efs))

	public := router.Group(api)
	publicRouter := ginRouter.GinRouter{Router: public}
	publicRouter.AutoRouter(
		&vampire.Index{},
		&vampire.Register{},
	)
	if err := router.Run(fmt.Sprintf(":%d", vampire.GConfig.System.Port)); err != nil {
		panic(err)
	}
}

func main() {
	InitConfig()
	InitDatabase()
	InitRun()
}
