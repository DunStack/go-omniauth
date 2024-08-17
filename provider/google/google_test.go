package google_test

import (
	"testing"

	"github.com/dunstack/go-omniauth/provider/google"
	gotest "github.com/dunstack/go-test"
	"golang.org/x/oauth2"
)

func TestClaims(t *testing.T) {
	p := google.NewProvider(&oauth2.Config{})

	tests := []struct {
		name   string
		token  *oauth2.Token
		claims any
		err    error
	}{
		{
			name:  "id_token not found",
			token: new(oauth2.Token),
			err:   google.ErrIDTokenNotFound,
		},
		{
			name: "id_token invalid",
			token: new(oauth2.Token).WithExtra(map[string]any{
				"id_token": "a.b.c",
			}),
			err: google.ErrIDTokenInvalid,
		},
		{
			name: "id_token valid",
			token: new(oauth2.Token).WithExtra(map[string]any{
				"id_token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiZ2l2ZW5fbmFtZSI6IkpvaG4iLCJmYW1pbHlfbmFtZSI6IkRvZSIsInBpY3R1cmUiOiJodHRwczovL2pvaG5kb2UucGljdHVyZS8iLCJlbWFpbCI6ImpvaG5AZG9lLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlfQ.l_IGqrGQdobxQKZdHWP3p64MxPKoObBbfm-kJpOx8Og",
			}),
			claims: google.Claims{
				UID:           "1234567890",
				Name:          "John Doe",
				GivenName:     "John",
				FamilyName:    "Doe",
				Picture:       "https://johndoe.picture/",
				Email:         "john@doe.com",
				EmailVerified: true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotest.WithT(t, func(gt *gotest.GoTest) {
				claims, err := p.Claims(test.token)
				if test.err != nil {
					gt.Expect(err).ToBe(test.err)
				}
				if test.claims != nil {
					gt.Expect(claims).ToEqual(test.claims)
				}
			})
		})
	}
}
