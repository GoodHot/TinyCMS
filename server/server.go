package server

import "github.com/gin-gonic/gin"

type HTTPServer struct {
	ServerAddr string `val:"${server.addr}"`
}

func (server *HTTPServer) Startup() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(server.ServerAddr)
}
