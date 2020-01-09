package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tqsq2005/gin-gorm/setting"
	"time"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)

	//载入配置
	//TODO:#疑问#为什么不载入配置就获取不到？
	setting.LoadConf()

	dbType = viper.GetString("DB_CONNECTION")
	dbName = viper.GetString("DB_DATABASE")
	user = viper.GetString("DB_USERNAME")
	password = viper.GetString("DB_PASSWORD")
	host = viper.GetString("DB_HOST") + ":" + viper.GetString("DB_PORT")
	tablePrefix = viper.GetString("TABLE_PREFIX")

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(dbType)
		log.Println(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user,
			password,
			host,
			dbName))
		log.Println(err)
	}

	//更改默认表名称
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	//禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	db.SingularTable(true)
	// true启用Logger，显示详细日志
	db.LogMode(true)
	// 连接池#SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(10)
	// 连接池#SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(100)
	// 设置连接的最大生命周期为一小时。
	// 设置为0的话意味着没有最大生命周期，连接总是可重用(默认行为)。
	db.DB().SetConnMaxLifetime(time.Hour)
}

func CloseDB()  {
	defer db.Close()
}
