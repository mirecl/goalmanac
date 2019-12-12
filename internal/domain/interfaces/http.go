package interfaces

// HTTPLogger ...
type HTTPLogger interface {
	Errorf(code *int, path, format string, args ...interface{})
	Infof(code *int, format string, args ...interface{})
}
