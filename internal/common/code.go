// Package common
// @Time  : 2021/12/9 上午11:51
// @Author: Jtyoui@qq.com
// @note  : 状态码和异常
package common

import "errors"

type CodeType uint

const (
	SuccessCode  CodeType = 200
	ParamCode    CodeType = 1
	RegisterCode CodeType = 2
	LoginCode    CodeType = 3
	DataCode     CodeType = 4
)

// 根据状态码来获取错误信息
func (c CodeType) getError() error {
	switch c {
	case ParamCode:
		return errors.New("参数错误")
	case SuccessCode:
		return errors.New("成功")
	case RegisterCode:
		return errors.New("注册用户信息失败")
	case LoginCode:
		return errors.New("僵尸码输入错误")
	case DataCode:
		return errors.New("获取数据错误")
	}
	return nil
}
