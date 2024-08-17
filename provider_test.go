package omniauth_test

import (
	"context"
	"errors"

	"github.com/dunstack/go-omniauth"
	"golang.org/x/oauth2"
)

var _ omniauth.Provider = new(FooProvider)

type FooProvider struct {
}

func (FooProvider) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return "https://foo.bar/"
}

func (FooProvider) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	if code == "" {
		return nil, errors.New("invalid code")
	}
	return &oauth2.Token{}, nil
}

func (FooProvider) Claims(token *oauth2.Token) (any, error) {
	return "xyz", nil
}
