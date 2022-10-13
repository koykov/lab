package pkg

type myiface interface {
	My() int
}

type X struct {
	New func() any
}
