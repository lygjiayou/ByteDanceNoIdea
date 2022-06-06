package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000...用户模块的错误

	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004 // 用户token不存在
	ERROR_TOKEN_RUNTIME    = 1005 // 用户token超时
	ERROR_TOKEN_WRONG      = 1006 // 用户token错误
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008

	//上传视频模块错误
	ERROR_LOAD_DATA  = 2001
	ERROR_SAVE_VIDEO = 2002
	ERROR_SAVE_COVER = 2003
	ERROR_UPDATE_DB  = 2004
)

// CodeMsg err_code -> err_msg
var CodeMsg = map[int]string{
	SUCCESS:                "OK~",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
	ERROR_USER_NO_RIGHT:    "该用户无权限",

	ERROR_LOAD_DATA:  "提取视频数据错误",
	ERROR_SAVE_VIDEO: "保存视频数据错误",
	ERROR_SAVE_COVER: "保存视频封面错误",
	ERROR_UPDATE_DB:  "更新数据库视频信息错误",
}

func GetErrMsg(code int) string {
	return CodeMsg[code]
}
