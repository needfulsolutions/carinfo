package main

import (
    _ "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

/*
func response(c *gin.Context) {

    c.JSON(http.StatusOK, gin.H{
        "message": "pong",
    })
}


func AddRoutes(rg *gin.RouterGroup) {

    rg.GET("/volvo/:model/*year", func(c *gin.Context) {
        
        model  := c.Param("model")
        year   := c.Param("year")
        msg    := "model=" + model + ";year=" + year + ";" + c.FullPath()
        c.String(http.StatusOK, msg)
    })
}
*/

func main() {
    
    r := gin.Default()

    // web
    r.LoadHTMLGlob("static/templates/**/*.html")

    // static routes for web
    r.Static("/assets/css", "./static/css")

    // show home page !
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "pages/root", gin.H {
            "header_title": "homepage",
        })
    })

    //v1 := r.Group("/v1")
    //AddRoutes(v1)
    //r.GET("/ping", response)

    r.Run()    
}

