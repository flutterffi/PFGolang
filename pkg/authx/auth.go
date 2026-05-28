package authx

import (
	"errors"
	"net/http"
	"strings"
)

var ErrMissingBearerToken = errors.New("missing bearer token")
var ErrInvalidToken = errors.New("invalid token")

func ExtractBearerToken(r *http.Request) (string, error) {
	header := strings.TrimSpace(r.Header.Get("Authorization"))
	if header == "" {
		return "", ErrMissingBearerToken
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(header, prefix) {
		return "", ErrInvalidToken
	}

	token := strings.TrimSpace(strings.TrimPrefix(header, prefix))
	if token == "" {
		return "", ErrInvalidToken
	}

	return token, nil
}

func ValidateStaticToken(r *http.Request, expected string) error {
	token, err := ExtractBearerToken(r)
	if err != nil {
		return err
	}

	if token != expected {
		return ErrInvalidToken
	}

	return nil
}
