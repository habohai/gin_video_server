package models

import (
	"fmt"
	"testing"
	"time"

	//"github.com/haibeichina/gin_video_server/0commonpkg/comdefs"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/defs"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/logging"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/setting"
)

var tempvid string
var tempuid int64

func clearTables() {
	db.Delete(&defs.User{})
	db.Delete(&defs.Comment{})
	db.Delete(&defs.Videoinfo{})
	db.Delete(&defs.Session{})
}

func TestMain(m *testing.M) {
	setting.SetUp()
	ForTestSetUp()
	logging.SetUp()
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	clearTables()
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	data := map[string]interface{}{
		"user_name": "haibei",
		"pwd":       "123",
	}
	if err := AddUser(data); err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	user, err := GetUserInfo("haibei")

	if err != nil || user.ID == 0 {
		t.Error("Error of GetUser")
	}
	tempuid = user.ID
}

func testDeleteUser(t *testing.T) {
	data := map[string]interface{}{
		"user_name": "haibei",
		"pwd":       "123",
	}
	if err := DelUserInfo(data); err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	user, err := GetUserInfo("haibei")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}

	if user.Pwd != "" {
		t.Error("RegetUser test failed!")
	}
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	data := map[string]interface{}{
		"user_id":    tempuid,
		"user_name":  "haibei",
		"video_name": "my_video",
	}
	vi, err := AddVideoInfo(data)
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}

	tempvid = vi.ID
}

func testGetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi.ID == "" {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi.ID != "" {
		t.Errorf("Error of RegetVideoInfo err : %v , vi : %v", err, vi)
	}
}

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", testListComments)
}

func testAddComments(t *testing.T) {
	data := map[string]interface{}{
		"user_id":   tempuid,
		"user_name": "haibei",
		"video_id":  "123456",
		"content":   "I like this video",
	}

	comment, err := AddComment(data)
	if err != nil {
		t.Errorf("Error of Addcomments: %v", err)
	}

	t.Log(comment)
}

func testListComments(t *testing.T) {
	data := map[string]interface{}{
		"video_id": "123456",
	}

	res2, err2 := GetComment("123456")
	if err2 != nil {
		t.Errorf("Error of GetComment err : %v", err2)
	}

	if res2.ID == "" {
		t.Errorf("Error of GetComment res: %v", err2)
	}

	t.Log(res2)

	timef := time.Unix(0, 0).Unix()
	timet := time.Now().Unix()
	res1, err1 := GetComments("123456", timef, timet)
	if err1 != nil || len(res1) == 0 {
		t.Errorf("Error of GetComments err : %v, res: %v", err1, res1)
	}

	res, err := GetCommentsByPage(0, setting.AppSetting.PageSize, data)
	if err != nil || len(res) == 0 {
		t.Errorf("Error of ListComments err : %v, res: %v", err, res)
	}

	for i, ele := range res {
		fmt.Printf("comment:%d, %v \n", i, ele)
	}
}
