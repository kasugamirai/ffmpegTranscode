package video

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func DownloadVideo(url string) string {
	//sent GET request to video URL
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//generate a unique filename
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	filename := fmt.Sprintf("downloaded_video_%d_%d.mp4", time.Now().Unix(), r.Int())

	//create a file to hold the video
	out, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	//copy the response body to file
	_, err = io.Copy(out, resp.Body)

	if err != nil {
		panic(err)
	}

	hlog.Info("Downloaded video to", filename)

	//return the unique filename
	return filename
}
