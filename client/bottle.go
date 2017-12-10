// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "cellar": bottle Resource Client
//
// Command:
// $ goagen
// --design=github.com/uzimith/goa-learn/design
// --out=$(GOPATH)/src/github.com/uzimith/goa-learn
// --version=v1.3.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ShowBottlePath computes a request path to the show action of bottle.
func ShowBottlePath(bottleID int) string {
	param0 := strconv.Itoa(bottleID)

	return fmt.Sprintf("/bottles/%s", param0)
}

// Get bottle by id
func (c *Client) ShowBottle(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowBottleRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowBottleRequest create the request corresponding to the show action endpoint of the bottle resource.
func (c *Client) NewShowBottleRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
