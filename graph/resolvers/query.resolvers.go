package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/phamstack/godek/graph/generated"
	"github.com/phamstack/godek/lib/auth"
	"github.com/phamstack/godek/models"
)

func (r *queryResolver) Me(ctx context.Context) (*models.User, error) {
	// TODO: Preload stuffs -> rethink fetching at middleware
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, errors.New("Unauthenticated")
	}

	return user, nil
}

func (r *queryResolver) Decks(ctx context.Context) ([]*models.Deck, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, errors.New("Unauthenticated")
	}

	decks, err := r.Services.Deck.Filter(user.ID)
	if err != nil {
		return nil, err
	}

	return decks, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, errors.New("Unauthenticated")
	}

	todos, err := r.Services.Todo.Filter(user.ID)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *queryResolver) Bookmarks(ctx context.Context) ([]*models.Bookmark, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, errors.New("Unauthenticated")
	}

	bookmarks, err := r.Services.Bookmark.Filter(user.ID)
	if err != nil {
		return nil, err
	}

	return bookmarks, nil
}

func (r *queryResolver) Snippets(ctx context.Context) ([]*models.Snippet, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, errors.New("Unauthenticated")
	}

	snippets, err := r.Services.Snippet.Filter(user.ID)
	if err != nil {
		return nil, err
	}

	return snippets, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
