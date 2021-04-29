package dao

import (
	"gf/app/api/grpc"
	"gf/app/api/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dao struct {
	Db   *gorm.DB
	Grpc grpc.DemoGrpcClient
}

func New(c *conf.Config) (d *Dao) {
	return &Dao{
		Db: newDb(c),
		Grpc: grpc.NewClient(c),
	}
}

func newDb(c *conf.Config) (db *gorm.DB) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: c.Db.Dsn,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return
}