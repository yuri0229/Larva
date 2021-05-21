package http

import (
	"context"
	"gf/pkg/verify"
	"net/http"
	"strconv"

	"gf/app/api/internal/model"
	"gf/app/api/internal/service"
	"github.com/gin-gonic/gin"
)

type demoService interface {
	ArticleDetail(ctx context.Context, id int64) (res *model.Article, err error)
}

var DemoServ demoService

func json(c *gin.Context, data interface{}, err error) {
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

func hello(c *gin.Context) {
	res := map[string]string{"say":"hello"}
	json(c, res, nil)
}

func detail(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		json(c, nil, err)
		return
	}
	res, err := DemoServ.ArticleDetail(c, id)
	if err != nil {
		json(c, nil, err)
		return
	}
	json(c, res, err)
}

func RouteRetister(e *gin.Engine, s *service.Service){
	DemoServ = s
	conf := &verify.Conf{s.Conf.Http.Secret}
	v := verify.New(conf)
	e.GET("/", hello)
	e.GET("/detail", v.Verify, detail)
}