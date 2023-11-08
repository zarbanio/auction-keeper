package domain

type SysLog struct {
	Id      int64  `json:"id"`
	Level   string `json:"string"`
	Service string `json:"service"`
	Message string `json:"message"`
	Fields  string `json:"fields"`
}
