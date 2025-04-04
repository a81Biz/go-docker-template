package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func proxyHandler(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, err := url.Parse(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid proxy target"})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)

		// Reescribe la ruta: /api/users/... => /users/...
		c.Request.URL.Path = "/users" + strings.TrimPrefix(c.Request.URL.Path, "/api/users")

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	router := gin.Default()
	//ésto cambió
	router.Use(func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
			c.Abort()
			return
		}
		c.Next()
	})

	// Ruta de ejemplo para users-module
	router.Any("/api/users/*proxyPath", proxyHandler("http://users-module:5235"))

	log.Println("API Manager corriendo en puerto 8080")
	router.Run(":8080")
}
