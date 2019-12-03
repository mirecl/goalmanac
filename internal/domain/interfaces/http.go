package interfaces

// HTTPLogger ...
type HTTPLogger interface {
	Error(args ...string)
	Info(args ...string)
	Infof(format string, args ...interface{})
}
