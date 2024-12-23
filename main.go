package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Any("/auth/*proxyPath", reverseProxy("http://103.127.132.149:8081"))
	router.Use(authentication())
	router.Any("/customer/*proxyPath", reverseProxy("http://103.127.132.149:8082"))
	router.Any("/order/*proxyPath", reverseProxy("http://103.127.132.149:8083"))
	router.Any("/notif/*proxyPath", reverseProxy("http://103.127.132.149:8084"))
	router.Any("/tiket/*proxyPath", reverseProxy("http://103.127.132.149:8085"))
	router.Any("/event/*proxyPath", reverseProxy("http://103.127.132.149:8086"))
	router.Any("/rating/*proxyPath", reverseProxy("http://103.127.132.149:8087"))
	router.Any("/payment/*proxyPath", reverseProxy("http://103.127.132.149:8088"))

	router.Run(":8080")
}
func reverseProxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetURL := target + c.Param("proxyPath")
		// c.JSON(http.StatusOK, response{http.StatusOK, targetURL, nil})
		http.Redirect(c.Writer, c.Request, targetURL, http.StatusTemporaryRedirect)
	}
}
func authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, response{http.StatusUnauthorized, "Unauthentication", nil})
			c.Abort()
			return
		}
		c.Next()
	}
}

type response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
