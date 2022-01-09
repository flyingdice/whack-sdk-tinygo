package whack

// #include <host.h>
import "C"

// HostSuccess calls the imported SD host function with the given result.
func HostSuccess(val []byte) {
	ptr, length := sliceToPointer(val)
	C.host_success(ptr, length)
}

// HostError calls the imported error host function with the given error.
func HostError(err error) {
	ptr, length := sliceToPointer([]byte(err.Error()))
	C.host_error(ptr, length)
}
