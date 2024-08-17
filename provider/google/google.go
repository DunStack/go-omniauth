package google

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/dunstack/go-omniauth"
	"golang.org/x/oauth2"
)

var (
	ErrIDTokenNotFound = errors.New("omniauth.google: id_token not found")
	ErrIDTokenInvalid  = errors.New("omniauth.google: id_token invalid")
)

type Claims struct {
	UID           string `json:"sub"`
	Name          string `json:""`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:""`
	Email         string `json:""`
	EmailVerified bool   `json:"email_verified"`
}

var _ omniauth.Provider = new(Provider)

func NewProvider(config *oauth2.Config) *Provider {
	return &Provider{config}
}

type Provider struct {
	*oauth2.Config
}

func (p *Provider) Claims(token *oauth2.Token) (any, error) {
	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, ErrIDTokenNotFound
	}
	payload := strings.Split(idToken, ".")[1]
	data, err := base64.RawStdEncoding.DecodeString(payload)
	if err != nil {
		return nil, ErrIDTokenInvalid
	}
	claims := Claims{}
	err = json.Unmarshal(data, &claims)
	return claims, err
}
