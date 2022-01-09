// imports.h
//
// This header file defines the Whack host SDK that is injected
// into the WASM instance via import object.
//
// These functions are available via "env.${func}".
#include <stdint.h>

// host_success sends value from successful function call back to Whack host runtime.
void host_success(void* ptr, int32_t length);

// host_error sends error from failed function call back to the Whack host runtime.
void host_error(void* ptr, int32_t length);
