package omniauth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"golang.org/x/oauth2"
)

type Provider interface {
	Claims(token *oauth2.Token) (Claims, error)
}

var ErrIDTokenNotFound = errors.New("id_token not found")

var _ Provider = new(BaseProvider)

type BaseProvider struct {
}

func (p BaseProvider) Claims(token *oauth2.Token) (Claims, error) {
	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, ErrIDTokenNotFound
	}
	payload := strings.Split(idToken, ".")[1]
	data, err := base64.RawStdEncoding.DecodeString(payload)
	if err != nil {
		return nil, err
	}
	claims := Claims{}
	err = json.Unmarshal(data, &claims)
	return claims, err
}
