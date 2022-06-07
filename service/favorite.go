package service

import (
	"ByteDanceNoIdea/model"
	"gorm.io/gorm"
)

func Action(userId int64, videoId int64, actionType int) bool {
	switch {
		//点赞操作；
		case actionType == 1:
			//已经点赞过
			if err := model.Db.Model(&model.Favorite{}).Where("user_id = ? and video_id = ?", userId, videoId).First(&model.Favorite{}).Error; err == nil {
					return true
			}
			//更新点赞信息到数据库
			favoriteInfo := model.Favorite{
				UserID:	userId,
				VideoID:	videoId,
			}
			if err := model.Db.Create(&favoriteInfo).Error; err != nil {
				return false
			}
			//更新该video的favorite_count信息
			if err := model.Db.Model(&model.Video{}).Where("id = ?", videoId).UpdateColumn("favorite_count", gorm.Expr("favorite_count + 1")).Error; err != nil {
				return false
			}
			return true
		//取消点赞操作；
		case actionType == 2:
			var favoriteInfo model.Favorite
			//没有点赞
			if err := model.Db.Model(&model.Favorite{}).Where("user_id = ? and video_id = ?", userId, videoId).First(&favoriteInfo).Error; err != nil {
				return true
			}
			//更新取消点赞信息到数据库
			if err := model.Db.Model(&model.Favorite{}).Delete(&favoriteInfo).Error; err != nil {
				return false
			}
			//更新该video的favorite_count信息
			if err := model.Db.Model(&model.Video{}).Where("id = ?", videoId).UpdateColumn("favorite_count", gorm.Expr("favorite_count - 1")).Error; err != nil {
				return false
			}
			return true
		default:
			return false
	}
}

func FavoriteListService(userId int64) []model.VideoInfo{
	var videos []int64
	var videoList []model.VideoInfo
	model.Db.Model(&model.Favorite{}).Select("video_id").Where("user_id = ?",userId).Find(&videos)

	for videoId := range videos {
		var video model.Video
		model.Db.Model(&model.Video{}).Where("id = ?", videoId).First(&video)

		var author model.UserInfo
		model.Db.Model(&model.User{}).Where("id = ?", video.AuthorID).First(&author)
		author.IsFollow = false
		var isFavorite int64
		model.Db.Model(&model.Favorite{}).Where("user_id = ? and video_id = ?",userId,video.ID).Count(&isFavorite)
		
		videoList = append(videoList, model.VideoInfo{
			ID: video.ID,
			Author: author,
			PlayUrl:	video.PlayUrl,
			CoverUrl:	video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:	video.CommentCount,
			IsFavorite:	isFavorite > 0,
			Title:	video.Title,
		})
	}
	return videoList
}