package whack

type Request struct {
	Bytes []byte
}

func (r *Request) Stdout(a ...interface{}) {
	stdout(a...)
}

func (r *Request) Stderr(a ...interface{}) {
	stderr(a...)
}

func NewRequest(ptr uintptr, length int32) *Request {
	return &Request{
		Bytes: pointerToSlice(ptr, length),
	}
}
