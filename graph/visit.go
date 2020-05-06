package graph

import (
	"github.com/hbbio/awless/cloud/rdf"
	tstore "github.com/wallix/triplestore"
)

type Visitor interface {
	Visit(*Graph) error
}

type visitEachFunc func(res *Resource, depth int) error

func VisitorCollectFunc(collect *[]*Resource) visitEachFunc {
	return func(res *Resource, depth int) error {
		*collect = append(*collect, res)
		return nil
	}
}

type ParentsVisitor struct {
	From        *Resource
	Each        visitEachFunc
	IncludeFrom bool
	Relation    string
}

func (v *ParentsVisitor) Visit(g *Graph) error {
	startNode, foreach, err := prepareRDFVisit(g, v.From, v.Each, v.IncludeFrom)
	if err != nil {
		return err
	}
	if v.Relation == "" {
		v.Relation = rdf.ParentOf
	}
	return tstore.NewTree(g.store.Snapshot(), v.Relation).TraverseAncestors(startNode, foreach)
}

type ChildrenVisitor struct {
	From        *Resource
	Each        visitEachFunc
	IncludeFrom bool
	Relation    string
}

func (v *ChildrenVisitor) Visit(g *Graph) error {
	startNode, foreach, err := prepareRDFVisit(g, v.From, v.Each, v.IncludeFrom)
	if err != nil {
		return err
	}
	if v.Relation == "" {
		v.Relation = rdf.ParentOf
	}
	return tstore.NewTree(g.store.Snapshot(), v.Relation).TraverseDFS(startNode, foreach)
}

type SiblingsVisitor struct {
	From        *Resource
	Each        visitEachFunc
	IncludeFrom bool
}

func (v *SiblingsVisitor) Visit(g *Graph) error {
	startNode, foreach, err := prepareRDFVisit(g, v.From, v.Each, v.IncludeFrom)
	if err != nil {
		return err
	}

	return tstore.NewTree(g.store.Snapshot(), rdf.ParentOf).TraverseSiblings(startNode, resolveResourceType, foreach)
}

func prepareRDFVisit(g *Graph, root *Resource, each visitEachFunc, includeRoot bool) (string, func(g tstore.RDFGraph, n string, i int) error, error) {
	rootNode := root.Id()

	foreach := func(rdfG tstore.RDFGraph, n string, i int) error {
		rT, err := resolveResourceType(rdfG, n)
		if err != nil {
			return err
		}
		res, err := g.GetResource(rT, n)
		if err != nil {
			return err
		}
		if includeRoot || !root.Same(res) {
			if err := each(res, i); err != nil {
				return err
			}
		}
		return nil
	}
	return rootNode, foreach, nil
}
