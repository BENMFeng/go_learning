package main

import (
 "log"
 "net/http"

 "github.com/gin-gonic/gin"
)

func MiddlewareA() gin.HandlerFunc {
 return func(c *gin.Context) {
  log.Println("MiddlewareA before request")
  // before request
  c.Next()
  // after request
  log.Println("MiddlewareA after request")
 }
}

func MiddlewareB() gin.HandlerFunc {
 return func(c *gin.Context) {
  log.Println("MiddlewareB before request")
  // before request
  c.Next()
  // after request
  log.Println("MiddlewareB after request")
 }
}

// This function's name is a must. App Engine uses it to drive the requests properly.
func main() {
 // Starts a new Gin instance with no middle-ware
 r := gin.New()
 r.Use(MiddlewareA(), MiddlewareB())
 r.GET("/ping", func(c *gin.Context) {
  c.String(http.StatusOK, "pong")
  log.Println("pong")
 })
 r.Run(":8080")
}
