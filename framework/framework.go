package framework

import "net/http"

type Core struct {
	router map[string]Controller
}

func NewCore() *Core {
	return &Core{}
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {

}

func (c *Core) Get(url string, handler Controller) {
	c.router[url] = handler

}
