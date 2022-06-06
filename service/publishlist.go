package service

import (
	"ByteDanceNoIdea/model"
)

// PublishListService  返回用户发布视频列表
func PublishListService(req model.PublishListRequest) []model.VideoInfo {
	userID := req.UserID

	//获取数据库内author信息
	var authorOrigin model.User

	//即将封装入结果的author信息
	var authorInfo model.UserInfo

	authorOrigin.ID = userID
	authorOrigin.FindUserInfoByID()

	//authorOrigin -> authorInfo
	authorInfo.ID = authorOrigin.ID
	authorInfo.UserName = authorOrigin.UserName
	authorInfo.FollowCount = authorOrigin.FollowCount
	authorInfo.FollowerCount = authorOrigin.FollowerCount
	authorInfo.IsFollow = false //temp

	videos := model.GetPublishList(userID)
	resVideos := make([]model.VideoInfo, len(videos))
	for i := 0; i < len(videos); i++ {
		//videos -> resVideos
		resVideos[i].ID = videos[i].ID
		resVideos[i].Author = authorInfo
		resVideos[i].PlayUrl = videos[i].PlayUrl
		resVideos[i].CoverUrl = videos[i].CoverUrl
		resVideos[i].FavoriteCount = videos[i].FavoriteCount
		resVideos[i].CommentCount = videos[i].CommentCount
		resVideos[i].IsFavorite = false //temp
		resVideos[i].Title = videos[i].Title
	}

	return resVideos
}

//// Publish check token then save upload file to public directory
//func Publish(c *gin.Context) {
//	token := c.PostForm("token")
//
//	if _, exist := usersLoginInfo[token]; !exist {
//		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
//		return
//	}
//
//	data, err := c.FormFile("data")
//	if err != nil {
//		c.JSON(http.StatusOK, Response{
//			StatusCode: 1,
//			StatusMsg:  err.Error(),
//		})
//		return
//	}
//
//	filename := filepath.Base(data.Filename)
//	user := usersLoginInfo[token]
//	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
//	saveFile := filepath.Join("./public/", finalName)
//	if err := c.SaveUploadedFile(data, saveFile); err != nil {
//		c.JSON(http.StatusOK, Response{
//			StatusCode: 1,
//			StatusMsg:  err.Error(),
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, Response{
//		StatusCode: 0,
//		StatusMsg:  finalName + " uploaded successfully",
//	})
//}
