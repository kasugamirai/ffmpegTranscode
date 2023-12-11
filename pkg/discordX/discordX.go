package discordX

import (
	"fmt"
	"freefrom.space/videoTransform/conf"
	"freefrom.space/videoTransform/pkg"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"math/rand"
	"os"
	"time"
)

func UploadVideo(s *discordgo.Session, videoPath string) string {
	// generate a unique filename for the converted video
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	convertedVideoPath := fmt.Sprintf("./converted_video_%d_%d.mp4", time.Now().Unix(), r1.Int())

	// convert the video format to .mp4
	pkg.ConvertVideoFormat(videoPath, convertedVideoPath)

	// open the converted video file
	file, err := os.Open(convertedVideoPath)
	if err != nil {
		log.Err(err).Msgf("error opening video file,", err)
	}
	defer file.Close()

	// create a new file message
	message := &discordgo.MessageSend{
		Files: []*discordgo.File{
			{
				Name:   convertedVideoPath,
				Reader: file,
			},
		},
	}

	// send the message with the video file
	msg, err := s.ChannelMessageSendComplex(conf.GetConf().Discord.ChannelId, message)
	if err != nil {
		log.Err(err).Msgf("error sending message,", err)
	}

	// delete the converted video file
	err = os.Remove(convertedVideoPath)
	if err != nil {
		log.Err(err).Msgf("error deleting converted video file,", err)
	}

	// return the URL of the video
	return msg.Attachments[0].URL
}

func DownloadAndUploadVideoToDiscord(videoURL string) (string, error) {
	// Download the video and get the unique filename
	filename := pkg.DownloadVideo(videoURL)

	dg, err := discordgo.New("Bot " + conf.GetConf().Discord.Token)
	if err != nil {
		log.Err(err).Msgf("error creating Discord session,", err)
		return "", err
	}

	// Use the unique filename when uploading the video
	videoLink := UploadVideo(dg, filename)
	log.Info().Msgf("Video link:", videoLink)
	return videoLink, nil
}
