// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "goa.learn": Application Contexts
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
	"strconv"
)

// ListArticleContext provides the article list action context.
type ListArticleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID []int
}

// NewListArticleContext parses the incoming request URL and body, performs validations and creates the
// context used by the article controller list action.
func NewListArticleContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListArticleContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListArticleContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		params := make([]int, len(paramID))
		for i, rawID := range paramID {
			if id, err2 := strconv.Atoi(rawID); err2 == nil {
				params[i] = id
			} else {
				err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
			}
		}
		rctx.ID = params
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListArticleContext) OK(r ArticleCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.article+json; type=collection")
	if r == nil {
		r = ArticleCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ListArticleContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListArticleContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowArticleContext provides the article show action context.
type ShowArticleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewShowArticleContext parses the incoming request URL and body, performs validations and creates the
// context used by the article controller show action.
func NewShowArticleContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowArticleContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowArticleContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
		if rctx.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`id`, rctx.ID, 1, true))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowArticleContext) OK(r *Article) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.article+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ShowArticleContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowArticleContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
