package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/phamstack/godek/graph/generated"
	"github.com/phamstack/godek/models"
)

func (r *queryResolver) Me(ctx context.Context) (*models.User, error) {
	me := &models.User{
		Name:  "Matt",
		Email: "matt@gooch.bc",
	}

	if err := r.Services.User.Create(me); err != nil {
		return nil, err
	}

	return me, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
