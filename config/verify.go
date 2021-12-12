// Package config
// @Time  : 2021/10/12 下午5:20
// @Author: Jtyoui@qq.com
// @note  : 验证码
package config

type VerifyCode struct {
	Key       int // 验证码个数
	ImgWidth  int // 验证码宽度
	ImgHeight int // 验证码高度
}
