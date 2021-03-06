// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package client

import (
	"context"
	"github.com/dopl-technologies/api-protos-go"
	"sync"
)

var (
	lockInterfaceMockClose        sync.RWMutex
	lockInterfaceMockRecordFrames sync.RWMutex
)

// Ensure, that InterfaceMock does implement Interface.
// If this is not the case, regenerate this file with moq.
var _ Interface = &InterfaceMock{}

// InterfaceMock is a mock implementation of Interface.
//
//     func TestSomethingThatUsesInterface(t *testing.T) {
//
//         // make and configure a mocked Interface
//         mockedInterface := &InterfaceMock{
//             CloseFunc: func()  {
// 	               panic("mock out the Close method")
//             },
//             RecordFramesFunc: func(sessionID uint64, deviceID uint64) (chan<- *dtprotos.Frame, context.CancelFunc, error) {
// 	               panic("mock out the RecordFrames method")
//             },
//         }
//
//         // use mockedInterface in code that requires Interface
//         // and then make assertions.
//
//     }
type InterfaceMock struct {
	// CloseFunc mocks the Close method.
	CloseFunc func()

	// RecordFramesFunc mocks the RecordFrames method.
	RecordFramesFunc func(sessionID uint64, deviceID uint64) (chan<- *dtprotos.Frame, context.CancelFunc, error)

	// calls tracks calls to the methods.
	calls struct {
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// RecordFrames holds details about calls to the RecordFrames method.
		RecordFrames []struct {
			// SessionID is the sessionID argument value.
			SessionID uint64
			// DeviceID is the deviceID argument value.
			DeviceID uint64
		}
	}
}

// Close calls CloseFunc.
func (mock *InterfaceMock) Close() {
	if mock.CloseFunc == nil {
		panic("InterfaceMock.CloseFunc: method is nil but Interface.Close was just called")
	}
	callInfo := struct {
	}{}
	lockInterfaceMockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	lockInterfaceMockClose.Unlock()
	mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedInterface.CloseCalls())
func (mock *InterfaceMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	lockInterfaceMockClose.RLock()
	calls = mock.calls.Close
	lockInterfaceMockClose.RUnlock()
	return calls
}

// RecordFrames calls RecordFramesFunc.
func (mock *InterfaceMock) RecordFrames(sessionID uint64, deviceID uint64) (chan<- *dtprotos.Frame, context.CancelFunc, error) {
	if mock.RecordFramesFunc == nil {
		panic("InterfaceMock.RecordFramesFunc: method is nil but Interface.RecordFrames was just called")
	}
	callInfo := struct {
		SessionID uint64
		DeviceID  uint64
	}{
		SessionID: sessionID,
		DeviceID:  deviceID,
	}
	lockInterfaceMockRecordFrames.Lock()
	mock.calls.RecordFrames = append(mock.calls.RecordFrames, callInfo)
	lockInterfaceMockRecordFrames.Unlock()
	return mock.RecordFramesFunc(sessionID, deviceID)
}

// RecordFramesCalls gets all the calls that were made to RecordFrames.
// Check the length with:
//     len(mockedInterface.RecordFramesCalls())
func (mock *InterfaceMock) RecordFramesCalls() []struct {
	SessionID uint64
	DeviceID  uint64
} {
	var calls []struct {
		SessionID uint64
		DeviceID  uint64
	}
	lockInterfaceMockRecordFrames.RLock()
	calls = mock.calls.RecordFrames
	lockInterfaceMockRecordFrames.RUnlock()
	return calls
}
