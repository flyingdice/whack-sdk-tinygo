package imports

// #include <host.h>
import "C"
import (
	"github.com/flyingdice/whack-sdk-tinygo/internal/memory"
)

// Success calls the imported success host function with the given result.
func Success(val []byte) {
	ptr, length := memory.SliceToPointer(val)
	C.success(ptr, length)
}

// Error calls the imported error host function with the given error.
func Error(err error) {
	ptr, length := memory.SliceToPointer([]byte(err.Error()))
	C.error(ptr, length)
}
