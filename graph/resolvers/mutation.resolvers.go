package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/phamstack/godek/graph/generated"
	"github.com/phamstack/godek/lib/auth"
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
		newUser := &models.User{
			GoogleID: token,
			Name:     name,
			Email:    email,
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

func (r *mutationResolver) UpdateGoogleAccount(ctx context.Context, name string) (*models.User, error) {
	// get user from context middleware
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, errors.New("You are not logged in yet")
	}

	if err := r.Services.User.Update(user, name); err != nil {
		return nil, err
	}

	user.Name = name

	return user, nil
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

	if err := r.Services.Deck.Create(newDeck); err != nil {
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

func (r *mutationResolver) CreateTodo(ctx context.Context, deckID *int, title string, description string, deadline time.Time) (*models.Todo, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, errors.New("You are not logged in yet")
	}

	newTodo := &models.Todo{
		UserID:      user.ID,
		DeckID:      *deckID,
		Title:       title,
		Description: description,
		Deadline:    deadline,
	}

	if err := r.Services.Todo.Create(newTodo); err != nil {
		return nil, err
	}

	return newTodo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, id int, deckID *int, title string, description string, deadline time.Time, complete bool) (*models.Todo, error) {
	// get user from context middleware
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, errors.New("You are not logged in yet")
	}

	// Find requested removed deck and verify
	updatedTodo, err := r.Services.Todo.ByID(id)
	if err != nil {
		return nil, err
	}

	// deck can only be removed by owner
	if user.ID != updatedTodo.UserID {
		return nil, errors.New("You are unauthorized to edit this deck")
	}

	updatedTodo.DeckID = *deckID
	updatedTodo.Title = title
	updatedTodo.Description = description
	updatedTodo.Deadline = deadline

	if err := r.Services.Todo.Update(updatedTodo); err != nil {
		return nil, err
	}

	return updatedTodo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id int) (*models.Todo, error) {
	// get user from context middleware
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, errors.New("You are not logged in yet")
	}

	// Find requested removed deck and verify
	removedTodo, err := r.Services.Todo.ByID(id)
	if err != nil {
		return nil, err
	}

	// deck can only be removed by owner
	if user.ID != removedTodo.UserID {
		return nil, errors.New("You are unauthorized to delete this deck")
	}

	// remove user
	if err := r.Services.Todo.Delete(removedTodo); err != nil {
		return nil, err
	}

	return removedTodo, nil
}

func (r *mutationResolver) CreateBookmark(ctx context.Context, url string) (*models.Bookmark, error) {
	fmt.Println("000000000000")

	user := auth.ForContext(ctx)
	if user == nil {
		return nil, errors.New("You are not logged in yet")
	}

	fmt.Println("111111111111")

	fetchedBookmark, err := r.Services.Bookmark.FetchURL(url)
	if err != nil {
		return nil, err
	}

	fmt.Println("22222222222")

	fetchedBookmark.UserID = user.ID

	fmt.Printf("%+v\n", fetchedBookmark)

	if err := r.Services.Bookmark.Create(fetchedBookmark); err != nil {
		return nil, err
	}

	fmt.Println("33333333333")

	return fetchedBookmark, nil
}

func (r *mutationResolver) UpdateBookmark(ctx context.Context, id int, deckID *int, title string, description string) (*models.Bookmark, error) {
	// get user from context middleware
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, errors.New("You are not logged in yet")
	}

	// Find requested removed deck and verify
	updatedBookmark, err := r.Services.Bookmark.ByID(id)
	if err != nil {
		return nil, err
	}

	// deck can only be removed by owner
	if user.ID != updatedBookmark.UserID {
		return nil, errors.New("You are unauthorized to edit this deck")
	}

	updatedBookmark.DeckID = *deckID
	updatedBookmark.Title = title
	updatedBookmark.Description = description

	if err := r.Services.Bookmark.Update(updatedBookmark); err != nil {
		return nil, err
	}

	return updatedBookmark, nil
}

func (r *mutationResolver) DeleteBookmark(ctx context.Context, id int) (*models.Bookmark, error) {
	// get user from context middleware
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, errors.New("You are not logged in yet")
	}

	// Find requested removed deck and verify
	removedBookmark, err := r.Services.Bookmark.ByID(id)
	if err != nil {
		return nil, err
	}

	// deck can only be removed by owner
	if user.ID != removedBookmark.UserID {
		return nil, errors.New("You are unauthorized to delete this deck")
	}

	// remove user
	if err := r.Services.Bookmark.Delete(removedBookmark); err != nil {
		return nil, err
	}

	return removedBookmark, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) FetchBookmark(ctx context.Context, url string) (*models.Bookmark, error) {
	fetchedBookmark, err := r.Services.Bookmark.FetchURL(url)
	if err != nil {
		return nil, err
	}

	return fetchedBookmark, nil
}
