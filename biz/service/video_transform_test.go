package service

import (
	"context"
	"testing"

	video "freefrom.space/videoTransform/hertz_gen/xy/video"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestVideoTransformService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewVideoTransformService(ctx, c)
	// init req and assert value
	req := &video.VideoTransformReq{}
	resp, err := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
