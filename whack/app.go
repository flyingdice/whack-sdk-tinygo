package whack

import (
	"github.com/flyingdice/whack-sdk-tinygo/whack/request"
	"github.com/flyingdice/whack-sdk-tinygo/whack/response"
)

// Global variable that stores the registered app.
var RegisteredApp App = nil

// Register tracks the 3rd party app using the SDK as
// the application to run when module is run by Whack.
func Register(a App) {
	RegisteredApp = a
}

// App is the interface 3rd party developers to implement
// to make their code executable by Whack.
type App interface {
	Main(request.Request) response.Response
}
