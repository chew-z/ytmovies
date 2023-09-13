package main

import (
	"context"
	"io"
	"log"
	"os"

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
