package omniauth_test

import (
	"testing"

	"github.com/dunstack/go-omniauth"
)

func TestOmniAuth(t *testing.T) {
	oa := omniauth.OmniAuth{
		"": omniauth.BaseProvider{},
	}
	if oa == nil {
		t.Error("error")
	}
}
