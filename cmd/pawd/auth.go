package main

import (
	"context"
	"errors"

	pawd "github.com/Xe/pawd/proto"
)

// Auth is the implementation of the grpc service pawd.Auth.
type Auth struct{ *Server }

func (a *Auth) Register(context.Context, *pawd.RegisterInfo) (*pawd.UserToken, error) {
	return nil, errors.New("not implemented")
}

func (a *Auth) Login(context.Context, *pawd.LoginInfo) (*pawd.UserToken, error) {
	return nil, errors.New("not implemented")
}

func (a *Auth) Logout(context.Context, *pawd.UserToken) (*pawd.Nil, error) {
	return nil, errors.New("not implemented")
}
