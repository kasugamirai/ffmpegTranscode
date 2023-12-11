package service

import (
	"context"
	"freefrom.space/videoTransform/pkg/discordX"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	video "freefrom.space/videoTransform/hertz_gen/xy/video"
	"github.com/cloudwego/hertz/pkg/app"
)

type VideoTransformService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewVideoTransformService(Context context.Context, RequestContext *app.RequestContext) *VideoTransformService {
	return &VideoTransformService{RequestContext: RequestContext, Context: Context}
}

func (h *VideoTransformService) Run(req *video.VideoTransformReq) (resp *video.VideoTransformResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	videoURL := req.VideoURL

	convertVideo, err := discordX.DownloadAndUploadVideoToDiscord(videoURL)
	if err != nil {
		hlog.Errorf("error converting video format,", err)
		return nil, err
	}
	resp = &video.VideoTransformResp{
		VideoURL: convertVideo,
		Ok:       true,
	}
	return
}
