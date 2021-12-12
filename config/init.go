// Package config
// @Time  : 2021/10/12 上午9:27
// @Author: Jtyoui@qq.com
// @note  : 所有配置类
package config

type Config struct {
	System     System
	Mysql      Mysql
	Jwt        Jwt
	VerifyCode VerifyCode
	Sqlite     Sqlite
	Email      Email
}
