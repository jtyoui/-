// Package config
// @Time  : 2021/10/12 下午3:11
// @Author: Jtyoui@qq.com
// @note  : jwt配置
package config

type Jwt struct {
	SigningKey  string // 盐
	ExpiresTime int64  // 过期时间
}
