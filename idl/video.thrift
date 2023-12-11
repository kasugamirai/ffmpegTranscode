namespace go xy.video

struct videoTransformReq{
	1: string video_url (api.body="video_url");
}

struct videoTransformResp{
    1: bool ok
    2: string video_url
}

service videoTransformService{
    videoTransformResp VideoTransform(1: videoTransformReq request) (api.post="api/video/transform");
}