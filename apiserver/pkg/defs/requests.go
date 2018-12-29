package defs

// ReqUserCredential 用户认证
type ReqUserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

// ReqNewComment 新增评论
type ReqNewComment struct {
	UserID  int64  `json:"user_id"`
	Content string `json:"content"`
}

// ReqNewVideo 新视频信息
type ReqNewVideo struct {
	UserID    int64  `json:"user_id"`
	VideoName string `json:"video_name"`
}

// ReqVideos 请求用户的视频信息
type ReqVideos struct {
	PageNum int `json:"page_num"`
}

// ReqComments 请求用户的视频信息
type ReqComments struct {
	PageNum int `json:"page_num"`
}
