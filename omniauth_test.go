package omniauth_test

import (
	"testing"

	"github.com/dunstack/go-omniauth"
	"golang.org/x/oauth2"
)

type FooProvider struct {
	*oauth2.Config
}

func (p FooProvider) Claims(token *oauth2.Token) (any, error) {
	return "foo", nil
}

var oa = omniauth.OmniAuth{
	"foo": FooProvider{},
}

func TestProvider(t *testing.T) {
	foo, _ := oa.Provider("foo")
	if _, ok := foo.(FooProvider); !ok {
		t.Errorf("provider: %T, want: %T", foo, FooProvider{})
	}
	_, err := oa.Provider("bar")
	if _, ok := err.(omniauth.ErrProviderNotFound); !ok {
		t.Errorf("err: %v, want: %v", err, omniauth.ErrProviderNotFound("bar"))
	}
}
