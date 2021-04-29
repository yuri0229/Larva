package grpc

import (
	"context"

	"gf/app/service/internal/model"
	"gf/app/service/internal/service"
)

type demoService interface {
	ArticleDetail(ctx context.Context, id int64) (res *model.Article, err error)
}

var DemoService demoService

type GrpcServ struct {}

func New(s *service.Service) (serv *GrpcServ) {
	DemoService = s
	serv = &GrpcServ{}
	return
}

func (r *GrpcServ) Detail(c context.Context, req *Req) (res *ItemResp, err error) {
	item, err := DemoService.ArticleDetail(c, req.Id)
	if err != nil {
		return
	}
	if item != nil {
		res = &ItemResp{
			Id: item.Id,
			Title: item.Title,
		}
	}
	return
}

func (s *GrpcServ)  mustEmbedUnimplementedDemoGrpcServer() {}