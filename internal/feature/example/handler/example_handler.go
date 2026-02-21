package examplehand

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	examplev1 "grpc-template/gen/grpc/example/v1"
	"grpc-template/internal/feature/example"
)

type iExampleServ interface {
	GetExample(ctx context.Context, id string) (*example.Example, error)
}

type ExampleHandler struct {
	examplev1.UnimplementedExampleServiceServer
	exampleService iExampleServ
}

func NewExampleHandler(exampleService iExampleServ) *ExampleHandler {
	return &ExampleHandler{exampleService: exampleService}
}

func (h *ExampleHandler) GetExample(ctx context.Context, req *examplev1.GetExampleRequest) (*examplev1.GetExampleResponse, error) {
	result, err := h.exampleService.GetExample(ctx, req.GetId())

	if err != nil {
		return nil, err
	}

	return &examplev1.GetExampleResponse{
		Example: &examplev1.Example{
			Id:          result.ID,
			Name:        result.Name,
			Description: result.Description,
			Status:      result.Status,
			CreatedAt:   timestamppb.New(result.CreatedAt),
			UpdatedAt:   timestamppb.New(result.UpdatedAt),
		},
	}, nil
}
