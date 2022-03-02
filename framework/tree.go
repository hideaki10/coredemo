package framework

import (
	"errors"
	"strings"
)

type Tree struct {
	root *node
}

type node struct {
	isLast  bool
	segment string
	handler Controller
	childs  []*node
}

func newNode() *node {
	return &node{
		isLast:  true,
		segment: "",
		childs:  []*node{},
	}

}

func NewTree() *Tree {
	root := newNode()
	return &Tree{
		root: root,
	}
}

func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}

	if isWildSegment(segment) {
		return n.childs
	}

	nodes := make([]*node, 0, len(n.childs))

	for _, cnode := range n.childs {
		if isWildSegment(cnode.segment) {
			nodes = append(nodes, cnode)
		} else if cnode.segment == segment {
			nodes = append(nodes, cnode)
		}
	}

	return nodes
}

func (n *node) matchNode(uri string) *node {

	segments := strings.SplitN(uri, "/", 2)

	segment := segments[0]

	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}

	cnodes := n.filterChildNodes(segment)

	if len(cnodes) == 0 || cnodes == nil {
		return nil
	}

	if len(segments) == 1 {
		for _, tn := range cnodes {
			if tn.isLast {
				return tn
			}
		}
		return nil
	}

	for _, tn := range cnodes {
		tnMatch := tn.matchNode(segments[1])
		if tnMatch != nil {
			return tnMatch
		}
	}

	return nil
}

func (tree *Tree) AddRouter(uri string, handler Controller) error {

	n := tree.root

	if n.matchNode(uri) != nil {
		return errors.New("route exists: " + uri)
	}

	segments := strings.Split(uri, "/")

	for index, segment := range segments {

		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}

		isLast := index == len(segments)-1

		var objNode *node

		childNodes := n.filterChildNodes(segment)

		if len(childNodes) > 0 {
			for _, cnode := range childNodes {
				if cnode.segment == segment {
					objNode = cnode
					break
				}
			}
		}

		if objNode == nil {
			cnode := newNode()
			cnode.segment = segment
			if isLast {
				cnode.isLast = true
				cnode.handler = handler
			}
			n.childs = append(n.childs, cnode)
			objNode = cnode
		}
		n = objNode

	}
	return nil
}

func (tree *Tree) FindHandler(uri string) Controller {
	matchnode := tree.root.matchNode(uri)
	if matchnode == nil {
		return nil
	}
	return matchnode.handler
}
