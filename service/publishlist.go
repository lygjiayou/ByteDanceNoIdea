package service

import (
	"ByteDanceNoIdea/model"
)

// PublishListService  返回用户发布视频列表
func PublishListService(req model.PublishListRequest) []model.ResVideo {
	userID := req.UserID

	//获取author信息
	var author model.User
	author.ID = userID
	author.FindUserInfoByID()

	videos := model.GetPublishList(userID)
	resVideos := make([]model.ResVideo, len(videos))
	for i := 0; i < len(videos); i++ {
		resVideos[i].User = author
		resVideos[i].Video = videos[i]
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
