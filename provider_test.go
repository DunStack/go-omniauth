package omniauth_test

import (
	"reflect"
	"testing"

	"github.com/dunstack/go-omniauth"
	"golang.org/x/oauth2"
)

func TestBaseProviderClaims(t *testing.T) {
	p := new(omniauth.BaseProvider)

	t.Run("id_token not found", func(t *testing.T) {
		token := new(oauth2.Token)
		_, err := p.Claims(token)
		if err != omniauth.ErrIDTokenNotFound {
			t.Errorf("error = %s; want %s", err, omniauth.ErrIDTokenNotFound)
		}
	})

	t.Run("id_token invalid", func(t *testing.T) {
		idToken := "a.b.c"
		token := new(oauth2.Token).WithExtra(map[string]any{"id_token": idToken})
		_, err := p.Claims(token)
		if err == nil {
			t.Error("error should not be nil")
		}
	})

	t.Run("id_token valid", func(t *testing.T) {
		idToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		token := new(oauth2.Token).WithExtra(map[string]any{"id_token": idToken})
		claims, _ := p.Claims(token)
		want := omniauth.Claims{
			"sub":  "1234567890",
			"name": "John Doe",
			"iat":  float64(1516239022),
		}
		if !reflect.DeepEqual(claims, want) {
			t.Errorf("claims = %v, want %v", claims, want)
		}
	})
}
