package graph

import "github.com/ContextLogic/inventory/graph/model"

var products = []*model.Product{
	{
		Upc:     "1",
		InStock: true,
		ShippingEstimate: 100,
	},
	{
		Upc:     "2",
		InStock: false,
		ShippingEstimate: 200,
	},
	{
		Upc:     "3",
		InStock: true,
		ShippingEstimate: 300,
	},
}
