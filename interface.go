package client

import (
	"context"

	dtprotos "github.com/dopl-technologies/api-protos-go"
)

// Interface for communicating with the state manager service
//go:generate moq -out client_mock.go . Interface
type Interface interface {
	RecordFrames(sessionID uint64, deviceID uint64) (chan<- *dtprotos.Frame, context.CancelFunc, error)

	Close()
}
