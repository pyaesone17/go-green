package core

// Bundle is the app bundle interface
type Bundle interface {
	GetRoutes() []Route
}
