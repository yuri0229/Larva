package api

import (
	"log"
	"net/http"
	"strconv"

	"gf/internal/conf"
	"gf/internal/service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type HttpServ struct {
	s *service.Service
	rpcClient DemoGrpcClient
}

func NewHttp(c *conf.Config, s *service.Service) (serv *HttpServ) {
	conn, err := grpc.Dial(c.Rpc.Addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	rpcClient := NewDemoGrpcClient(conn)

	serv = &HttpServ{
		s: s,
		rpcClient: rpcClient,
	}
	return
}

func (r *HttpServ) json(c *gin.Context, data interface{}, err error) {
	var (
		code int
		msg string
	)
	if err != nil {
		code = 1
		msg = err.Error()
	}
	c.JSON(http.StatusOK, gin.H{"code":code, "msg":msg, "data":data})
}

func (r *HttpServ) Hello(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		r.json(c, nil, err)
		return
	}
	res, err := r.s.ArticleDetail(c, id)
	if err != nil {
		r.json(c, nil, err)
		return
	}
	r.json(c, res, err)
}

func (r *HttpServ) Grpc(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
		return
	}
	res, err :=  r.rpcClient.Detail(c, &Req{Id: id})
	r.json(c, res, err)
}