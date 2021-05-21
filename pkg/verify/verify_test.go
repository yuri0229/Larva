package verify

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestVerify(t *testing.T) {
	v := New(&Conf{Secret:"abc"})
	engine := gin.Default()
	engine.GET("/verify", v.Verify, func(context *gin.Context) {
		context.JSON(200, gin.H{"status":0})
	})
	engine.POST("/verify", v.Verify, func(context *gin.Context) {
		context.JSON(200, gin.H{"status":0})
	})
	engine.Run("127.0.0.1:8200")
	return
}