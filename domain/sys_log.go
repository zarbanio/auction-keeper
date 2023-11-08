package domain

type SysLog struct {
	Id      int64  `json:"id"`
	Level   string `json:"string"`
	Message string `json:"message"`
	Fields  string `json:"fields"`
}
