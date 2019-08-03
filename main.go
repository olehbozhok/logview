package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/gobuffalo/packr/v2"
)

func main() {

	box := packr.New("static", "static")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "test")
	})
	r.GET("file_list", filesList)

	r.GET("/", func(c *gin.Context) {
		b, err := box.Find("./templates/base.html")
		if err != nil {
			c.Writer.Write([]byte(err.Error()))
		}
		c.Writer.Write(b)
	})

	r.GET("/show", showLog(box))

	r.Use(static.Serve("/static", &ServeFileSystem{box}))

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8000")
}
