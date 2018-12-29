package defs

//Data model

// // VideoInfo 视频信息结构
// type VideoInfo struct {
// 	ID        string `json:"id"`
// 	UserID    int    `json:"author_id"`
// 	UserName  string `json:"user_name"`
// 	VideoName string `json:"video_name"`
// }

// // Comment 评论结构
// type Comment struct {
// 	ID           string `json:"id"`
// 	UserID       int    `json:"user_id"`
// 	UserName     string `json:"user_name"`
// 	VideoID      string `json:"video_id"`
// 	Content      string `json:"content"`
// 	DisplayCtime string `json:"display_ctime"`
// }

// SimpleSession session结构
type SimpleSession struct {
	UserName string //login name
	TTL      int64
}

// Session 视频信息表结构
type Session struct {
	ID       string `gorm:"primary_key" json:"id"`
	UserName string `json:"user_name"`
	TTL      int64  `json:"ttl"`
}

// User 认证表结构
type User struct {
	ID       int64  `gorm:"primary_key" json:"id"`
	UserName string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

// Delrec 删除记录表（用于硬删除）
type Delrec struct {
	VideoID string `json:"video_id"`
}

// Model 公共
type Model struct {
	ID        string `gorm:"primary_key" json:"id"`
	CreatedOn int64  `json:"created_on"`
}

// Comment 	评论
type Comment struct {
	Model
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	VideoID  string `json:"video_id"`
	Content  string `json:"content"`
}

// Videoinfo 视频信息表结构
type Videoinfo struct {
	Model
	UserID    int64  `json:"user_id"`
	UserName  string `json:"user_name"`
	VideoName string `json:"video_name"`
}
