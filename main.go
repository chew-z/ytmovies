package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var (
	ctx, cancel = context.WithCancel(context.Background())
	path        = os.Getenv("MEDIA_PATH")
	port        = os.Getenv("PORT")
	uRL         = os.Getenv("URL")
	ytCmd       = os.Getenv("YT_DLP")
)

func init() {
	defer cancel()
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		if id := c.Query("id"); id == "" {
			// c.Redirect(http.StatusMovedPermanently, uRL+"media.html")
			c.HTML(http.StatusOK, "empty.html", gin.H{})
		} else {
			format := c.DefaultQuery("f", "")
			youTubeUrl := "https://www.youtube.com/watch?v=" + id
			log.Println(youTubeUrl)
			go ytDownload(id, format)
			u := uRL + id + ".mp4"
			c.HTML(http.StatusOK, "main.html", gin.H{
				"URL": u,
			})
		}
	})
	router.Run(":" + port)
}
