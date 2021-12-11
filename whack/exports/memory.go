package exports

import (
	"reflect"
	"runtime"
	"unsafe"
)

// Malloc allocates block of memory of the given size.
//
//export malloc
func malloc(length int32) uintptr {
	buffer := make([]byte, length)
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	runtime.KeepAlive(buffer)
	return header.Data
}

// Free releases block of memory at the given pointer + length.
//
//export free
func free(ptr uintptr, length int32) {
	var buffer []byte

	header := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	header.Data = ptr
	header.Len = int(length)
	header.Cap = int(length)

	buffer = nil
}
