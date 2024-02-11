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
	l, r := unsafe.Pointer(&s.nodes[i]), unsafe.Pointer(n)
	if uintptr(l) != uintptr(r) {
		s.nodes[i] = *n
	}
}

func (s *storage) putNodeUnsafe1(i int, n *node) {
	lh := *(*reflect.SliceHeader)(unsafe.Pointer(&s.nodes))
	lhb := reflect.SliceHeader{Data: lh.Data, Len: lh.Len * nodeSZ, Cap: lh.Cap * nodeSZ}
	lb := *(*[]byte)(unsafe.Pointer(&lhb))
	off := i * nodeSZ

	rp := uintptr(unsafe.Pointer(n))
	rh := reflect.SliceHeader{Data: rp, Len: nodeSZ, Cap: nodeSZ}
	rb := *(*[]byte)(unsafe.Pointer(&rh))

	copy(lb[off:off+nodeSZ], rb)
}

func (s *storage) putNodeUnsafe2(i int, n *node) {
	rp := uintptr(unsafe.Pointer(n))
	rh := reflect.SliceHeader{Data: rp, Len: nodeSZ, Cap: nodeSZ}
	rb := *(*[]byte)(unsafe.Pointer(&rh))

	off := i * nodeSZ
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&s.nodes[0])), len(s.nodes)*nodeSZ)[off:off+nodeSZ], rb)
}

func (s *storage) reset() { s.ln = 0 }

var nodeSZ int

func init() {
	var n node
	nodeSZ = int(unsafe.Sizeof(n))
}
