package duffcopy

import (
	"reflect"
	"unsafe"
)

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

func (s *storage) getNodeBCE() (i int, n *node) {
	n_ := len(s.nodes)
	if s.ln < n_ {
		_ = s.nodes[n_-1]
		n = &s.nodes[s.ln]
		n.typ = 0
	} else {
		s.nodes = append(s.nodes, node{})
		n_++
		_ = s.nodes[n_-1]
		n = &s.nodes[n_-1]
	}
	s.ln++
	i = s.ln - 1
	return
}

func (s *storage) putNode(i int, n *node) {
	s.nodes[i] = *n
}

func (s *storage) putNodeUnsafe(i int, n *node) {
	l := unsafe.Pointer(&s.nodes[i])
	r := unsafe.Pointer(n)
	l = r
	_ = l
}

func (s *storage) reset() { s.ln = 0 }
