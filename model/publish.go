package model

type PublishResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type PublishListResponse struct {
	PublishResponse
	VideoList  []VideoInfo `json:"video_list"`  // 用户发布的视频列表
}