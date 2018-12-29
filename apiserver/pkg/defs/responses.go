package defs

//response

// ResSignedUp 注册回复
type ResSignedUp struct {
	Success   bool   `json:"success"`
	SessionID string `json:"session_id"`
}

// ResUserSession 用户session回复
type ResUserSession struct {
	Username  string `json:"user_name"`
	SessionID string `json:"session_id"`
}

// ResUserInfo 用户信息回复
type ResUserInfo struct {
	ID int64 `json:"id"`
}

// ResSignedIn 登陆回复
type ResSignedIn struct {
	Success   bool   `json:"success"`
	SessionID string `json:"session_id"`
}

// ResVideoInfos 视频信息
type ResVideoInfos struct {
	VideoInfos []Videoinfo `json:"videoinfos"`
	Count      int         `json:"total"`
}

// ResComments 评论信息
type ResComments struct {
	Comments []Comment `json:"comments"`
	Count    int       `json:"total"`
}

// ResAddComment 评论信息
type ResAddComment struct {
	Status string `json:"status"`
}

// ResAddVideoInfo 评论信息
type ResAddVideoInfo struct {
	ID string `json:"id"`
}
