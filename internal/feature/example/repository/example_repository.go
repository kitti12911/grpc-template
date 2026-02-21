package examplerepo

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"grpc-template/internal/apperror"
	"grpc-template/internal/feature/example"
)

type exampleRepository struct{}

func NewExampleRepository() *exampleRepository {
	return &exampleRepository{}
}

func (r *exampleRepository) GetExample(ctx context.Context, id string) (*example.Example, error) {
	if id == "notfound" {
		return nil, apperror.NotFound("example not found", nil)
	}

	if id == "dberror" {
		sqlErr := fmt.Errorf("pq: relation \"examples\" does not exist")

		slog.ErrorContext(ctx, "failed to query example",
			"error", sqlErr.Error(),
			"id", id,
		)

		return nil, apperror.Internal("failed to retrieve example", sqlErr)
	}

	return &example.Example{
		ID:          id,
		Name:        "Example",
		Description: "Example description",
		Status:      "active",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
