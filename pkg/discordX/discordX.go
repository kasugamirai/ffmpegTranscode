package discordX

import (
	"context"
	"fmt"
	"freefrom.space/videoTransform/biz/dal/sqlite"
	"freefrom.space/videoTransform/conf"
	v "freefrom.space/videoTransform/ent/video"
	"freefrom.space/videoTransform/pkg/video"
	"github.com/bwmarrin/discordgo"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"math/rand"
	"os"
	"time"
)

func uploadVideo(s *discordgo.Session, videoPath string) (string, error) {
	// generate a unique filename for the converted video
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	convertedVideoPath := fmt.Sprintf("./converted_video_%d_%d.mov", time.Now().Unix(), r1.Int())

	// convert the video format to .mp4
	video.ConvertVideoFormat(videoPath, convertedVideoPath)

	// open the converted video file
	file, err := os.Open(convertedVideoPath)
	if err != nil {
		hlog.Error("error opening video file,", err)
		return "", err
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
		hlog.Error("error sending message,", err)
		return "", err
	}

	// delete the converted video file
	err = os.Remove(convertedVideoPath)
	if err != nil {
		hlog.Error("error deleting converted video file,", err)
		return "", err
	}

	if err != nil {
		hlog.Error("error sending message,", err)
		return "", err
	}

	// return the URL of the video
	return msg.Attachments[0].URL, nil
}

func DownloadAndUploadVideoToDiscord(videoURL string) (string, error) {
	ok := checkURLInDB(videoURL, context.Background())
	fmt.Println(ok)
	if ok {
		return findConvertURL(context.Background(), videoURL)
	}

	// Download the video and get the unique filename
	filename := video.DownloadVideo(videoURL)

	dg, err := discordgo.New("Bot " + conf.GetConf().Discord.Token)
	if err != nil {
		hlog.Error("error creating Discord session,", err)
		return "", err
	}

	// Use the unique filename when uploading the video
	videoLink, err := uploadVideo(dg, filename)
	if err != nil {
		hlog.Error("error uploading video,", err)
		return "", err
	}
	err = storeOriginAndConvertURLInDB(context.Background(), videoURL, videoLink)
	if err != nil {
		hlog.Error("error storing convert url,", err)
		return "", err
	}
	hlog.Error("Video link:", videoLink)
	return videoLink, nil
}

func checkURLInDB(videoURL string, ctx context.Context) bool {
	exists, err := sqlite.Client.Video.Query().Where(v.OriginURLEQ(videoURL)).Exist(ctx)

	if err != nil {
		hlog.Error("error querying video URL")
		return false
	}
	return exists
}

func storeOriginAndConvertURLInDB(ctx context.Context, originURL, convertURL string) error {
	err := sqlite.Client.Video.Create().
		SetOriginURL(originURL).
		SetConvertURL(convertURL).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func findConvertURL(ctx context.Context, originURL string) (string, error) {
	id, err := sqlite.Client.Video.Query().Where(v.OriginURLEQ(originURL)).OnlyID(ctx)
	if err != nil {
		return "", err
	}

	convertedVideo, err := sqlite.Client.Video.Get(ctx, id)
	if err != nil {
		return "", err
	}

	return convertedVideo.ConvertURL, nil
}
