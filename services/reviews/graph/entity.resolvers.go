package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ContextLogic/reviews/graph/generated"
	"github.com/ContextLogic/reviews/graph/model"
)

func (r *entityResolver) FindProductByUpc(ctx context.Context, upc string) (*model.Product, error) {
	return &model.Product{Upc: upc}, nil
}

func (r *entityResolver) FindReviewByID(ctx context.Context, id string) (*model.Review, error) {
	for _, review := range reviews {
		if review.Author.ID == id {
			return review, nil
		}
	}
	return nil, nil
}

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	return &model.User{ID: id}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
