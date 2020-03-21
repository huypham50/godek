package auth

import (
	"context"
	"net/http"

	"github.com/phamstack/godek/models"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

// ContextKey -> context id
type contextKey struct {
	name string
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(s *models.Services) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get string will convert to object if necessary
			// undefined, "", nil
			token := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if token == "" {
				next.ServeHTTP(w, r)
				return
			}

			googleID, err := s.User.ParseAuthToken(token)
			if err != nil {
				http.Error(w, "Invalid token 403 status codes are important", http.StatusForbidden)
				return
			}

			// get the user from the database
			user, _ := s.User.ByGoogleID(googleID)

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *models.User {
	user, _ := ctx.Value(userCtxKey).(*models.User)
	return user
}
