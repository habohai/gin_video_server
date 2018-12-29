package models

import (
	"testing"

	//"github.com/haibeichina/gin_video_server/0commonpkg/comdefs"
	"github.com/haibeichina/gin_video_server/scheduler/pkg/defs"
)

var temdelrecs []string

func clearTables() {
	db.Delete(&defs.Delrec{})
}

func TestMain(m *testing.M) {
	ForTestSetUp()
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	clearTables()
	t.Run("Add", testAddDelRec)
	t.Run("Get", testGetDelRecs)
	t.Run("Del", testDeleteDelRecs)
	t.Run("Reget", testDeleteDelRec)
}

func testAddDelRec(t *testing.T) {
	if err := AddDelRec("0abc"); err != nil {
		t.Errorf("Error of AddDelRec: %v", err)
	}

	if err := AddDelRec("1abc"); err != nil {
		t.Errorf("Error of AddDelRec: %v", err)
	}

	if err := AddDelRec("2abc"); err != nil {
		t.Errorf("Error of AddDelRec: %v", err)
	}

	if err := AddDelRec("3abc"); err != nil {
		t.Errorf("Error of AddDelRec: %v", err)
	}

	if err := AddDelRec("4abc"); err != nil {
		t.Errorf("Error of AddDelRec: %v", err)
	}

	if err := AddDelRec("5abc"); err != nil {
		t.Errorf("Error of AddDelRec: %v", err)
	}
}

func testGetDelRecs(t *testing.T) {
	dr, err := GetDelRecs(3)

	if err != nil || cap(dr) == 0 {
		t.Error("Error of GetDelRecs")
	}

	for _, v := range dr {
		temdelrecs = append(temdelrecs, v.VideoID)
	}
}

func testDeleteDelRecs(t *testing.T) {
	if err := DeleteDelRecs(temdelrecs); err != nil {
		t.Errorf("Error of DeleteDelRecs: %v", err)
	}
}

func testDeleteDelRec(t *testing.T) {
	dr, err := GetDelRecs(1)

	if err != nil || cap(dr) == 0 {
		t.Error("Error of GetDelRecs")
	}

	err = DeleteDelRec(dr[0].VideoID)
	if err != nil {
		t.Errorf("Error of DeleteDelRec: %v", err)
	}
}
