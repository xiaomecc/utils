package errorx

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg error) error {
	return &CodeError{Code: code, Msg: msg.Error()}
}

func (e *CodeError) Error() string {
	return e.Msg
}

func NewDefaultError(msg error) error {
	return NewCodeError(ErrCode.E_FAILED, msg)
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
