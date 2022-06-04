package service

import (
	"douyin/model"
	"douyin/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoListResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list"`
}

// PublishList 返回用户发布列表
func PublishList(c *gin.Context) {
	userId := c.Query("user_id")
	videos, err := repository.GetPublishList(userId)
	var resp = Response{}
	if err != nil {
		resp.StatusCode = 0
		resp.StatusMsg = err.Error()
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response:  resp,
		VideoList: videos,
	})
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
