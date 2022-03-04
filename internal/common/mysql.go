// Package common
// @Time  : 2021/12/8 下午4:13
// @Author: Jtyoui@qq.com
// @note  : 数据库
package common

import (
	"database/sql"
	"fmt"
	"github.com/jtyoui/my-date-with-a-vampire/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GDb *gorm.DB

// 创建数据库并初始化
func create(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

func InitMysql() (drive gorm.Dialector) {
	conn := config.GConfig.Mysql
	if err := create(conn.String(), "mysql", conn.CreateSQL()); err != nil { // 不存在数据库自动初始化
		panic("创建数据库失败:" + err.Error())
	}
	drive = mysql.New(mysql.Config{
		DSN:                       conn.DSN(),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	})
	return
}
