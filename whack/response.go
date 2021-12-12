package whack

type Response interface {
	Success() []byte
	Error() error
}

type response struct {
	success []byte
	err     error
}

func (r *response) Success() []byte { return r.success }
func (r *response) Error() error    { return r.err }

func Success(val []byte) *response {
	return &response{
		success: val,
	}
}

func Error(err error) *response {
	return &response{
		err: err,
	}
}
