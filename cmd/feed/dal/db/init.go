package db

import (
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var (
	DB *gorm.DB
	// RDClient *redis.RDClient
)

// Init Mysql DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(conf.GetConf().GetString("db.mysql.dsn")),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	DB.AutoMigrate(&Video{})

	// RDClient = redis.NewRDClient(conf.GetConf().GetInt("db.redis.id.feed"))
}
