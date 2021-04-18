package service

import (
	"context"
	"gf/internal/conf"
	"gf/internal/dao"
	"gf/internal/model"
	"gorm.io/gorm"
)

type Service struct {
	dao *dao.Dao
}

func New(c *conf.Config) (s *Service) {
	s = &Service{
		dao: dao.New(c),
	}
	return
}

func (s *Service) ArticleDetail(ctx context.Context, id int64) (res *model.Article, err error) {
	res = &model.Article{}
	db := s.dao.Db
	if err = db.Table("Article").Where("id=?", id).First(res).Error; err == gorm.ErrRecordNotFound {
		res = nil
		err = nil
	}
	return
}
