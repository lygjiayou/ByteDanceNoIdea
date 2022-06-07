package service

import (
	"ByteDanceNoIdea/middleware"
	"ByteDanceNoIdea/model"
	"ByteDanceNoIdea/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	model.Response
	VideoList []model.VideoInfo `json:"video_list,omitempty"`
	NextTime  int64             `json:"next_time,omitempty"`
}

func FeedService(c *gin.Context) {
	// 因为要通过token来解析出username,进而获取到用户信息，所以就不用route来验证token了，但是这样要自己验证token
	//判断有没有用户登录
	token := c.Query("token")
	// 如果token不存在
	if token == "" {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: errmsg.ERROR_TOKEN_EXIST,
			StatusMsg:  errmsg.GetErrMsg(errmsg.ERROR_TOKEN_EXIST),
		})
		return
	}
	//获取user id
	//解析username
	key, err := middleware.CheckToken(token)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: errmsg.ERROR_TOKEN_WRONG,
			StatusMsg:  errmsg.GetErrMsg(errmsg.ERROR_TOKEN_WRONG),
		})
		return
	}

	//根据username获取user的全部信息
	user := model.User{
		UserName: key.Username,
	}
	_, is_found := model.FindUserInfo(user)
	if is_found != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: errmsg.ERROR_USER_NOT_EXIST,
			StatusMsg:  errmsg.GetErrMsg(errmsg.ERROR_USER_NOT_EXIST),
		})
		return
	}
	// 上面判断token的逻辑是yanghai的代码
	latest_time := c.Query("latest_time")
	//videoList := []model.VideoInfo{}
	var videoList []model.VideoInfo
	var video model.Video
	//var user model.User 上面定义过了
	strconv.Atoi(latest_time)
	//把数据库里所有视频放在videoList内,且按照创建时间降序排列
	model.Db.Order("created_time desc").Find(&video)

	// 因为用之前的video，所以下面这四行就换成下面的
	//for i := 0; i < len(videoList); i++ {
	//	user := common.User{}
	//	model.Db.Where("token = ?", videoList[i].PublisherToken).Find(&user)
	//	videoList[i].Author = user
	//}
	for i := 0; i < len(videoList); i++ {
		videoList[i].ID = video.ID
		videoList[i].CommentCount = video.CommentCount
		videoList[i].CoverUrl = video.CoverUrl
		videoList[i].FavoriteCount = video.FavoriteCount
		videoList[i].IsFavorite = false
		videoList[i].PlayUrl = video.PlayUrl
		videoList[i].Title = video.Title
		videoList[i].Author.UserName = user.UserName
		videoList[i].Author.ID = user.ID
		videoList[i].Author.IsFollow = false
		videoList[i].Author.FollowCount = user.FollowCount
		videoList[i].Author.FollowerCount = user.FollowerCount
	}
	//if token == "" {
	//	//每次获取先把默认点赞标识改为false
	//	for i := 0; i < len(videoList); i++ {
	//		model.Db.Model(&common.Video{}).Update("is_favorite", false)
	//		videoList[i].IsFavorite = false
	//	}
	//	users := []common.User{}
	//	model.Db.Find(&users)
	//	for i := 0; i < len(users); i++ {
	//		users[i].IsFollow = false
	//		model.Db.Model(&common.User{}).Update("is_follow", false)
	//	}
	//	c.JSON(http.StatusOK, FeedResponse{
	//		Response:  common.Response{StatusCode: 0},
	//		VideoList: videoList,
	//		NextTime:  time.Now().Unix(),
	//	})
	//
	//} else {

	c.JSON(http.StatusOK, FeedResponse{
		Response:  model.Response{StatusCode: 200},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	})

	//}

}
