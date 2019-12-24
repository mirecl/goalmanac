package interfaces

import "google.golang.org/grpc/codes"

// GRPCLogger ...
type GRPCLogger interface {
	Errorf(status *codes.Code, path, format string, args ...interface{})
	Infof(status *codes.Code, format string, args ...interface{})
}
