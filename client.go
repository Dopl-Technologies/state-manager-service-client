package client

import (
	"context"
	"fmt"
	"strconv"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	dtprotos "github.com/dopl-technologies/api-protos-go"
)

// Client device service client
type Client struct {
	client dtprotos.StateManagerServiceClient
	conn   *grpc.ClientConn
}

// New creates a new client that connects to the
// given address
func New(address string) (Interface, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		client: dtprotos.NewStateManagerServiceClient(conn),
		conn:   conn,
	}, nil
}

// RecordFrames records frames
func (c *Client) RecordFrames(sessionID uint64, deviceID uint64) (chan<- *dtprotos.Frame, context.CancelFunc, error) {
	canceleableCtx, cancel := context.WithCancel(context.Background())
	ctx := metadata.AppendToOutgoingContext(
		canceleableCtx,
		"session-id", strconv.FormatUint(sessionID, 10),
		"device-id", strconv.FormatUint(deviceID, 10),
	)

	stream, err := c.client.RecordFrames(ctx)
	if err != nil {
		return nil, cancel, err
	}

	frameCh := make(chan *dtprotos.Frame)
	go func() {
		for {
			stop := false

			select {
			case <-ctx.Done():
				stop = true
			case frame := <-frameCh:
				req := &dtprotos.RecordFramesRequest{
					Data:    frame,
					Created: ptypes.TimestampNow(),
				}

				if err := stream.Send(req); err != nil {
					fmt.Println("Error sending record frame request:", err)
					stop = true
				}
			}

			if stop {
				break
			}
		}

		close(frameCh)
		stream.CloseAndRecv()
	}()

	return frameCh, cancel, nil
}

// Close closes the connection
func (c *Client) Close() {
	c.conn.Close()
}
