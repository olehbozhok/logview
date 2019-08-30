package main

import (
	"flag"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/gobuffalo/packr/v2"
)

var hostAddr = flag.String("host", ":8000", "host addr")

func main() {
	flag.Parse()

	box := packr.New("static", "static")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "test")
	})
	r.GET("file_list", filesList)

	r.GET("/", func(c *gin.Context) {
		b, err := box.Find("templates/base.html")
		if err != nil {
			c.Writer.Write([]byte(err.Error()))
		}
		c.Writer.Write(b)
	})

	r.GET("/show", showLog(box))

	r.Use(static.Serve("/static/", &ServeFileSystem{box}))

	r.Run(*hostAddr)
}
