package providers

import (
	"context"
	"errors"
	"net/url"
)

type Repo interface {
	Create(ctx context.Context, provider *Provider) (*Provider, error)
	GetAll(ctx context.Context, limit, offset uint) (providers []Provider, total uint, err error)
}

type Provider struct {
	ID   uint
	Name string
	URL  string
}

func ValidateURL(input string) error {
	parsedURL, err := url.ParseRequestURI(input)
	if err != nil {
		return errors.New("invalid URL")
	}

	// Check if the scheme (protocol) and host are set in the parsed URL
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return errors.New("invalid URL")
	}
	return nil
}
