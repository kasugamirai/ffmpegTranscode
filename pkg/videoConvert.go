package pkg

import (
	"github.com/rs/zerolog/log"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
)

func ConvertVideoFormat(inputPath string, outputPath string) {
	err := ffmpeg.Input(inputPath).
		Output(outputPath, ffmpeg.KwArgs{"c:v": "libx264"}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		log.Err(err).Msgf("error converting video format,", err)
	}

	// delete the input video file
	err = os.Remove(inputPath)
	if err != nil {
		log.Err(err).Msgf("error deleting input video file,", err)
	}
}