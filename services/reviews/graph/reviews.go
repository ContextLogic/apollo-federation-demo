package graph

import (
	"github.com/ContextLogic/reviews/graph/model"
	"fmt"
)

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
			ID: fmt.Sprintf("%d", i+1),
			Product: &model.Product{
				Upc: fmt.Sprintf("%d", i+1),
			},
			Author: &model.User{
				ID:       fmt.Sprintf("%d", i/2+1),
				Username: usernames[i/2],
			},
			Body: r,
		})
	}
}
