package model

type Result struct {
	Message string      `json:"message"`
	Code    string      `json:"code"`
	Data    interface{} `json:"data"`
}

func OfSeccess(data interface{}) *Result {
	return &Result{"成功", "0", data}
}

func OfFailed() *Result {
	return &Result{"失败", "0", nil}
}
