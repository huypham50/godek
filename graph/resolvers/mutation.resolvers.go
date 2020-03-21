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
		fmt.Println("4444")
		userCount := r.Services.User.Count()
		username := db.GenerateUsername(email, userCount)

		newUser := &models.User{
			GoogleID: token,
			Name:     name,
			Email:    email,
			Username: username,
			Avatar:   avatar,
		}
		r.Services.User.Create(newUser)

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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
