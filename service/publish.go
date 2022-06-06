package service

import (
	"ByteDanceNoIdea/model"
	"bytes"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

//保存视频信息到数据库
func SaveVideo(userID int64, videoUrl string, coverUrl string) int64 {
	//获取ip
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return 0
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip := localAddr.IP.String()

	//更新数据库
	video := model.Video{
		AuthorID:   userID,
		PlayUrl:    fmt.Sprintf("%s://%s:%s/%s", "http", ip, "8080", videoUrl),
		CoverUrl:   fmt.Sprintf("%s://%s:%s/%s", "http", ip, "8080", coverUrl),
		CreateTime: time.Now().Unix(),
	}
	return model.Db.Create(&video).RowsAffected
}

//保存视频封面
func SaveCover(videoUrl string, title string) bool {
	videoName := strings.Split(videoUrl, ".")
	coverPic := strings.Join(videoName[:len(videoName)-1], "") + ".jpg"
	cmdStr1 := fmt.Sprintf("ffmpeg -ss 0.5 -i %s -vframes 1 -s 720x1080 -f image2 %s", videoUrl, coverPic)
	args := strings.Fields(cmdStr1)

	_, err := Cmd(args[0], args[1:])
	if err != nil {
		return false
	}
	return true
}

//命令行接口
func Cmd(commandName string, params []string) (string, error) {
	//fmt.Println("命令行调用")
	cmd := exec.Command(commandName, params...)
	//fmt.Println("Cmd", cmd.Args)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	return out.String(), err
}

//查找userid对应的发布视频列表
func PublishList(userid int64) []model.VideoInfo {
	var videos []model.Video
	model.Db.Model(&model.Video{}).Where("author_id = ?", userid).Find(&videos)
	videoInfos := make([]model.VideoInfo, 0)
	for _, video := range videos {
		var author model.User
		model.Db.Model(&model.User{}).Where("id = ?", video.AuthorID).First(&author)
		var authorInfo model.UserInfo
		//authorOrigin -> authorInfo
		authorInfo.ID = author.ID
		authorInfo.UserName = author.UserName
		authorInfo.FollowCount = author.FollowCount
		authorInfo.FollowerCount = author.FollowerCount
		authorInfo.IsFollow = false //temp

		videoInfos = append(videoInfos, model.VideoInfo{
			ID:            video.ID,
			Author:        authorInfo,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
		})
	}
	return videoInfos
}
