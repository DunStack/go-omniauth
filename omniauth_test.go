package omniauth_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dunstack/go-omniauth"
	gotest "github.com/dunstack/go-test"
)

var oa = omniauth.OmniAuth{
	"foo": FooProvider{},
}

func TestAuthCodeURL(t *testing.T) {
	tests := []struct {
		provider string
		url      string
		err      error
	}{
		{"foo", "https://foo.bar/", nil},
		{"bar", "", omniauth.ErrProviderNotFound("bar")},
	}

	for _, test := range tests {
		t.Run(test.provider, func(t *testing.T) {
			gotest.WithT(t, func(gt *gotest.GoTest) {
				url, err := oa.AuthCodeURL(test.provider, "")
				gt.Expect(url).ToBe(test.url)
				gt.Expect(err).ToBe(err)
			})
		})
	}
}

func TestExchangeAndClaims(t *testing.T) {
	tests := []struct {
		provider string
		code     string
		claims   any
		err      error
	}{
		{"bar", "abc", nil, omniauth.ErrProviderNotFound("bar")},
		{"foo", "", nil, errors.New("invalid code")},
		{"foo", "abc", "xyz", nil},
	}

	for _, test := range tests {
		t.Run(test.provider+"-"+test.code, func(t *testing.T) {
			gotest.WithT(t, func(gt *gotest.GoTest) {
				claims, err := oa.ExchangeAndClaims(test.provider, context.Background(), test.code)
				gt.Expect(claims).ToEqual(test.claims)
				gt.Expect(err).ToEqual(test.err)
			})
		})
	}
}

func TestErrProviderNotFound(t *testing.T) {
	gotest.WithT(t, func(gt *gotest.GoTest) {
		err := omniauth.ErrProviderNotFound("bar")
		gt.Expect(err.Error()).ToBe(`omniauth: provider "bar" not found`)
	})
}
