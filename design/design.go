package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("goa.learn", func() {
	Title("goa learn")
	Description("learning goa")
	Scheme("http")
	Host("localhost:8080")
})

var Article = MediaType("application/vnd.article+json", func() {
	Description("Article")
	Attributes(func() {
		Attribute("id", Integer, "ID", func() {
			Example(1)
		})
		Attribute("text", String, "text", func() {
		})
		Attribute("created_by", Integer, "作成者", func() {
			Example(1)
		})
		Attribute("created_at", DateTime, "作成日時")
		Required("id", "text", "created_at")
	})
	View("default", func() {
		Attribute("id")
		Attribute("text")
		Attribute("created_by")
		Attribute("created_at")
	})
})

var _ = Resource("article", func() {
	BasePath("/article")
	DefaultMedia(Article)
	Action("show", func() {
		Description("Get arcitle by id")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer, "id", func() {
				Minimum(1)
			})
		})
		Response(OK, func() {
			Media(Article, "default")
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("List")
		Params(func() {
			Param("id", ArrayOf(Integer), "Filter by ids")
		})
		Response(OK, func() {
			Media(CollectionOf(Article, func() {
				View("default")
			}))
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})
