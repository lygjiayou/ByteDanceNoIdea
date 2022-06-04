package api

import(
	"douyin/model"
	"douyin/service"
	"douyin/utils/errmsg"
	"douyin/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"fmt"
	"strings"
)

func Publish(c *gin.Context) {
	//获取token
	token := c.PostForm("token")
	//token不存在
	if token == "" {
		c.JSON(http.StatusOK, model.Response{
			StatusCode:errmsg.ERROR_TOKEN_EXIST,
			StatusMsg:errmsg.GetErrMsg(errmsg.ERROR_TOKEN_EXIST),
		})
		return
	}

	//获取user id
	//解析username
	key,err := middleware.CheckToken(token)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, model.Response{
			StatusCode:errmsg.ERROR_TOKEN_WRONG,
			StatusMsg:errmsg.GetErrMsg(errmsg.ERROR_TOKEN_WRONG),
		})
		return
	}

	//根据username获取userid
	user := model.User{
		UserName:key.Username,
	}
	userId, is_found :=  user.FindByUsername()
	if(is_found != nil) {
		c.JSON(http.StatusOK,model.Response{
			StatusCode:errmsg.ERROR_USER_NOT_EXIST,
			StatusMsg:errmsg.GetErrMsg(errmsg.ERROR_USER_NOT_EXIST),
		})
		return
	}

	//保存视频到本地public
	data,errLoad := c.FormFile("data")
	if errLoad != nil {
		c.JSON(http.StatusOK,model.Response{
			StatusCode:errmsg.ERROR_LOAD_DATA,
			StatusMsg:errmsg.GetErrMsg(errmsg.ERROR_LOAD_DATA),
		})
		return
	}

	fileName := filepath.Base(data.Filename)
	finalVideoName := fmt.Sprintf("%d_%s", userId, fileName)
	saveVideoFile := filepath.Join("./public/", finalVideoName)

	errSave := c.SaveUploadedFile(data,saveVideoFile)
	if errSave != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode:errmsg.ERROR_SAVE_VIDEO,
			StatusMsg:errmsg.GetErrMsg(errmsg.ERROR_SAVE_VIDEO),
		})
		return
	}

	//提取视频第一帧作为封面保存至pulic
	//saveCoverFile := "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg"
	videoTitle := c.PostForm("title")
	videoName := strings.Split(saveVideoFile, ".")
	saveCoverFile := strings.Join(videoName[:len(videoName)-1], "") + ".jpg"
	if saveCoverSuccess := service.SaveCover(saveVideoFile, videoTitle); saveCoverSuccess == false {
		c.JSON(http.StatusOK, model.Response{
			StatusCode:errmsg.ERROR_SAVE_COVER,
			StatusMsg:errmsg.GetErrMsg(errmsg.ERROR_SAVE_COVER),
		})
		return
	}

	//保存视频和封面信息到数据库
	success := service.SaveVideo(userId, saveVideoFile, saveCoverFile)
	if success == 0 {
		c.JSON(http.StatusOK,model.Response{
			StatusCode:errmsg.ERROR_UPDATE_DB,
			StatusMsg:errmsg.GetErrMsg(errmsg.ERROR_UPDATE_DB),
		})
		return
	}

	//上传成功
	c.JSON(http.StatusOK,model.Response{
		StatusCode:errmsg.SUCCESS,
		StatusMsg:"上传成功",
	})
}

func PublishList(c *gin.Context) {
	//获取token
	token := c.Query("token")
	//token不存在
	if token == "" {
		c.JSON(http.StatusOK, model.PublishListResponse{
			PublishResponse: model.PublishResponse{
				StatusCode: errmsg.ERROR_TOKEN_EXIST,
				StatusMsg: errmsg.GetErrMsg(errmsg.ERROR_TOKEN_EXIST),
			},
			VideoList: []model.VideoInfo{},
		})
		return
	}

	//获取user id
	//解析username
	key,err := middleware.CheckToken(token)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, model.Response{
			StatusCode:errmsg.ERROR_TOKEN_WRONG,
			StatusMsg:errmsg.GetErrMsg(errmsg.ERROR_TOKEN_WRONG),
		})
		return
	}
	//根据username获取userid
	user := model.User{
		UserName:key.Username,
	}
	userId, is_found :=  user.FindByUsername()
	if(is_found != nil) {
		c.JSON(http.StatusOK,model.Response{
			StatusCode:errmsg.ERROR_USER_NOT_EXIST,
			StatusMsg:errmsg.GetErrMsg(errmsg.ERROR_USER_NOT_EXIST),
		})
		return
	}

	//返回投稿视频
	c.JSON(http.StatusOK, model.PublishListResponse{
		PublishResponse: model.PublishResponse{
			StatusCode: errmsg.SUCCESS,
			StatusMsg: errmsg.GetErrMsg(errmsg.SUCCESS),
		},
		VideoList: service.PublishList(userId),
	})

}