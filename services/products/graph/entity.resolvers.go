package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ContextLogic/products/graph/generated"
	"github.com/ContextLogic/products/graph/model"
)

func (r *entityResolver) FindProductByUpc(ctx context.Context, upc string) (*model.Product, error) {
	fmt.Printf("FindProductByUpc: %s\n", upc)
	for _, product := range products {
		if product.Upc == upc {
			return product, nil
		}
	}
	return nil, nil
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
var products []*model.Product

func init() {
	products = []*model.Product{
		{
			Upc:    "1",
			Name:   "Table",
			Price:  899,
			Weight: 100,
		},
		{
			Upc:    "2",
			Name:   "Couch",
			Price:  1299,
			Weight: 1000,
		},
		{
			Upc:    "3",
			Name:   "Chair",
			Price:  54,
			Weight: 50,
		},
	}
}
