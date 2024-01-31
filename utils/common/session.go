package common

import (
	"errors"
	"simple-bank/config"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type Session interface {
	StoreSession(email string, ctx *gin.Context) error
	ReadSession(ctx *gin.Context) (string, error)
	DeleteSession(ctx *gin.Context) error
}

type session struct {
	cfg config.SessionConfig
}

func (s *session) StoreSession(email string, ctx *gin.Context) error {
	var store = sessions.NewCookieStore([]byte(s.cfg.SessionKey))

	session, err := store.Get(ctx.Request, "isAuthenticated")
	if err != nil {
		return err
	}

	session.Values["authenticated"] = true
	session.Values["email"] = email

	session.Options.MaxAge = s.cfg.MaxAge

	err = session.Save(ctx.Request, ctx.Writer)
	if err != nil {
		return err
	}

	return nil
}

func (s *session) ReadSession(ctx *gin.Context) (string, error) {
	var store = sessions.NewCookieStore([]byte(s.cfg.SessionKey))

	session, err := store.Get(ctx.Request, "isAuthenticated")
	if err != nil {
		return "", err
	}

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return "", errors.New("unauthenticated, please login first")
	}

	email := session.Values["email"]

	return email.(string), nil
}

func (s *session) DeleteSession(ctx *gin.Context) error {
	var store = sessions.NewCookieStore([]byte(s.cfg.SessionKey))

	session, err := store.Get(ctx.Request, "isAuthenticated")
	if err != nil {
		return err
	}

	session.Values["authenticated"] = false
	session.Values["email"] = nil

	err = session.Save(ctx.Request, ctx.Writer)
	if err != nil {
		return err
	}

	return nil
}

func NewSession(cfg config.SessionConfig) Session {
	return &session{cfg: cfg}
}
