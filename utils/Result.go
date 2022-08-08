package utils

type Result struct {
	Code    int         `json:"code"`
	Count   int         `json:"count"`
	Message interface{} `json:"message"`
	Data    string      `json:"data"`
}

func (r *Result) Ok(msg string) (result *Result) {
	r.Code = 200
	r.Message = msg
	return r
}
