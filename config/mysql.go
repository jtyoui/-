// Package config
// @Time  : 2021/10/12 上午9:01
// @Author: Jtyoui@qq.com
// @note  : mysql数据库配置
package config

import "fmt"

type Mysql struct {
	Path     string // 服务器地址
	Port     int    // 端口
	Property string // 数据库属性
	Db       string // 数据库名
	Username string // 数据库用户名
	Password string // 数据库密码
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", m.Username, m.Password, m.Path, m.Port, m.Db, m.Property)
}
