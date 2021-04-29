package service

import (
	"context"
	"gf/app/api/grpc"
	"gf/app/api/internal/conf"
	"gf/app/api/internal/dao"
	"gf/app/api/internal/model"
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
	resp, err := s.dao.Grpc.Detail(ctx, &grpc.Req{Id: id})
	if err != nil {
		return
	}
	res = &model.Article{
		Id: resp.Id,
		Title: resp.Title,
	}
	return
}
