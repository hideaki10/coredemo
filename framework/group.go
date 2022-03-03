package framework

type IGroup interface {
	Get(string, Controller)
	Post(string, Controller)
	Put(string, Controller)
	Delete(string, Controller)
}

type Group struct {
	core        *Core
	parent      *Group
	prefix      string
	middlewares []Controller
}

func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:        core,
		parent:      nil,
		prefix:      prefix,
		middlewares: []Controller{},
	}
}

func (g *Group) Get(uri string, handlers ...Controller) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Get(uri, allHandlers...)
}
func (g *Group) Post(uri string, handlers ...Controller) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Post(uri, allHandlers...)
}

func (g *Group) Put(uri string, handlers ...Controller) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Put(uri, allHandlers...)
}

func (g *Group) Delete(uri string, handlers ...Controller) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Delete(uri, allHandlers...)
}

func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}

	return g.parent.getAbsolutePrefix() + g.prefix
}

func (g *Group) getMiddlewares() []Controller {
	if g.parent == nil {
		return g.middlewares
	}

	return append(g.parent.middlewares, g.middlewares...)
}

func (g *Group) Use(middlewares ...Controller) {
	g.middlewares = append(g.middlewares, middlewares...)
}

func (g *Group) Group(uri string) *Group {
	cgroup := NewGroup(g.core, uri)
	cgroup.parent = g
	return cgroup
}
