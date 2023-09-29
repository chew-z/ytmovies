package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/levenlabs/golib/timeutil"
	"github.com/wader/goutubedl"
)

func ytDownload(id string, format string) {
	ytUrl := "https://www.youtube.com/watch?v=" + id
	goutubedl.Path = ytCmd
	opts := &goutubedl.Options{}
	result, err := goutubedl.New(context.Background(), ytUrl, *opts)
	if err != nil {
		log.Fatal(err)
	}
	downloadResult, err := result.Download(context.Background(), format)
	if err != nil {
		log.Fatal(err)
	}
	defer downloadResult.Close()
	f, err := os.Create(id + ".mp4")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	io.Copy(f, downloadResult)
}

func extractInfo(id string) {
	var info goutubedl.Info
	movie := &Movie{
		Id: id,
	}
	file := path + "-.info.json"
	infoFile, err := os.Open(file)
	if err != nil {
		log.Println("opening -.info.json file", err.Error())
	}
	jsonParser := json.NewDecoder(infoFile)
	if err = jsonParser.Decode(&info); err != nil {
		log.Println("parsing -.info.json file", err.Error())
	}

	log.Println(info.ID)
	if info.ID == id {
		movie.Channel = info.Channel
		movie.Title = info.Title
		movie.Duration = fmt.Sprintf("%s", time.Duration(info.Duration*float64(time.Second)))
		ts := timeutil.TimestampFromFloat64(info.Timestamp)
		movie.Timestamp = fmt.Sprintf("%s", ts)
		movie.UploadDate = info.UploadDate
		movie.ReleaseDate = info.ReleaseDate
		movie.Url = "https://www.youtube.com/watch?v=" + id
		movie.Link = uRL + id + ".mp4"

		log.Printf("%s - %s (%s)", info.Channel, info.Title, time.Duration(info.Duration*float64(time.Second)))

		movieJson, err := json.MarshalIndent(movie, "", "    ")
		if err != nil {
			log.Println(err.Error())
		}
		file := "/var/www/html/media/" + id + ".json"
		err = ioutil.WriteFile(file, movieJson, 0664)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
