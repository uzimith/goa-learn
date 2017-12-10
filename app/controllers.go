// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "goa.learn": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/uzimith/goa-learn/design
// --out=$(GOPATH)/src/github.com/uzimith/goa-learn
// --version=v1.3.0

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// ArticleController is the controller interface for the Article actions.
type ArticleController interface {
	goa.Muxer
	List(*ListArticleContext) error
	Show(*ShowArticleContext) error
}

// MountArticleController "mounts" a Article resource controller on the given service.
func MountArticleController(service *goa.Service, ctrl ArticleController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListArticleContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/article", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Article", "action", "List", "route", "GET /article")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowArticleContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/article/:id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Article", "action", "Show", "route", "GET /article/:id")
}
