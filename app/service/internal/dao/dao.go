package dao

import (
	"gf/app/service/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dao struct {
	Db   *gorm.DB
}

func New(c *conf.Config) (d *Dao) {
	return &Dao{
		Db: newDb(c),
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