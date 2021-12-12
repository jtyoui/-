// Package config
// @Time  : 2021/12/10 上午9:37
// @Author: Jtyoui@qq.com
// @note  : 邮箱
package config

import "strings"

type Email struct {
	Address string `toml:"address"`
	SSL     bool   `toml:"ssl"`
	Secret  string `toml:"secret"`
}

func (e *Email) GetHostPort() (host string, port int) {
	switch e.suffix() {
	case "qq.com":
		host = "smtp.qq.com"
	}
	if e.SSL {
		port = 945
	} else {
		port = 25
	}
	return
}

func (e *Email) suffix() string {
	return strings.Split(e.Address, "@")[1]
}
