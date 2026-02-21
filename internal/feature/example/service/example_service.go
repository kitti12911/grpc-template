package exampleserv

import (
	"context"
	"grpc-template/internal/feature/example"
)

type iExampleRepo interface {
	GetExample(ctx context.Context, id string) (*example.Example, error)
}

type exampleService struct {
	exampleRepo iExampleRepo
}

func NewExampleService(exampleRepo iExampleRepo) *exampleService {
	return &exampleService{exampleRepo: exampleRepo}
}

func (s *exampleService) GetExample(ctx context.Context, id string) (*example.Example, error) {
	return s.exampleRepo.GetExample(ctx, id)
}
