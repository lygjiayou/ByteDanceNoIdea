package service

import (
	"ByteDanceNoIdea/model"
	"ByteDanceNoIdea/utils/errmsg"
	"gorm.io/gorm"
)

func FollowAction(actionType int, targetId int64, userId int64) int {
	switch actionType {
	case 1:
		if err := model.Db.Model(&model.Relation{}).Where("user_id = ? and follower_id = ?", targetId, userId).First(&model.Relation{}).Error; err == nil {
			return errmsg.ERROR_ALREADY_FOLLOW
		}
		relation := model.Relation{
			UserID:	targetId,
			FollowerID:	userId,
		}

		if err := model.Db.Model(&model.User{}).Where("id = ?", targetId).UpdateColumn("follower_count", gorm.Expr("follower_count + 1")).Error; err != nil {
			return errmsg.ERROR_FOLLOW_FAIL
		} else {
			if err := model.Db.Model(&model.User{}).Where("id = ?", userId).UpdateColumn("follow_count", gorm.Expr("follow_count + 1")).Error; err != nil {
				model.Db.Model(&model.User{}).Where("id = ?", targetId).UpdateColumn("follower_count", gorm.Expr("follower_count - 1"))
				return errmsg.ERROR_FOLLOW_FAIL
			} else {
				if err := model.Db.Create(&relation).Error; err != nil {
					model.Db.Model(&model.User{}).Where("id = ?", targetId).UpdateColumn("follower_count", gorm.Expr("follower_count - 1"))
					model.Db.Model(&model.User{}).Where("id = ?", userId).UpdateColumn("follow_count", gorm.Expr("follow_count - 1"))
					return errmsg.ERROR_FOLLOW_FAIL
				} else {
					return errmsg.FOLLOW_SUCCESS
				}
			}
		}
		
	case 2:
		var relation model.Relation
		if err := model.Db.Model(&model.Relation{}).Where("user_id = ? and follower_id = ?", targetId,userId).First(&relation).Error; err != nil {
			return errmsg.ERROR_ALREADY_CANCEL
		}

		if err := model.Db.Model(&model.User{}).Where("id = ?", targetId).UpdateColumn("follower_count", gorm.Expr("follower_count - 1")).Error; err != nil {
			return errmsg.ERROR_CANCEL_FAIL
		} else {
			if err := model.Db.Model(&model.User{}).Where("id = ?", userId).UpdateColumn("follow_count", gorm.Expr("follow_count - 1")).Error; err != nil {
				model.Db.Model(&model.User{}).Where("id = ?", targetId).UpdateColumn("follower_count", gorm.Expr("follower_count + 1"))
				return errmsg.ERROR_CANCEL_FAIL
			} else {
				if err := model.Db.Delete(&relation).Error; err != nil {
					model.Db.Model(&model.User{}).Where("id = ?", targetId).UpdateColumn("follower_count", gorm.Expr("follower_count + 1"))
					model.Db.Model(&model.User{}).Where("id = ?", userId).UpdateColumn("follow_count", gorm.Expr("follow_count + 1"))
					return errmsg.ERROR_CANCEL_FAIL
				} else {
					return errmsg.CANCEL_FOLLOW_SUCCESS
				}
			}
		}
	default:
		return errmsg.ERROR
	}
}

func FollowList(userId int64) (userList []model.UserInfo, ret bool) {
	var pairs []model.Relation
	if err:= model.Db.Model(&model.Relation{}).Where("follower_id = ?", userId).Find(&pairs).Error; err != nil {
		return userList, false
	}
	
	for _, pair := range pairs {
		var userInfo model.UserInfo
		if err := model.Db.Model(&model.User{}).Where("id = ?", pair.UserID).First(&userInfo).Error; err != nil {
			return userList, false
		}
		userInfo.IsFollow = true
		userList = append(userList, userInfo)
	}
	return userList, true
}

func FollowerList(userId int64) (userList []model.UserInfo, ret bool) {
	var pairs []model.Relation
	if err:= model.Db.Model(&model.Relation{}).Where("user_id = ?", userId).Find(&pairs).Error; err != nil {
		return userList, false
	}
	
	for _, pair := range pairs {
		var userInfo model.UserInfo
		if err := model.Db.Model(&model.User{}).Where("id = ?", pair.FollowerID).First(&userInfo).Error; err != nil {
			return userList, false
		}
		if model.Db.Model(&model.Relation{}).
		Where("user_id = ? and follower_id = ?", pair.FollowerID, userId).
		First(&model.Relation{}).RowsAffected > 0 {
			userInfo.IsFollow = true
		}
		userList = append(userList, userInfo)
	}
	return userList, true
}