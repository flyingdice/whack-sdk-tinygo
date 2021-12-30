package whack

import (
	"reflect"
	"runtime"
	"unsafe"
)

// Whack entrypoint.
//
// This is called by the Whack runtime when it's ready to run the 3rd party
// app.
//export whack_main
func main(ptr uintptr, length int32) {
	stdout("debug: received main request %d %d", ptr, length)

	// Convert input into a Req. This memory is allocated
	// and managed by Whack runtime before/after execution of the module.
	// We are not responsible for freeing it.
	req := NewRequest(ptr, length)

	// Execute all apps registered by 3rd party using the SDK.
	for _, app := range Apps() {
		res := app.Main(req)
		// Return Success/error of registered app code
		// back to the host.
		if res.Error != nil {
			HostError(res.Error)
		} else if res.Success != nil {
			HostSuccess(res.Success)
		}
	}
}

// Malloc allocates block of memory of the given size.
//
//export whack_malloc
func malloc(length int32) uintptr {
	stdout("debug: received malloc request %d", length)

	buffer := make([]byte, length)
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))

	// NOTE (hawker) This is a no-op in tinygo atm.
	runtime.KeepAlive(buffer)

	return header.Data
}

// Free releases block of memory at the given pointer + length.
//
//export whack_free
func free(ptr uintptr, length int32) {
	stdout("debug: received free request %d %d", ptr, length)

	var buffer []byte

	header := (*reflect.SliceHeader)(unsafe.Pointer(&buffer))
	header.Data = ptr
	header.Len = uintptr(length)
	header.Cap = uintptr(length)

	buffer = nil
}
