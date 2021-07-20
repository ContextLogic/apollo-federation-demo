package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/ContextLogic/inventory/graph/generated"
	"github.com/ContextLogic/inventory/graph/model"
)

func (r *productResolver) InStock(ctx context.Context, obj *model.Product) (bool, error) {
	for _, p := range products {
		if obj.Upc == p.Upc {
			return p.InStock, nil
		}
	}
	return false, nil
}

func (r *productResolver) ShippingEstimate(ctx context.Context, obj *model.Product) (float64, error) {
	for _, p := range products {
		if obj.Upc == p.Upc {
			return p.ShippingEstimate, nil
		}
	}
	return 0, nil
}

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

type productResolver struct{ *Resolver }
