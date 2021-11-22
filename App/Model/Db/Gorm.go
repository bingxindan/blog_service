package Db

import (
	"fmt"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"time"
)

type Type int32

const (
	MYSQL Type = iota
	POSTGRESQL
	SQLITE
	SQLSERVER
	CLICKHOUSE
)

type Conf struct {
	Type         Type
	Host         string
	Port         int
	Dbname       string
	User         string
	Password     string
	Charset      string
	ParseTime    string
	Loc          string
	SslMode      string
	TimeZone     string
	ReadTimeout  int
	WriteTimeout int
	Pool         struct {
		MaxIdle     int
		MaxOpen     int
		MaxLifetime time.Duration
	}
}

// GetDb 获取数据库链接，连接池使用默认值
func GetDb(conf Conf) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	switch conf.Type {
	case MYSQL:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s", conf.User, conf.Password, conf.Host, conf.Port, conf.Dbname, conf.Charset, conf.ParseTime, conf.Loc)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case POSTGRESQL:
		/*dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", conf.Host, conf.User, conf.Password, conf.Dbname, conf.Port, conf.SslMode, conf.TimeZone)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})*/
	case SQLITE:
		//db, err = gorm.Open(sqlite.Open(conf.Dbname), &gorm.Config{})
	case SQLSERVER:
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", conf.User, conf.Password, conf.Host, conf.Port, conf.Dbname)
		db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	case CLICKHOUSE:
		dsn := fmt.Sprintf("tcp://%s:%d?database=%s&username=%s&password=%s&read_timeout=%d&write_timeout=%d", conf.Host, conf.Port, conf.Dbname, conf.User, conf.Password, conf.ReadTimeout, conf.WriteTimeout)
		db, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	}
	return db, err
}

// GetPoolDb 自定义连接池
func GetPoolDb(conf Conf) (*gorm.DB, error) {
	db, err := GetDb(conf)
	if err != nil {
		return nil, err
	}
	// 获取通用数据库对象 sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if conf.Pool.MaxIdle > 0 {
		// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
		sqlDB.SetMaxIdleConns(conf.Pool.MaxIdle)
	}
	if conf.Pool.MaxOpen > 0 {
		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(conf.Pool.MaxOpen)
	}
	if conf.Pool.MaxLifetime > 0 {
		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(conf.Pool.MaxLifetime)
	}
	return db, nil
}
