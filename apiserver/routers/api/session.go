package api

import (
	"sync"
	"time"

	"github.com/haibeichina/gin_video_server/0commonpkg/comutils"

	"github.com/haibeichina/gin_video_server/apiserver/models"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/defs"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/logging"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	models.DeleteSession(sid)
}

// LoadSessionsFromDB 从数据库中获取所有session
func LoadSessionsFromDB() {
	rs, err := models.GetSessions()
	if err != nil {
		logging.Fatal("get all session error")
		return
	}

	for _, s := range rs {
		ss := &defs.SimpleSession{
			UserName: s.UserName,
			TTL:      s.TTL,
		}
		sessionMap.Store(s.ID, ss)
	}
	return
}

// GenerateNewSessionID 生成一个session id
func GenerateNewSessionID(un string) string {
	id, _ := comutils.NewUUID()
	ct := time.Now().Unix()
	ttl := ct + 30*60

	ss := &defs.SimpleSession{
		UserName: un,
		TTL:      ttl,
	}
	sessionMap.Store(id, ss)

	data := make(map[string]interface{})
	data["id"] = id
	data["user_name"] = un
	data["ttl"] = ttl
	models.AddSession(data)

	return id
}

// IsSessionExpired 判断session是否过期
func IsSessionExpired(sid string) (string, bool) {
	ct := time.Now().Unix()
	if ss, ok := sessionMap.Load(sid); ok {
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}

		return ss.(*defs.SimpleSession).UserName, false
	}

	session, err := models.GetSession(sid)
	if err != nil {
		logging.Info("get session from DB err")
	}

	if session.TTL < ct {
		deleteExpiredSession(sid)
		return "", true
	}

	s := &defs.SimpleSession{
		UserName: session.UserName,
		TTL:      session.TTL,
	}

	sessionMap.Store(sid, s)
	return s.UserName, false
}
