package apihandlerservice

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/haibeichina/gin_video_server/web/pkg/setting"
)

// var httpClient *http.Client

// func init() {
// 	httpClient = &http.Client{}
// }

// APIHander 透传用处理
type APIHandler struct {
	URL     string
	Method  string
	ReqBody string

	HTTPRequest *http.Request
	HTTPClient  *http.Client
}

// APIHandlerRequest api转发请求
func (a *APIHandler) APIHandlerRequest() (*http.Response, error) {
	var resp *http.Response
	var err error

	u, _ := url.Parse(a.URL)
	u.Host = setting.AppSetting.APIAddr + ":" + u.Port()
	newURL := u.String()

	switch a.Method {
	case http.MethodGet:
		req, _ := http.NewRequest("GET", newURL, nil)
		req.Header = a.HTTPRequest.Header
		resp, err = a.HTTPClient.Do(req)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	case http.MethodPost:
		req, _ := http.NewRequest("POST", newURL, bytes.NewBuffer([]byte(a.ReqBody)))
		req.Header = a.HTTPRequest.Header
		resp, err = a.HTTPClient.Do(req)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	case http.MethodDelete:
		req, _ := http.NewRequest("DELETE", newURL, nil)
		req.Header = a.HTTPRequest.Header
		resp, err = a.HTTPClient.Do(req)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	default:
		return nil, errors.New("no match method")
	}

	return resp, nil
}

// func normalResponse(w http.ResponseWriter, r *http.Response) {
// 	res, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		re, _ := json.Marshal(ErrorInternalFaults)
// 		w.WriteHeader(500)
// 		io.WriteString(w, string(re))
// 		return
// 	}

// 	w.WriteHeader(r.StatusCode)
// 	io.WriteString(w, string(res))
// }
