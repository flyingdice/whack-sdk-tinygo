package exports

import (
	"github.com/flyingdice/whack-sdk-tinygo/whack"
	"github.com/flyingdice/whack-sdk-tinygo/whack/imports"
	"github.com/flyingdice/whack-sdk-tinygo/whack/request"
)

// Entrypoint called by the Whack host execution engine
// that is responsible shuffling bytes to/from the 3rd party
// registered app code.
//
//export whack_main
func whack_main(ptr uintptr, length int32) {
	// Convert input into a request. This memory is allocated
	// and managed by Whack before/after  execution of the module.
	// We are not responsible for freeing it.
	req := request.New(ptr, length)

	// Run app code registered by 3rd party using the SDK.
	res := whack.RegisteredApp.Main(req)

	// Return success/error of registered app code
	// back to the host.
	if res.Error() != nil {
		imports.Error(res.Error())
	} else {
		imports.Success(res.Success())
	}
}
