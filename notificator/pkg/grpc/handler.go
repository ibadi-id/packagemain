package grpc

import (
	"context"

	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "github.com/ibadi-id/packagemain/notificator/pkg/endpoint"
	pb "github.com/ibadi-id/packagemain/notificator/pkg/grpc/pb"
)

// makeSendEmailHandler creates the handler logic
func makeSendEmailHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SendEmailEndpoint, decodeSendEmailRequest, encodeSendEmailResponse, options...)
}

// decodeSendEmailResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain SendEmail request.
// TODO implement the decoder
func decodeSendEmailRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.SendEmailRequest)
	return endpoint.SendEmailRequest{Email: req.Email, Content: req.Content}, nil
}

// encodeSendEmailResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSendEmailResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.SendEmailResponse)
	return &pb.SendEmailReply{Id: res.Id}, nil
}
func (g *grpcServer) SendEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailReply, error) {
	_, rep, err := g.sendEmail.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SendEmailReply), nil
}
