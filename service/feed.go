package service

import (
	"ByteDanceNoIdea/model"
	"time"
)

func FeedService(latestTime int64, userID int64) (videoInfos []model.VideoInfo, nextTime int64) {
	//查询给出时间以前的最近30条video数据
	var videos []model.Video
	model.Db.Model(&model.Video{}).Where("create_time < ?", latestTime).Order("create_time desc").Limit(30).Find(&videos)
	
	//下次刷新的时间，此处取本次刷新得到的最早的时间
	if len(videos) > 0 {
		nextTime = videos[len(videos) - 1].CreateTime
	} else {
		nextTime = time.Now().Unix()
	}
	
	
	//写入返回videoInfos信息
	var favoriteCount int64
	for _, video := range videos {
		favoriteCount = 0
		var author model.UserInfo
		model.Db.Model(&model.User{}).Where("id = ?", video.AuthorID).First(&author)
		if userID != -1 {
			//查询用户是否follow，此处暂时默认未follow
			author.IsFollow = false

			//查询视频是否点赞
			model.Db.Model(&model.Favorite{}).Where("user_id = ? and video_id = ?", userID, video.ID).Count(&favoriteCount)
		}
		videoInfos = append(videoInfos, model.VideoInfo{
			ID:	video.ID,
			Author:	author,
			PlayUrl:	video.PlayUrl,
			CoverUrl:	video.CoverUrl,
			FavoriteCount:	video.FavoriteCount,
			CommentCount: video.CommentCount,
			IsFavorite:	favoriteCount > 0,
			Title:	video.Title,
		})
	}
	return
}