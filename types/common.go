package types

type BaseResponse struct {
	Code       int         `json:"bk_error_code"`
	Message    string      `json:"bk_error_msg"`
	Permission interface{} `json:"permission"`
	Result     bool        `json:"result"`
}

type PostResponse struct {
	BaseResponse
	Data map[string]interface{} `json:"data"`
}
