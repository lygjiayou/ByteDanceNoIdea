package repository

type Comment struct {
	Id          int64  `gorm:"column:id"`           // 视频评论ID
	VideoId     string `gorm:"column:video_id"`     // 视频ID
	CommenterId string `gorm:"column:commenter_id"` // 评论用户ID
	Content     string `gorm:"column:content"`      // 评论内容
	CreateTime  string `gorm:"column:create_time"`  // 评论发布日期
}

//func (Comment) TableName() string {
//	return "comment"
//}
//
//type CommentDao struct {
//}
//
//var commentDao *CommentDao
//var commentOnce sync.Once
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
