package model

type FeedResponse struct {
	Response
	VideoList []VideoInfo `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}