package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	ERROR_EXIST_VIDEO:     "已经存在该视频",
	ERROR_EXIST_USER_NAME: "用户名已经存在",

	ERROR_NOT_EXIST_VIDEO:    "该视频不存在",
	ERROR_NOT_EXIST_USERID:   "该用户不存在",
	ERROR_NOT_EXIST_USERNAME: "该用户名已经存在",

	ERROR_AUTH_INVALID_SESSION:   "session已经无效",
	ERROR_AUTH_PWD_UNMATCH:       "密码不匹配",
	ERROR_AUTH_USERNAME_MISMATCH: "用户名不匹配",

	ERROR_ADD_COMMENT_FAIL:   "添加评论失败",
	ERROR_ADD_USER_FAIL:      "添加用户失败",
	ERROR_ADD_VIDEOINFO_FAIL: "添加视频信息错误",
	ERROR_ADD_DELRECS_FAIL:   "添加应删除的视频信息失败",

	ERROR_GET_COMMENTS_FAIL:   "获取评论失败",
	ERROR_GET_VIDEOINFOS_FAIL: "获取用户视频失败",

	ERROR_UPLOAD_VIDEO_FILE_TOO_BIG:       "上传的视频文件太大了",
	ERROR_UPLOAD_VIDEO_FORMAT_UNSUPPORTED: "上传的视频文件格式不正确",
	ERROR_UPLOAD_VIDEO_SAVE_FAIL:          "上传的视频文件保存失败",

	ERROR_REQUEST_NOT_RECOGNIZED:             "请求没有认证",
	ERROR_REQUEST_BODY_PARSE_FAILED:          "请求body解析失败",
	ERROR_REQUEST_MISMATCH_USERNAME_URL_BODY: "请求中url和body中的username不一致",
	ERROR_REQUEST_UNMARSHAL_BODY_FAILED:      "解析请求中json格式的body失败",
	ERROR_REQUEST_BODY_UNMARSHAL_JSON_FAIL:   "http请求中body中json数据解析错误",

	ERROR_DELETE_USER_FAIL:      "删除用户失败",
	ERROR_DELETE_VIDEOINFO_FAIL: "删除视频信息失败",

	ERROR_MODIFY_USER_FAIL: "修改用户失败",

	ERROR_FIND_USER_FAIL: "查找用户失败",

	ERROR_COUNT_COMMENT_FAIL:   "获取评论数量失败",
	ERROR_COUNT_VIDEOINFO_FAIL: "获取视频信息数量失败",

	ERROR_DB_ERROR_INFO:  "数据库错误！",
	ERROR_OPEN_FILE_FAIL: "打开文件失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
