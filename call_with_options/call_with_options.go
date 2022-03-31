package call_with_options

import "github.com/koykov/inspector"

type T struct {
	l, r string
	buf  []byte
}

func (t *T) F(s string, o OptionsInterface) {
	t.r = s
	raw := o.GetOption("inspector")
	if ins, ok := raw.(inspector.Inspector); ok {
		ins.DeepEqual(&t.l, &t.r)
	}
	t.buf = append(t.buf, s...)
}

func (t *T) reset() {
	t.buf = t.buf[:0]
}
