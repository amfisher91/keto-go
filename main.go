package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
    router := gin.Default()

    router.GET("/", home)

    router.Run()
}

func home(c *gin.Context) {
    c.String(http.StatusOK, "Hello, World!")
}
