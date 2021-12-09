// Package vampire
// @Time  : 2021/12/9 下午3:53
// @Author: Jtyoui@qq.com
// @note  : 数据库操作
package vampire

type DB interface {
	TableName() string
}

func login[T DB](key string, value interface{}) (t T, ok bool) {
	if err := GDb.Where(key+" = ?", value).First(&t).Error; err != nil {
		return
	}
	ok = true
	return
}
