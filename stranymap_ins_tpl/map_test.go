package stranymap_ins_tpl

import (
	"bytes"
	"testing"

	"github.com/koykov/dyntpl"
	"github.com/koykov/inspector"
)

const (
	tpl = `before{
{% ctx list = m.x.y.([]string) %}
{% for _, x := range list separator | %}
  {%= x %}
{% endfor %}
}after`
	expect = "before{1|2|3}after"
)

func init() {
	tree, _ := dyntpl.Parse([]byte(tpl), false)
	dyntpl.RegisterTpl(-1, "map_", tree)
}

var (
	m = map[string]any{
		"x": map[string]any{
			"y": []string{"1", "2", "3"},
		},
	}
)

func TestMap(t *testing.T) {
	var ins inspector.StringAnyMapInspector
	ctx := dyntpl.AcquireCtx()
	defer dyntpl.ReleaseCtx(ctx)
	ctx.Set("m", m, ins)

	var buf bytes.Buffer
	_ = dyntpl.Write(&buf, "map_", ctx)
	if buf.String() != expect {
		t.FailNow()
	}
}
