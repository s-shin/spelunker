package error

import "google.golang.org/grpc"

// Error codes
const (
	EcodeInvalidMessage = iota
	EcodeUnsupported
	EcodeUnauthorized
)

var errorMessages = map[int]string{
	EcodeInvalidMessage: "invalid message",
	EcodeUnsupported:    "unsupported",
	EcodeUnauthorized:   "unauthorized",
}

func CreateGrpcError(code int) error {
	return grpc.Errorf(code, errorMessages[code])
}
