package main

import (
	"github.com/goadesign/goa"
	"github.com/uzimith/goa-learn/app"
)

// ArticleController implements the article resource.
type ArticleController struct {
	*goa.Controller
}

// NewArticleController creates a article controller.
func NewArticleController(service *goa.Service) *ArticleController {
	return &ArticleController{Controller: service.NewController("ArticleController")}
}

// List runs the list action.
func (c *ArticleController) List(ctx *app.ListArticleContext) error {
	// ArticleController_List: start_implement

	// Put your logic here

	// ArticleController_List: end_implement
	res := app.AjaGreenArticleCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *ArticleController) Show(ctx *app.ShowArticleContext) error {
	// ArticleController_Show: start_implement

	// Put your logic here

	// ArticleController_Show: end_implement
	res := &app.AjaGreenArticle{}
	return ctx.OK(res)
}
