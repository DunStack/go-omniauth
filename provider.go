package omniauth

import (
	"context"

	"golang.org/x/oauth2"
)

type Provider interface {
	AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
	Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
	Claims(token *oauth2.Token) (any, error)
}
