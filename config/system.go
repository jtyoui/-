// Package config
// @Time  : 2021/10/12 上午9:58
// @Author: Jtyoui@qq.com
// @note  : 系统配置
package config

// System 该系统的一些基本配置
type System struct {
	PrefixRouter  string // 一级路由的前缀
	Port          int    // 接口开放的端口
	Release       bool
	DataBase      string `toml:"dbType"` // 默认mysql数据库；如果需要切换不同数据，只需要添加对应的数据库（后续）
	UseMultipoint bool   `toml:"stopUseMultipoint"`
}
