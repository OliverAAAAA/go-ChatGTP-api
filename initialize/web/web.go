package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"go-chatgpt-api/config"
	"go-chatgpt-api/handlers"
	"go-chatgpt-api/utils/ip"
	"net/http"
	"time"
)

func Init() {

	router := gin.Default()
	gin.ForceConsoleColor()
	router.Use(RateLimitMiddleware(time.Second, 5, 5))
	//// # Headers
	// Allow CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	})
	//// # Add routes
	//==============业务============
	router.GET("/api/ask", handlers.Ask)
	router.GET("/api/createImg", handlers.GenerateImg)
	//==============业务end============

	//==============工具start============
	router.GET("/ip", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": ip.GetIpAddress(c.ClientIP()),
			"msg":  "success",
			"code": "200",
		})
	})
	// Add a health endpoint
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	//==============工具end============

	router.LoadHTMLFiles("html/search.html")
	//获取form参数
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "search.html", nil)
	})

	router.POST("/search", handlers.AskSearch)
	router.POST("/createImg", handlers.CreateImg)

	router.Run(fmt.Sprintf("%s:%s", *config.GetIp(), *config.GetPort()))
}

func RateLimitMiddleware(fillInterval time.Duration, cap, quantum int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			c.String(http.StatusForbidden, "请求人数过多请稍后再试..")
			c.Abort()
			return
		}
		c.Next()
	}
}
