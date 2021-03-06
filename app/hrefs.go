// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "goa.learn": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/uzimith/goa-learn/design
// --out=$(GOPATH)/src/github.com/uzimith/goa-learn
// --version=v1.3.0

package app

import (
	"fmt"
	"strings"
)

// ArticleHref returns the resource href.
func ArticleHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/article/%v", paramid)
}
