package repository

type Video struct {
	Id            int64  `gorm:"column:id"`             // 视频ID
	Title         string `gorm:"column:title"`          // 视频标题
	Name          string `gorm:"column:name"`           // 视频作者ID
	PlayUrl       string `gorm:"column:play_url"`       // 视频播放地址
	CoverUrl      string `gorm:"column:cover_url"`      // 	视频封面地址
	FavoriteCount string `gorm:"column:favorite_count"` // 视频点赞总数
	CommentCount  string `gorm:"column:comment_count"`  // 视频评论总数
	IssueTime     string `gorm:"column:issue_time"`     // 投稿时间
}

//func (Video) TableName() string {
//	return "video"
//}
//
//type VideoDao struct {
//}
//
//var videoDao *VideoDao
//var videoOnce sync.Once
//
//func NewVideoDaoInstance() *VideoDao {
//	videoOnce.Do(
//		func() {
//			videoDao = &VideoDao{}
//		})
//	return videoDao
//}
//
//func (*VideoDao) QueryPostById(id int64) (*Video, error) {
//	var video Video
//	err := db.Where("id = ?", id).Find(&video).Error
//	if err == gorm.ErrRecordNotFound {
//		return nil, nil
//	}
//	if err != nil {
//		//util.Logger.Error("find post by id err:" + err.Error())
//		return nil, err
//	}
//	return &video, nil
//}
//
//func (*VideoDao) QueryPostByParentId(parentId int64) ([]*Video, error) {
//	var videos []*Video
//	err := db.Where("parent_id = ?", parentId).Find(&videos).Error
//	if err != nil {
//		//util.Logger.Error("find posts by parent_id err:" + err.Error())
//		return nil, err
//	}
//	return videos, nil
//}
//
//func (*VideoDao) CreateVideo(video *Video) error {
//	if err := db.Create(video).Error; err != nil {
//		//util.Logger.Error("insert post err:" + err.Error())
//		return err
//	}
//	return nil
//}
