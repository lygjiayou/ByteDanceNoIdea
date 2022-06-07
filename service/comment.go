package service

import (
	"ByteDanceNoIdea/model"
	"ByteDanceNoIdea/utils/errmsg"
	"gorm.io/gorm"
	"time"
)

func CommentActionService(commentActionParam model.CommentRequestParam, userId int64) (commentInfo model.CommentInfo, err int) {
	videoId := commentActionParam.VideoID
	switch commentActionParam.ActionType {
	//发表评论
	case 1:
		comment := model.Comment{
			VideoID:	videoId,
			CommentUserID:	userId,
			Content:	commentActionParam.CommentText,
			CreateTime:	time.Now().Unix(),
		}
		
		//先更新评论数量为+1；
		if err := model.Db.Model(&model.Video{}).Where("id = ?", videoId).UpdateColumn("comment_count", gorm.Expr("comment_count + 1")).Error; err != nil {
			return commentInfo,errmsg.ERROR_COMMENT_FAIL
		} else {
			//更新评论信息数据库，
			//若失败，更新评论数量为-1；
			if err := model.Db.Create(&comment).Error; err != nil {
				model.Db.Model(&model.Video{}).Where("id = ?", videoId).UpdateColumn("comment_count", gorm.Expr("comment_count - 1"))
				return commentInfo, errmsg.ERROR_COMMENT_FAIL
			//成功，返回评论详细信息
			} else {
				var author model.User
				model.Db.Model(&model.User{}).Where("id = ?", userId).First(&author)
				var authorInfo model.UserInfo
				//authorOrigin -> authorInfo
				authorInfo.ID = author.ID
				authorInfo.UserName = author.UserName
				authorInfo.FollowCount = author.FollowCount
				authorInfo.FollowerCount = author.FollowerCount
				authorInfo.IsFollow = false //temp

				model.Db.Model(&model.Comment{}).Where("video_id = ? and comment_user_id = ?", videoId, userId).First(&comment)
				commentInfo = model.CommentInfo{
					ID:	comment.ID,
					Author:	authorInfo,
					Content: comment.Content,
					CreateTime:	comment.CreateTime,
				}
				return commentInfo, errmsg.COMMENT_SUCCESS
			}
		}
		
	//删除评论
	case 2:
		//先更新评论数量为-1；
		if err := model.Db.Model(&model.Video{}).Where("id = ?", videoId).UpdateColumn("comment_count", gorm.Expr("comment_count - 1")).Error; err != nil {
			return commentInfo,errmsg.ERROR_DELETE_FAIL
		} else {
			var comment model.Comment
			model.Db.Model(&model.Comment{}).Where("id = ?", commentActionParam.CommentID).First(&comment)
			if model.Db.Model(&model.Comment{}).Delete(&comment).RowsAffected == 0 {
				model.Db.Model(&model.Video{}).Where("id = ?", videoId).UpdateColumn("comment_count", gorm.Expr("comment_count + 1"))
				return commentInfo, errmsg.ERROR_DELETE_FAIL
			} else {
				return commentInfo, errmsg.DELETE_COMMENT_SUCCESS
			}
		}	
	default:
		return commentInfo, errmsg.ERROR_COMMENT_FAIL
	}
}

func CommentListService(userid int64, videoid int64) []model.CommentInfo {
	var commentList []model.CommentInfo
	var comments []model.Comment
	model.Db.Model(&model.Comment{}).Where("video_id = ?", videoid).Find(&comments)

	for _, comment := range comments {
		var author model.User
		model.Db.Model(&model.User{}).Where("id = ?", comment.CommentUserID).First(&author)
		var authorInfo model.UserInfo
		//authorOrigin -> authorInfo
		authorInfo.ID = author.ID
		authorInfo.UserName = author.UserName
		authorInfo.FollowCount = author.FollowCount
		authorInfo.FollowerCount = author.FollowerCount
		authorInfo.IsFollow = false //temp

		commentList = append(commentList, model.CommentInfo{
			ID:	comment.ID,
			Author: authorInfo,
			Content:	comment.Content,
			CreateTime: comment.CreateTime,
		})
	}
	return commentList
}