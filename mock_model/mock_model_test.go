package mock_model

import "testing"

func BenchmarkMockModel(b *testing.B) {
	b.Run("model", func(b *testing.B) {
		b.ReportAllocs()
		ctx := Ctx{buf: make([]string, 512)}
		for i := 0; i < b.N; i++ {
			m := Model{ctx: ctx}
			_ = m.GetID(1)
		}
	})
	b.Run("mock", func(b *testing.B) {
		b.ReportAllocs()
		ctx := Ctx{buf: make([]string, 512)}
		for i := 0; i < b.N; i++ {
			m := MockModel{ctx: ctx}
			_ = m.GetID(1)
		}
	})
}
