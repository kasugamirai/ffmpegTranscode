package video

import (
	"context"

	"freefrom.space/videoTransform/biz/service"
	"freefrom.space/videoTransform/biz/utils"
	video "freefrom.space/videoTransform/hertz_gen/xy/video"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// VideoTransform .
// @router api/video/transform [POST]
func VideoTransform(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video.VideoTransformReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewVideoTransformService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
