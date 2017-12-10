package mock

import (
	"time"

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
	res := make(app.ArticleCollection, 0, len(ctx.ID))
	for idx, id := range ctx.ID {
		res = append(res, &app.Article{
			ID:        id,
			Text:      "aiueo",
			CreatedBy: &idx,
			CreatedAt: time.Now(),
		})
	}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *ArticleController) Show(ctx *app.ShowArticleContext) error {
	createdBy := 1
	res := app.Article{
		ID:        ctx.ID,
		Text:      "aiueo",
		CreatedBy: &createdBy,
		CreatedAt: time.Now(),
	}
	return ctx.OK(&res)
}
