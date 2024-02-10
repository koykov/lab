package duffcopy

import "reflect"

type byteptr = reflect.SliceHeader

type node struct {
	typ                       int
	key, val                  byteptr
	idx, depth, offset, limit int
	vptr, pptr                uintptr
}

type storage struct {
	nodes []node
	ln    int
}

func (s *storage) getNode() (i int, n *node) {
	if s.ln < len(s.nodes) {
		n = &s.nodes[s.ln]
		n.typ = 0
	} else {
		n = &node{}
		s.nodes = append(s.nodes, *n)
	}
	s.ln++
	i = s.ln - 1
	return
}

func (s *storage) putNode(i int, n *node) {
	s.nodes[i] = *n
}

func (s *storage) reset() { s.ln = 0 }
