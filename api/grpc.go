package api

import (
	"context"
	"gf/internal/service"
)

type GrpcServ struct {
	s	*service.Service
}

func NewGrpc(s *service.Service) (serv *GrpcServ) {
	serv = &GrpcServ{
		s: s,
	}
	return
}

func (r *GrpcServ) Detail(c context.Context, req *Req) (res *ItemResp, err error) {

	item, err := r.s.ArticleDetail(c, req.Id)
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