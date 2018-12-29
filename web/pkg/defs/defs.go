package defs

// HomePage 替换到home页面模版上的对象name
type HomePage struct {
	Name string
}

// UserPage 替换到user页面模版上的对象name
type UserPage struct {
	Name string
}

// APIBody api 透传用
type APIBody struct {
	URL     string `json:"url"`
	Method  string `json:"method"`
	ReqBody string `json:"req_body"`
}

// // Err 错误结构
// type Err struct {
// 	Error     string `json:"error"`
// 	ErrorCode string `json:"error_code"`
// }

// var (
// 	// ErrorRequestNotRecognized 没有认证错误
// 	ErrorRequestNotRecognized = Err{
// 		Error:     "api not recognized, bad request",
// 		ErrorCode: "001",
// 	}

// 	// ErrorRequestBodyParseFailed 请求页面body解析错误
// 	ErrorRequestBodyParseFailed = Err{
// 		Error:     "request body is not correct",
// 		ErrorCode: "002",
// 	}

// 	// ErrorInternalFaults 内部错误
// 	ErrorInternalFaults = Err{
// 		Error:     "internal server errror",
// 		ErrorCode: "003",
// 	}
// )
