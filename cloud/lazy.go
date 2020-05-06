package cloud

import (
	"io"
	"sync"
)

var _ GraphAPI = new(LazyGraph)

type LazyGraph struct {
	LoadingFunc func() GraphAPI
	once        sync.Once
	api         GraphAPI
}

func (g *LazyGraph) load() {
	g.once.Do(func() {
		g.api = g.LoadingFunc()
	})
}

func (g *LazyGraph) Find(q Query) ([]Resource, error) {
	g.load()
	return g.api.Find(q)
}

func (g *LazyGraph) FindWithProperties(props map[string]interface{}) ([]Resource, error) {
	g.load()
	return g.api.FindWithProperties(props)
}

func (g *LazyGraph) FindOne(q Query) (Resource, error) {
	g.load()
	return g.api.FindOne(q)
}

func (g *LazyGraph) FilterGraph(q Query) (GraphAPI, error) {
	g.load()
	return g.api.FilterGraph(q)
}

func (g *LazyGraph) MarshalTo(w io.Writer) error {
	g.load()
	return g.api.MarshalTo(w)
}

func (g *LazyGraph) ResourceRelations(r Resource, relation string, recursive bool) ([]Resource, error) {
	g.load()
	return g.api.ResourceRelations(r, relation, recursive)
}

func (g *LazyGraph) VisitRelations(r Resource, relation string, includeResource bool, each func(Resource, int) error) error {
	g.load()
	return g.api.VisitRelations(r, relation, includeResource, each)
}

func (g *LazyGraph) ResourceSiblings(r Resource) ([]Resource, error) {
	g.load()
	return g.api.ResourceSiblings(r)
}

func (g *LazyGraph) Merge(aG GraphAPI) error {
	g.load()
	return g.api.Merge(aG)
}
