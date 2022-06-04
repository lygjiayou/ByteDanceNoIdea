package model

// 登录
type LoginRequest struct {
	//Name string `binding:"required"`
	//Password string `binding:"required"`
	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type LoginResponse struct {
	//StatusCode int64   `json:"status_code"`// 状态码，0-成功，其他值-失败
	//StatusMsg  *string `json:"status_msg"` // 返回状态描述
	Response
	Token  string `json:"token"`             // 用户鉴权token
	UserID int64  `json:"user_id,omitempty"` // 用户id
}

// 注册
type RegisterRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// 根据user_id和token查询用户
type UserRequest struct {
	UserID int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}

// user的Response 下面有定义
type UserInfoResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	User       User   `json:"user"`        // 用户信息
}

type RegisterResponse struct {
	//StatusCode int64   `json:"status_code"`// 状态码，0-成功，其他值-失败
	//StatusMsg  *string `json:"status_msg"` // 返回状态描述
	Response
	Token  string `json:"token"`   // 用户鉴权token
	UserID int64  `json:"user_id"` // 用户id
}

type ErrNo int

const (
	OK             ErrNo = 0
	ParamInvalid   ErrNo = 1 // 参数不合法
	UserHasExisted ErrNo = 2 // 该 Username 已存在
	UserHasDeleted ErrNo = 3 // 用户已删除
	UserNotExisted ErrNo = 4 // 用户不存在
	WrongPassword  ErrNo = 5 // 密码错误
	LoginRequired  ErrNo = 6 // 用户未登录

	UnknownError ErrNo = 255 // 未知错误
)

// 关注操作等等会使用到
//type Response struct {
//	StatusCode int32  `json:"status_code"`
//	StatusMsg  string `json:"status_msg,omitempty"`
//}

//转移到model里
//type Video struct {
//	Id            int64  `json:"id,omitempty"`
//	Author        User   `json:"author"`
//	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
//	CoverUrl      string `json:"cover_url,omitempty"`
//	FavoriteCount int64  `json:"favorite_count,omitempty"`
//	CommentCount  int64  `json:"comment_count,omitempty"`
//	IsFavorite    bool   `json:"is_favorite,omitempty"`
//}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

// 用户信息，关注列表，粉丝列表操作会用到
type UserResponse struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

// 视频流
type VideoApifoxModal struct {
	NextTime   *int64  `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 视频列表
}

// 用户注册和登录
type UserApifoxModal struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	Token      string `json:"token"`       // 用户鉴权token
	UserID     int64  `json:"user_id"`     // 用户id
}

// 用户信息
type UserInfoApifoxModal struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	User       User    `json:"user"`        // 用户信息
}

// 投稿接口
type ContributeApifoxModal struct {
	StatusCode int64   `json:"status_code"`
	StatusMsg  *string `json:"status_msg"`
}

type PublishListRequest struct {
	UserID int64 `json:"user_id"`
}

// 发布列表，以及点赞列表用到这个了（点赞列表之一）
type PublishListResponse struct {
	StatusCode int64      `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string     `json:"status_msg"`  // 返回状态描述
	VideoList  []ResVideo `json:"video_list"`  // 用户发布的视频列表
}

// Video，以及点赞列表用到这个了（点赞列表之一）
type PublishVideo struct {
	Author        User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}

// 赞操作
type PraiseApifoxModal struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

// 点赞操作
type LikeApifoxModal struct {
	StatusCode string  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 用户点赞视频列表
}

// 点赞操作-Video
type LikeVideo struct {
	Author        User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}

// 点赞操作-视频作者信息
type LikeUser struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	ID            int64  `json:"id"`             // 用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Name          string `json:"name"`           // 用户名称
}

// 评论操作
type CommentApifoxModal struct {
	Comment    Comment `json:"comment"`     // 评论成功返回评论内容，不需要重新拉取整个列表
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
}

// 评论列表
type CommentListApifoxModal struct {
	CommentList []Comment `json:"comment_list"` // 评论列表
	StatusCode  int64     `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   *string   `json:"status_msg"`   // 返回状态描述
}

//
//// 关注操作
//type ApifoxModal struct {
//	StatusCode int64  `json:"status_code"`// 状态码，0-成功，其他值-失败
//	StatusMsg  string `json:"status_msg"` // 返回状态描述
//}

// 关注列表，粉丝列表会用到
type FollowApifoxModal struct {
	StatusCode string  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	UserList   []User  `json:"user_list"`   // 用户信息列表
}
