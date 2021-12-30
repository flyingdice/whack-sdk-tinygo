package whack

// RegisteredApps is global that stores the 3rd party registered apps.
var RegisteredApps []App

// Register tracks the 3rd party app using the SDK as
// the application to run when module is run by Whack.
func Register(a App) {
	RegisteredApps = append(RegisteredApps, a)
}

// Apps lists all registered 3rd party apps.
func Apps() []App {
	return RegisteredApps
}

// App is the interface 3rd party developers to implement
// to make their code executable by Whack.
type App interface {
	Main(*Request) *Response
}
