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

func (m *Mysql) DSN() string {
	return fmt.Sprintf("%s%s?%s", m.String(), m.Db, m.Property)
}

func (m *Mysql) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/", m.Username, m.Password, m.Path, m.Port)
}

func (m *Mysql) CreateSQL() string {
	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", m.Db)
	return sql
}
