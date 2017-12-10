// +build mock
//go:generate goagen bootstrap -d github.com/uzimith/goa-learn/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/uzimith/goa-learn/app"
	"github.com/uzimith/goa-learn/mock"
)

func main() {
	// Create service
	service := goa.New("goa.learn")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "article" controller
	c := mock.NewArticleController(service)
	app.MountArticleController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
