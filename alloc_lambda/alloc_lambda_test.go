package alloc_lambda

import "testing"

type Row struct {
	X string
}

type Event struct {
	Y []byte
}

func process(row *Row, event *Event) {
	fn := func(row *Row, event *Event) {
		row.X = "foobar"
		_ = event
	}
	fn(row, event)
}

func BenchmarkAllocLambda(b *testing.B) {
	var (
		row   Row
		event Event
	)
	b.ReportAllocs()
	for range b.N {
		process(&row, &event)
	}
}
