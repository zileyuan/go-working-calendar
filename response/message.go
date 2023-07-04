package response

type StatusType int

const (
	StatusOK            StatusType = iota // 成功
	StatusFailure                         // 失败
	StatusMaintainError                   // 维护数据错误
	StatusParamsError                     // 参数错误
)

// Message Respone响应对象
type Message struct {
	Status StatusType  `json:"status"`
	Data   interface{} `json:"data"`
}
