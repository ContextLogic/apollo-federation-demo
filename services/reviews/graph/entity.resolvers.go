package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ContextLogic/reviews/graph/generated"
	"github.com/ContextLogic/reviews/graph/model"
)

func (r *entityResolver) FindProductByUpc(ctx context.Context, upc string) (*model.Product, error) {
	fmt.Printf("FindProductByUpc: %s\n", upc)
	return &model.Product{Upc: upc}, nil
}

func (r *entityResolver) FindReviewByID(ctx context.Context, id string) (*model.Review, error) {
	fmt.Printf("FindReviewByID: %s\n", id)
	for _, review := range reviews {
		if review.Author.ID == id {
			return review, nil
		}
	}
	return nil, nil
}

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	fmt.Printf("FindUserByID: %s\n", id)
	return &model.User{ID: id}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var reviews []*model.Review

func init() {
	usernames := []string{"@ada", "@complete"}
	reviewBodies := []string{
		"Love it!",
		"Too expensive",
		"Could be better.",
		"Prefer something else.",
	}
	for i, r := range reviewBodies {
		reviews = append(reviews, &model.Review{
			ID: fmt.Sprintf("%d", i + 1),
			Product: &model.Product{
				Upc: fmt.Sprintf("%d", i + 1),
			},
			Author: &model.User{
				ID:       fmt.Sprintf("%d", i/2+1),
				Username: usernames[i/2],
			},
			Body: r,
		})
	}
	for _, r := range reviews {
		fmt.Println(r.ID, r.Product.Upc, r.Author.ID, r.Author.Username, r.Body)
	}
}
