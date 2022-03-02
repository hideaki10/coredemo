package framework

type IGroup interface {
	Get(string, Controller)
	Post(string, Controller)
	Put(string, Controller)
	Delete(string, Controller)
}

type Group struct {
	core   *Core
	parent *Group
	prefix string
}

func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		parent: nil,
		prefix: prefix,
	}
}

func (g *Group) Get(uri string, handler Controller) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Get(uri, handler)
}

func (g *Group) Post(uri string, handler Controller) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Post(uri, handler)
}

func (g *Group) Put(uri string, handler Controller) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Put(uri, handler)
}

func (g *Group) Delete(uri string, handler Controller) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Delete(uri, handler)
}

func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}

	return g.parent.getAbsolutePrefix() + g.prefix
}

func (g *Group) Group(uri string) *Group {
	cgroup := NewGroup(g.core, uri)
	cgroup.parent = g
	return cgroup
}
