package graph

import "github.com/ContextLogic/inventory/graph/model"

var products = []*model.Product{
	{
		Upc:     "1",
		InStock: true,
	},
	{
		Upc:     "2",
		InStock: false,
	},
	{
		Upc:     "3",
		InStock: true,
	},
}
