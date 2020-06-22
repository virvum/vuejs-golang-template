package main

import (
	"errors"
	"net/http"

	// TODO
	_ "github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
)

// Session holds a user session.
type Session struct {
	UserID   int64
	Username string
	Token    string
}

var (
	sessions = make(map[string]Session)

	// ErrAPIAuthNoCookie holds the error when the session cookie was not found.
	ErrAPIAuthNoCookie = errors.New("session cookie not found")

	// ErrAPIAuthTokenMap holds the error when no token was found for a session.
	ErrAPIAuthTokenMap = errors.New("invalid token: unable to map token to session")
)

// APIAuthenticate checks whether a user is logged in and returns the Session object.
func APIAuthenticate(ctx *gin.Context) (*Session, error) {
	// TODO

	return nil, nil
}

// APIGetSession returns the Session struct stored in the Gin context by the
// authentication middleware.
func APIGetSession(ctx *gin.Context) Session {
	s, ok := ctx.Get("session")
	if !ok {
		// Key "session" should always exist, otherwise there's a logical code error.
		panic("Gin-Gonic session key doesn't exist")
	}

	return *s.(*Session)
}

// APIAuthMiddleware is the authentication middleware, which ensures that a
// user is logged in and stores the Session struct in Gin's context.
func APIAuthMiddleware(ctx *gin.Context) {
	session, err := APIAuthenticate(ctx)

	if err == ErrAPIAuthNoCookie || err == ErrAPIAuthTokenMap {
		APIError(ctx, http.StatusUnauthorized, err.Error())
		return
	} else if err != nil {
		APIError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Set("session", session)
	ctx.Next()
}

// APIAuthGet checks whether a user is logged in.
func APIAuthGet(ctx *gin.Context) {
	// The middleware APIAuthMiddleware() takes care of authentication, so
	// we can just return a successful status here.
	ctx.JSON(http.StatusOK, gin.H{})
}

// APIAuthPost authenticates a user and creates a new session.
func APIAuthPost(ctx *gin.Context) {
	// TODO
	ctx.JSON(http.StatusOK, gin.H{})
}

// APIAuthDelete delete an existing session (log user out).
func APIAuthDelete(ctx *gin.Context) {
	// TODO
	// session := APIGetSession(ctx)

	ctx.JSON(http.StatusOK, gin.H{})
}
