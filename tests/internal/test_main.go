package main

import (
	"github.com/XIAHUALOU/fitness-gin/fit"
	"github.com/XIAHUALOU/fitness-gin/tests/internal/Configuration"
	"github.com/XIAHUALOU/fitness-gin/tests/internal/classes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func cros() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,X-Token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

	}
}
func errorFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				c.AbortWithStatusJSON(400, gin.H{"my": e})
			}
		}()
		c.Next()
	}
}
func main() {
	//Ignite方法 支持 配置原始Gin 中间件，全局的
	FitGin.Ignite(cros(), errorFunc()).
		//Config(Configuration.NewMyConfig()).
		//Attach(fairing.NewGlobalFairing()).
		Mount("", classes.NewIndexClass()). //控制器，挂载到v1
		Config(Configuration.NewRouterConfig()).
		Launch()
}
