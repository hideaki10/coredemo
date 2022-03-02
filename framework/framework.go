package framework

import (
	"net/http"
	"strings"
)

type Core struct {
	router map[string]map[string]Controller
}

func NewCore() *Core {

	getRouter := map[string]Controller{}
	postRouter := map[string]Controller{}
	putRouter := map[string]Controller{}
	deleteRouter := map[string]Controller{}

	router := map[string]map[string]Controller{}
	router["GET"] = getRouter
	router["POST"] = postRouter
	router["PUT"] = putRouter
	router["DELETE"] = deleteRouter

	return &Core{router: router}
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	ctx := NewContext(request, response)
	router := c.FindRouteByRequest(request)

	if router != nil {
		ctx.Json(404, "not found")
		return
	}

	if err := router(ctx); err != nil {
		ctx.Json(500, "internal server error")
		return
	}
}

func (c *Core) Get(url string, handler Controller) {
	upperUrl := strings.ToUpper(url)
	c.router["GET"][upperUrl] = handler
}

func (c *Core) Post(url string, handler Controller) {
	upperUrl := strings.ToUpper(url)
	c.router["POST"][upperUrl] = handler
}
func (c *Core) Put(url string, handler Controller) {
	upperUrl := strings.ToUpper(url)
	c.router["PUT"][upperUrl] = handler
}
func (c *Core) Delete(url string, handler Controller) {
	upperUrl := strings.ToUpper(url)
	c.router["DELETE"][upperUrl] = handler
}
func (c *Core) Group(prefix string) *Group {
	return NewGroup(c, prefix)
}

func (c *Core) FindRouteByRequest(request *http.Request) Controller {

	uri := request.URL.Path
	method := request.Method
	upperUri := strings.ToUpper(uri)
	upperMethod := strings.ToUpper(method)

	if methodHandlers, ok := c.router[upperMethod]; ok {
		if handler, ok := methodHandlers[upperUri]; ok {
			return handler
		}
	}

	return nil
}
