package res

type RespSt struct {
	Errno  int         `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}
