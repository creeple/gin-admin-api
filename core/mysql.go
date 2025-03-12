/**
 * @Author: lixiang
 * @Date: 2025/3/11 15:26
 * @Description: mysql.go
 */

package core

import (
	"fmt"
	"gin-admin-api/config"
	"gin-admin-api/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var Db *gorm.DB

func MysqlInit() error {
	var err error
	var dbConfig = config.Config.Mysql

	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Db,
		dbConfig.Charset)
	Db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	if Db.Error != nil {
		log.Fatal(err)
	}
	connect, err := Db.DB()
	if err != nil {
		return err
	}
	connect.SetMaxIdleConns(dbConfig.MaxIdle)
	global.Log.Infof("mysql init success")
	return nil
}
