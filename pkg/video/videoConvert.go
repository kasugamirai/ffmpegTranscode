package video

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
)

func ConvertVideoFormat(inputPath string, outputPath string) {
	err := ffmpeg.Input(inputPath).
		Output(outputPath, ffmpeg.KwArgs{
			"c:v":       "libx265",
			"crf":       "28",   // 控制编码的质量，值越小质量越好，文件越大。常见值在18-28之间
			"preset":    "fast", // 控制编码速度和压缩率的平衡，可选值包括ultrafast, superfast, veryfast, faster, fast, medium, slow, slower, veryslow
			"profile:v": "main", // 控制编码级别，可选值包括main, main10, main-still-picture
			"f":         "mov",  // 指定输出文件的格式为MOV
		}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		hlog.Error("error converting video format,", err)
	}

	// delete the input video file
	err = os.Remove(inputPath)
	if err != nil {
		hlog.Error("error deleting input video file,", err)
	}
}
