// Package pkg
// @Time  : 2022/3/4 下午2:21
// @Author: Jtyoui@qq.com
// @note  :
package pkg

import (
	"net"
	"strings"
)

// GetLocalBoundIP 获取本机向外绑定的Ipv4地址
func GetLocalBoundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return "www.我和僵尸有个约会.com"
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return strings.Split(localAddr.String(), ":")[0]
}
