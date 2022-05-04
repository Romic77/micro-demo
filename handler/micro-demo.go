package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	microdemo "micro-demo/proto"
)

type MicroDemo struct{}

// Return a new handler
func New() *MicroDemo {
	return &MicroDemo{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *MicroDemo) Call(ctx context.Context, req *microdemo.Request, rsp *microdemo.Response) error {
	log.Info("Received MicroDemo.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *MicroDemo) Stream(ctx context.Context, req *microdemo.StreamingRequest, stream microdemo.MicroDemo_StreamStream) error {
	log.Infof("Received MicroDemo.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&microdemo.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *MicroDemo) PingPong(ctx context.Context, stream microdemo.MicroDemo_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&microdemo.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
