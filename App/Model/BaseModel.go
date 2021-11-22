package Model

import (
	"blog_service/App/Model/Db"
	"github.com/bingxindan/bxd_go_lib/tools/confutil"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"strings"
	"sync"
	"time"
)

var dbMap = make(map[string]*gorm.DB)
var lock sync.Mutex

func GetDb(connection string) (*gorm.DB, error) {
	if con, ok := dbMap[connection]; ok {
		return con, nil
	}
	conf := getConf(connection)
	db, err := Db.GetPoolDb(conf)
	if err != nil {
		return nil, err
	}
	//加锁，避免竞争重复建连接
	lock.Lock()
	dbMap[connection] = db
	lock.Unlock()
	return db, nil
}
func getConf(connection string) Db.Conf {
	dbType := strings.ToLower(confutil.GetConf(connection, "type"))
	host := confutil.GetConf(connection, "host")
	port := cast.ToInt(confutil.GetConf(connection, "port"))
	dbName := confutil.GetConf(connection, "dbName")
	user := confutil.GetConf(connection, "user")
	pass := confutil.GetConf(connection, "pass")
	charset := confutil.GetConf(connection, "charset")
	parseTime := confutil.GetConf(connection, "parseTime")
	loc := confutil.GetConf(connection, "loc")
	sslMode := confutil.GetConf(connection, "sslMode")
	timeZone := confutil.GetConf(connection, "timeZone")
	readTimeout := cast.ToInt(confutil.GetConf(connection, "readTimeout"))
	writeTimeout := cast.ToInt(confutil.GetConf(connection, "writeTimeout"))
	maxIdle := cast.ToInt(confutil.GetConf(connection, "charset"))
	maxOpen := cast.ToInt(confutil.GetConf(connection, "charset"))
	maxLifetime := cast.ToInt(confutil.GetConf(connection, "charset"))
	var dbT Db.Type
	switch dbType {
	case "mysql":
		dbT = Db.MYSQL
	case "postgresql":
		dbT = Db.POSTGRESQL
	case "sqlite":
		dbT = Db.SQLITE
	case "sqlserver":
		dbT = Db.SQLSERVER
	case "clickhouse":
		dbT = Db.CLICKHOUSE
	default:
		panic("unsupported database:" + dbType)
	}
	conf := Db.Conf{
		Type:         dbT,
		Host:         host,
		Port:         port,
		Dbname:       dbName,
		User:         user,
		Password:     pass,
		Charset:      charset,
		ParseTime:    parseTime,
		Loc:          loc,
		SslMode:      sslMode,
		TimeZone:     timeZone,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		Pool: struct {
			MaxIdle     int
			MaxOpen     int
			MaxLifetime time.Duration
		}{
			MaxIdle:     maxIdle,
			MaxOpen:     maxOpen,
			MaxLifetime: time.Duration(maxLifetime) * time.Second,
		},
	}
	return conf
}
