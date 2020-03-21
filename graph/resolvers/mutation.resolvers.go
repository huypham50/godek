package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/phamstack/godek/graph/generated"
	"github.com/phamstack/godek/lib/auth"
	"github.com/phamstack/godek/lib/db"
	"github.com/phamstack/godek/models"
)

func (r *mutationResolver) LoginGoogle(ctx context.Context, token string, name string, email string, avatar string) (*models.Auth, error) {
	fmt.Println("#FETCHING", token, email)
	user, err := r.Services.User.ByGoogleID(token)

	fmt.Println("User is:")
	fmt.Printf("%+v\n", user)

	if err != nil && err != models.ErrNotFound {
		return nil, err
	}

	if user != nil && user.Email != email {
		return nil, errors.New("ggid != email")
	}

	if err == models.ErrNotFound {
		userCount := r.Services.User.Count()
		username := db.GenerateUsername(email, userCount)

		newUser := &models.User{
			GoogleID: token,
			Name:     name,
			Email:    email,
			Username: username,
			Avatar:   avatar,
		}
		err := r.Services.User.Create(newUser)
		if err != nil {
			return nil, err
		}

		authToken := r.Services.User.GenerateAuthToken(newUser)

		return &models.Auth{
			User:  newUser,
			Token: authToken,
		}, nil
	}

	authToken := r.Services.User.GenerateAuthToken(user)
	return &models.Auth{
		User:  user,
		Token: authToken,
	}, nil
}

func (r *mutationResolver) Logout(ctx context.Context) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) LogoutAll(ctx context.Context) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteGoogleAccount(ctx context.Context, email string) (*models.User, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, errors.New("Unauthenticated")
	}

	if user.Email == email {
		err := r.Services.User.Delete(user)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (r *mutationResolver) UpdateGoogleAccount(ctx context.Context, name string, username string) (*models.User, error) {
	// get user from context middleware
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, errors.New("You are not logged in yet")
	}

	// Find requested removed deck and verify
	updatedUser, err := r.Services.User.ID(user.ID)
	if err != nil {
		return nil, err
	}

	// deck can only be removed by owner
	if user.ID != updatedUser.ID {
		return nil, errors.New("You are unauthorized to edit this deck")
	}

	updatedUser.Name = name

	if err := r.Services.User.Update(updatedUser); err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (r *mutationResolver) CreateDeck(ctx context.Context, title string, description string, label string, color string) (*models.Deck, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, errors.New("You are not logged in yet")
	}

	newDeck := &models.Deck{
		UserID:      user.ID,
		Title:       title,
		Description: description,
		Label:       label,
		Color:       color,
	}

	err := r.Services.Deck.Create(newDeck)
	if err != nil {
		return nil, err
	}

	return newDeck, nil
}

func (r *mutationResolver) UpdateDeck(ctx context.Context, id int, title string, description string, label string, color string, archive bool) (*models.Deck, error) {
	// get user from context middleware
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, errors.New("You are not logged in yet")
	}

	// Find requested removed deck and verify
	updatedDeck, err := r.Services.Deck.ByID(id)
	if err != nil {
		return nil, err
	}

	// deck can only be removed by owner
	if user.ID != updatedDeck.UserID {
		return nil, errors.New("You are unauthorized to edit this deck")
	}

	updatedDeck.Title = title
	updatedDeck.Description = description
	updatedDeck.Label = label
	updatedDeck.Color = color
	updatedDeck.Archive = archive

	if err := r.Services.Deck.Update(updatedDeck); err != nil {
		return nil, err
	}

	return updatedDeck, nil
}

func (r *mutationResolver) DeleteDeck(ctx context.Context, id int) (*models.Deck, error) {
	// get user from context middleware
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, errors.New("You are not logged in yet")
	}

	// Find requested removed deck and verify
	removedDeck, err := r.Services.Deck.ByID(id)
	if err != nil {
		return nil, err
	}

	// deck can only be removed by owner
	if user.ID != removedDeck.UserID {
		return nil, errors.New("You are unauthorized to delete this deck")
	}

	// remove user
	if err := r.Services.Deck.Delete(removedDeck); err != nil {
		return nil, err
	}

	return removedDeck, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
