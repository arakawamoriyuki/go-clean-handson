package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MyBenchLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// request 処理の前
		// サンプル変数を設定
		c.Set("example", "12345")

		c.Next()

		// request 処理の後
		// レイテンシ表示
		latency := time.Since(t)
		log.Print(latency)

		// 送信予定のステータスコードを表示
		status := c.Writer.Status()
		log.Println(status)
	}
}

func AuthRequiredMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 認証失敗させる
		c.JSON(http.StatusUnauthorized, gin.H{"error": true, "from": "AuthRequiredMiddleware"})
		c.Abort()
	}
}

func benchmarkEndpoint(c *gin.Context) {
	// MyBenchLoggerMiddlewareで設定された変数を表示
	example := c.MustGet("example").(string)
	log.Println(example)

	c.JSON(http.StatusOK, gin.H{"error": false, "from": "benchmarkEndpoint"})
}

func meEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"error": false, "from": "meEndpoint"})
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/benchmark", MyBenchLoggerMiddleware(), benchmarkEndpoint)

	authorized := r.Group("/auth")
	authorized.Use(AuthRequiredMiddleware())
	{
		authorized.GET("/me", meEndpoint)
	}

	r.Run(":8080")
}
