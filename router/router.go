package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"learn_go/go_api_template/handler/api"
	"learn_go/go_api_template/router/middleware"
	"time"
)

func RecoverErr() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				logrus.Error("发生错误:", err)
				c.AbortWithStatusJSON(500, gin.H{"msg": "server error!"})
			}
		}()
		c.Next()
	}
}

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 运行访问日志
	//g.Use(gin.Recovery())
	//g.Use(middleware.RunLog)
	g.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	//g.Use(middleware.NoCache)
	//g.Use(middleware.Options)
	//g.Use(middleware.Secure)
	g.Use(middleware.LogIp)
	//g.Use(mw...)
	g.Use(RecoverErr())
	g.RedirectTrailingSlash = false
	g.HandleMethodNotAllowed = true
	g.NoMethod(func(c *gin.Context) {
		c.JSON(403, gin.H{"msg": "method not allowed"})
	})
	g.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"msg": "route not found"})
	})

	// 服务发现路由
	svcd := g.Group("/api")
	{
		svcd.GET("/", api.HelloWorld)
	}

	// 业务路由

	return g
}
