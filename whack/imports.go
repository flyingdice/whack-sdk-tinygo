package whack

// #include <imports.h>
import "C"

// SD calls the imported SD host function with the given result.
func HostSuccess(val []byte) {
	ptr, length := sliceToPointer(val)
	C.success(ptr, length)
}

// ED calls the imported error host function with the given error.
func HostError(err error) {
	ptr, length := sliceToPointer([]byte(err.Error()))
	C.error(ptr, length)
}
