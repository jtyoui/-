// Package config
// @Time  : 2021/10/12 上午9:27
// @Author: Jtyoui@qq.com
// @note  : 所有配置类
package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	"io/fs"
)

type Config struct {
	System     System
	Mysql      Mysql
	Jwt        Jwt
	VerifyCode VerifyCode
	Sqlite     Sqlite
	Email      Email
}

var GConfig Config

func InitConfig(efs fs.FS) {
	paths := flag.String("f", "", "配置文件的路径")
	flag.Parse()
	var err error
	if *paths != "" {
		_, err = toml.DecodeFile(*paths, &GConfig)
	} else {
		_, err = toml.DecodeFS(efs, "config.toml", &GConfig)
	}
	if err != nil {
		panic("配置文件加载失败")
	}
}
