package omniauth

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

type OmniAuth map[string]Provider

func (oa OmniAuth) AuthCodeURL(name string, state string, opts ...oauth2.AuthCodeOption) (string, error) {
	p, err := oa.Provider(name)
	if err != nil {
		return "", err
	}
	return p.AuthCodeURL(state, opts...), nil
}

func (oa OmniAuth) ExchangeAndClaims(name string, ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (any, error) {
	p, err := oa.Provider(name)
	if err != nil {
		return nil, err
	}
	token, err := p.Exchange(ctx, code, opts...)
	if err != nil {
		return nil, err
	}
	return p.Claims(token)
}

func (oa OmniAuth) Provider(name string) (Provider, error) {
	provider, ok := oa[name]
	if !ok {
		return nil, ErrProviderNotFound(name)
	}
	return provider, nil
}

type ErrProviderNotFound string

func (p ErrProviderNotFound) Error() string {
	return fmt.Sprintf("omniauth: provider %q not found", string(p))
}
