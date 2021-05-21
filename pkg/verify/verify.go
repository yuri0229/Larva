package verify

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	"net/url"
	"strings"
)

type Conf struct {
	Secret 	string
}

type Verify struct {
	secret	string
}

func New(conf *Conf) *Verify {
	return &Verify{
		secret: conf.Secret,
	}
}

func (v *Verify) Verify(ctx *gin.Context) {
	err := v.verify(ctx);
	if err != nil {
		ctx.JSON(200, gin.H{"code":1, "msg":err.Error()})
		ctx.Abort()
		return
	}
}

func (v *Verify) verify(ctx *gin.Context) error {
	params := ctx.Request.URL.Query()
	timestamp := params.Get("timestamp")
	if timestamp == "" {
		return errors.New("缺少timestamp")
	}
	sign := params.Get("sign")
	params.Del("sign")
	if sign != Sign(params, v.secret) {
		return errors.New("签名失败")
	}
	return nil
}

func Sign(params url.Values, secret string) string {
	p := params.Encode()
	if strings.IndexByte(p, '+') > -1 {
		p = strings.Replace(p, "+", "%20", -1)
	}
	d := md5.Sum([]byte(p + secret))
	return hex.EncodeToString(d[:])
}