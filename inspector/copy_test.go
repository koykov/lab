package main

import (
	"testing"

	"github.com/koykov/fastconv"
	"github.com/koykov/inspector/testobj"
)

func inspectBytes(x *testobj.TestObject) (c int) {
	c += len(x.Id)
	c += len(x.Name)
	for k := range x.Flags {
		c += len(k)
	}
	for j := 0; j < len(x.Finance.History); j++ {
		c += len(x.Finance.History[j].Comment)
	}
	return
}
func bufferize(buf, p []byte) ([]byte, []byte) {
	off := len(buf)
	buf = append(buf, p...)
	return buf, buf[off:]
}
func bufferizeS(buf []byte, s string) ([]byte, string) {
	off := len(buf)
	buf = append(buf, s...)
	return buf, fastconv.B2S(buf[off:])
}
func cpy(x *testobj.TestObject) *testobj.TestObject {
	bc := inspectBytes(x)
	buf := make([]byte, 0, bc)

	var c testobj.TestObject
	buf, c.Id = bufferizeS(buf, x.Id)
	buf, c.Name = bufferize(buf, x.Name)
	c.Status = x.Status
	c.Ustate = x.Ustate
	c.Cost = x.Cost
	if x.Permission != nil && len(*x.Permission) > 0 {
		x1 := make(testobj.TestPermission, len(*x.Permission))
		for k, v := range *x.Permission {
			x1[k] = v
		}
		c.Permission = &x1
	}
	if len(x.HistoryTree) > 0 {
		c.HistoryTree = make(map[string]*testobj.TestHistory, len(x.HistoryTree))
		for k, v := range x.HistoryTree {
			x2 := testobj.TestHistory{}
			if v != nil {
				x2.DateUnix = v.DateUnix
				x2.Cost = v.Cost
				buf, x2.Comment = bufferize(buf, x2.Comment)
			}
			var k1 string
			buf, k1 = bufferizeS(buf, k)
			c.HistoryTree[k1] = &x2
		}
	}
	if len(x.Flags) > 0 {
		x1 := make(testobj.TestFlag, len(x.Flags))
		for k, v := range x.Flags {
			var k1 string
			buf, k1 = bufferizeS(buf, k)
			x1[k1] = v
		}
		c.Flags = x1
	}
	if x.Finance != nil {
		c.Finance = &testobj.TestFinance{}
		c.Finance.MoneyIn = x.Finance.MoneyIn
		c.Finance.MoneyOut = x.Finance.MoneyOut
		c.Finance.Balance = x.Finance.Balance
		c.Finance.AllowBuy = x.Finance.AllowBuy
		if len(x.Finance.History) > 0 {
			c.Finance.History = make([]testobj.TestHistory, len(x.Finance.History))
			for i := 0; i < len(x.Finance.History); i++ {
				h := testobj.TestHistory{}
				h.DateUnix = x.Finance.History[i].DateUnix
				h.Cost = x.Finance.History[i].Cost
				buf, h.Comment = bufferize(buf, x.Finance.History[i].Comment)
				c.Finance.History[i] = h
			}
		}
	}

	return &c
}

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

	// b.Run("inspect bytes", func(b *testing.B) {
	// 	b.ReportAllocs()
	// 	for i := 0; i < b.N; i++ {
	// 		c := inspectBytes(origin)
	// 		p := make([]byte, 0, c)
	// 		_ = p
	// 	}
	// })
	b.Run("copy", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			c := cpy(origin)
			_ = c.Status
		}
	})
}
