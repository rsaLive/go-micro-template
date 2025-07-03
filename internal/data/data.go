package data

import (
	"fmt"
	configData "go-micro-template/config"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"strings"
	"time"
)

var ProviderSet = wire.NewSet(NewData)

type Data struct {
	dbCast *gorm.DB
	rCache *redis.Client
}

func NewData(confData *configData.AppConfig) *Data {
	fmt.Println("-----------------------------------")
	fmt.Printf("%+v", confData)
	fmt.Printf("host-----%+v", confData.DBFiee.Host)
	zap.L().Info("confData", zap.Any("confData", confData))
	zap.L().Info("confData.DBFiee", zap.Any("confData.DBFiee", confData.DBFiee))
	fmt.Println("-----------------------------------")
	connFiee := strings.Join([]string{confData.DBFiee.User, ":", confData.DBFiee.Password,
		"@tcp(", confData.DBFiee.Host, ":", confData.DBFiee.Port, ")/",
		confData.DBFiee.DbName, "?charset=utf8mb4&parseTime=true"}, "")
	DBFiee := loadMysqlConn(connFiee)
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     confData.Redis.Addr,
		Password: confData.Redis.Password,
		DB:       confData.Redis.DB,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	data := &Data{
		dbCast: DBFiee,
		rCache: RedisClient,
	}
	migration(data)
	return data
}

func loadMysqlConn(conn string) *gorm.DB {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conn,  // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  //设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) //打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	return db
}

func migration(data *Data) {
	//addColumn(&model_account.Users{}, "action_code")
	err := data.dbCast.AutoMigrate()
	if err != nil {
		zap.L().Error("register table fail--", zap.Error(err))
		os.Exit(0)
	}
}
