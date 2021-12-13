package whack

import (
	"github.com/flyingdice/whack-sdk-tinygo/whack/imports"
)

// Malloc allocates block of memory of the given size.
//export whack_main
func whack_main(ptr uintptr, length int32) {
	// Convert input into a request. This memory is allocated
	// and managed by Whack before/after  execution of the module.
	//
	// We are NOT responsible for freeing it.
	req := NewRequest(ptr, length)

	// Execute all apps registered by 3rd party using the SDK.
	for _, app := range Apps() {
		res := app.Main(req)
		// Return success/error of registered app code
		// back to the host.
		if res.Error() != nil {
			imports.Error(res.Error())
		} else {
			imports.Success(res.Success())
		}
	}
}
