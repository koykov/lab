package iface_copy

import "testing"

func TestIfaceCopy(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		ins := TI{}
		x := &T{
			id: 15,
			ct: 3.1415,
			nm: "foo",
			pl: []byte("bar"),
		}
		x1, _ := ins.Copy(x)
		x.id = 16
		t.Log(x1)
	})
	t.Run("scalar", func(t *testing.T) {
		ins := SI{}
		x := float32(3.1415)
		x1, _ := ins.Copy(&x)
		x = 15.16
		t.Log(x1)
	})
	t.Run("string", func(t *testing.T) {
		ins := SI{}
		x := "foobar"
		x1, _ := ins.Copy(&x)
		x = "qwerty"
		t.Log(x1)
	})
	t.Run("bytes", func(t *testing.T) {
		ins := SI{}
		x := []byte("foobar")
		x1, _ := ins.Copy(&x)
		x[0] = 'q'
		x[1] = 'w'
		x[2] = 'e'
		t.Log(string(x1.([]byte)))
	})
}
