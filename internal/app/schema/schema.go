package schema

// ErrorItem 响应错误项
type ErrorItem struct {
	Code    int    `json:"code"`    // 错误码
	Message string `json:"message"` // 错误信息
}
// ErrorResult 响应错误
type ErrorResult struct {
	Error ErrorItem `json:"error"` // 错误项
}