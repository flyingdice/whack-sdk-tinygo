package whack

// RegisteredApp is global that stores the 3rd party registered app.
var RegisteredApp App

// Register tracks the 3rd party app using the SDK as
// the application to run when module is run by Whack.
func Register(a App) {
	RegisteredApp = a
}

// App is the interface 3rd party developers to implement
// to make their code executable by Whack.
type App interface {
	Main(Request) Response
}
