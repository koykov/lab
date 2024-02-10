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
	h := *(*reflect.SliceHeader)(unsafe.Pointer(&s.nodes))
	hr := reflect.SliceHeader{Data: h.Data, Len: h.Len * nodeSZ, Cap: h.Cap * nodeSZ}
	raw := *(*[]byte)(unsafe.Pointer(&hr))
	_ = raw
	off := i * nodeSZ

	nr := uintptr(unsafe.Pointer(n))
	nh := reflect.SliceHeader{Data: nr, Len: nodeSZ, Cap: nodeSZ}
	nb := *(*[]byte)(unsafe.Pointer(&nh))

	copy(raw[off:off+nodeSZ], nb)
}

func (s *storage) reset() { s.ln = 0 }

var nodeSZ int

func init() {
	var n node
	nodeSZ = int(unsafe.Sizeof(n))
}
