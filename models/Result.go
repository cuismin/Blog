package models

//数据返回结构体设计
type Result struct {
	Message string    `json:"message"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	State   int         `json:"state"`
	Sid     string     `json:"sid"`
}

func SetResult(message string , data interface{} , success bool , state int , sid string) Result  {
	return Result{Message:message,Data:data,Success:success,State:state,Sid:sid}
}

