package http

import (
	"gf/api"
	"gf/internal/conf"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(c *conf.Config, serv *api.HttpServ) (engine *gin.Engine) {
	engine = gin.Default()
	engine.Use(cors())
	engine.GET("/", serv.Hello)
	engine.GET("/grpc", serv.Grpc)
	engine.Run(c.Http.Addr)
	return
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		header := c.Writer.Header()
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Headers", "Content-Length, Content-Type, XTOKEN")
		header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		header.Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, Authorization")
		header.Set("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}