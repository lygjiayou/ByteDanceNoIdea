package service

import (
	"douyin/common"
	"douyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

func FeedService(c *gin.Context) {
	//判断有没有用户登录
	token := c.Query("token")
	videoList := []common.Video{}
	//把数据库里所有视频放在videoList内,且按照创建时间降序排列
	model.Db.Order("created_at desc").Find(&videoList)

	for i := 0; i < len(videoList); i++ {
		user := common.User{}
		model.Db.Where("token = ?", videoList[i].PublisherToken).Find(&user)
		videoList[i].Author = user
	}

	if token == "" {
		//每次获取先把默认点赞标识改为false
		for i := 0; i < len(videoList); i++ {
			model.Db.Model(&common.Video{}).Update("is_favorite", false)
			videoList[i].IsFavorite = false
		}
		users := []common.User{}
		model.Db.Find(&users)
		for i := 0; i < len(users); i++ {
			users[i].IsFollow = false
			model.Db.Model(&common.User{}).Update("is_follow", false)
		}
		c.JSON(http.StatusOK, FeedResponse{
			Response:  common.Response{StatusCode: 0},
			VideoList: videoList,
			NextTime:  time.Now().Unix(),
		})

	} else {

		c.JSON(http.StatusOK, FeedResponse{
			Response:  common.Response{StatusCode: 0},
			VideoList: videoList,
			NextTime:  time.Now().Unix(),
		})

	}

}
