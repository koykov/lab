package main

import (
	"testing"

	"github.com/koykov/inspector/testobj"
)

func BenchmarkCopy(b *testing.B) {
	origin := &testobj.TestObject{
		Id:         "foo",
		Name:       []byte("bar"),
		Cost:       12.34,
		Status:     78,
		Permission: &testobj.TestPermission{15: true, 23: false},
		Flags: testobj.TestFlag{
			"export": 17,
			"ro":     4,
			"rw":     7,
			"Valid":  1,
		},
		Finance: &testobj.TestFinance{
			MoneyIn:  3200,
			MoneyOut: 1500.637657,
			Balance:  9000,
			AllowBuy: true,
			History: []testobj.TestHistory{
				{
					152354345634,
					14.345241,
					[]byte("pay for domain"),
				},
				{
					153465345246,
					-3.0000342543,
					[]byte("got refund"),
				},
				{
					156436535640,
					2325242534.35324523,
					[]byte("maintenance"),
				},
			},
		},
	}

	inspectBytes := func(origin *testobj.TestObject) (c int) {
		c += len(origin.Id)
		c += len(origin.Name)
		for k := range origin.Flags {
			c += len(k)
		}
		for j := 0; j < len(origin.Finance.History); j++ {
			c += len(origin.Finance.History[j].Comment)
		}
		return
	}

	b.Run("inspect bytes", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			c := inspectBytes(origin)
			p := make([]byte, 0, c)
			_ = p
		}
	})
}
