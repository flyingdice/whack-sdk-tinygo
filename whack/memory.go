package whack

import (
	"reflect"
	"unsafe"
)

// sliceToPointer returns unsafe.Pointer + length for the given
// byte array.
func sliceToPointer(buffer []byte) (unsafe.Pointer, int32) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	return unsafe.Pointer(header.Data), int32(len(buffer))
}

// pointerToSlice returns a byte array for the given pointer + length.
func pointerToSlice(ptr uintptr, length int32) []byte {
	buffer := make([]byte, length)
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	header.Data = ptr
	header.Len = uintptr(length)
	header.Cap = uintptr(length)
	return buffer
}
