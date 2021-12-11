package memory

import (
	"reflect"
	"unsafe"
)

// UnsafePointer returns an unsafe.Pointer + length for the given
// byte array.
func SliceToPointer(buffer []byte) (unsafe.Pointer, int32) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	return unsafe.Pointer(header.Data), int32(len(buffer))
}

// PointerToSlice returns a byte array for the given pointer + length.
func PointerToSlice(ptr uintptr, length int32) []byte {
	buffer := make([]byte, length)
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	header.Data = ptr
	header.Len = int(length)
	header.Cap = int(length)
	return buffer
}
