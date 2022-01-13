package whack

type Request struct {
	Bytes []byte
}

func (r *Request) Stdout(s string, a ...interface{}) {
	stdout(s, a...)
}

func (r *Request) Stderr(s string, a ...interface{}) {
	stderr(s, a...)
}

func NewRequest(ptr uintptr, length int32) *Request {
	return &Request{
		Bytes: pointerToSlice(ptr, length),
	}
}
