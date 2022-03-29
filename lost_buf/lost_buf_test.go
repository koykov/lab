package lost_buf

import (
	"testing"

	"github.com/koykov/bytebuf"
)

func WriteData(dst []byte, prefix, suffix []byte, id int32) {
	cb := bytebuf.ChainBuf(dst)
	cb.WriteStr("foo")
	cb.Write(prefix)
	cb.WriteStr("bar")
	cb.WriteInt(int64(id))
	cb.Write(suffix)
}

func AppendData(dst []byte, prefix, suffix []byte, id int32) []byte {
	cb := bytebuf.ChainBuf(dst)
	cb.WriteStr("foo")
	cb.Write(prefix)
	cb.WriteStr("bar")
	cb.WriteInt(int64(id))
	cb.Write(suffix)
	return cb
}

func TestLostBuf(t *testing.T) {
	t.Run("write", func(t *testing.T) {
		buf := make([]byte, 15)
		WriteData(buf[:0], []byte("aaa"), []byte("bbb"), 12345678)
		t.Log(string(buf))
	})
	t.Run("append", func(t *testing.T) {
		buf := make([]byte, 15)
		buf = AppendData(buf[:0], []byte("aaa"), []byte("bbb"), 12345678)
		t.Log(string(buf))
	})
}
