package interfaces

// HTTPLogger ...
type HTTPLogger interface {
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
}
