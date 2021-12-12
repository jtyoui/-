// Package vampire
// @Time  : 2021/12/8 下午4:13
// @Author: Jtyoui@qq.com
// @note  : 数据库
package vampire

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
	password := GConfig.Mysql.Password
	host := GConfig.Mysql.Path
	port := GConfig.Mysql.Port
	username := GConfig.Mysql.Username
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", username, password, host, port)
	dbName := GConfig.Mysql.Db
	createDB := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", dbName)
	// 不存在数据库自动初始化
	if err := create(dsn, "mysql", createDB); err != nil {
		panic("创建数据库失败:" + err.Error())
	}

	drive = mysql.New(mysql.Config{
		DSN:                       GConfig.Mysql.Dsn(),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	})
	return
}
