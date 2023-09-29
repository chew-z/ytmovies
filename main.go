package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type Params struct {
	Id     string `form:"id"`
	Format string `form:"format"`
}

type Movie struct {
	Id          string
	Title       string
	Channel     string
	Duration    string
	Link        string
	ReleaseDate string
	Timestamp   string
	UploadDate  string
	Url         string
}

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
		id := c.DefaultQuery("id", "")
		format := c.DefaultQuery("f", "")
		if id != "" {
			go ytDownload(id, format)
			go func() {
				time.Sleep(10 * time.Second)
				extractInfo(id)

			}()
		}
		u := uRL + id + ".mp4"
		c.HTML(http.StatusOK, "main.html", gin.H{
			"URL": u,
		})

	})
	router.POST("/", startPage)
	router.Run(":" + port)
}

func startPage(c *gin.Context) {
	var params Params
	id := ""
	format := ""
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		log.Println(params.Id)
		id = params.Id
		format = params.Format
		go ytDownload(id, format)
	}
	u := uRL + id + ".mp4"
	c.Redirect(302, u)
}
