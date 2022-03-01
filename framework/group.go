package framework

type IGroup interface {
	Get(string, Controller)
	Post(string, Controller)
	Put(string, Controller)
	Delete(string, Controller)
}

type Group struct {
	core   *Core
	prefix string
}

func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		prefix: prefix,
	}
}

func (g *Group) Get(uri string, handler Controller) {
	uri = g.prefix + uri
	g.core.Get(uri, handler)
}

func (g *Group) Post(uri string, handler Controller) {
	uri = g.prefix + uri
	g.core.Post(uri, handler)
}

func (g *Group) Put(uri string, handler Controller) {
	uri = g.prefix + uri
	g.core.Put(uri, handler)
}

func (g *Group) Delete(uri string, handler Controller) {
	uri = g.prefix + uri
	g.core.Delete(uri, handler)
}

func (c *Core) Group(prefix string) *Group {
	return NewGroup(c, prefix)
}
