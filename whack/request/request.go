package request

import "github.com/flyingdice/whack-sdk-tinygo/internal/memory"

type Request interface {
	Bytes() []byte
}

type request struct {
	bytes []byte
}

func (r *request) Bytes() []byte { return r.bytes }

func New(ptr uintptr, length int32) *request {
	return &request{
		bytes: memory.PointerToSlice(ptr, length),
	}
}
