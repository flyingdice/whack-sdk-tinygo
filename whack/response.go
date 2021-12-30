package whack

type Response struct {
	Success []byte
	Error   error
}

func Success(val []byte) *Response {
	return &Response{
		Success: val,
	}
}

func Error(err error) *Response {
	return &Response{
		Error: err,
	}
}
