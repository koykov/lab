package main

import (
	"log"

	"github.com/koykov/cbyte"
)

const allocLimit = uint64(1024 * 1024 * 1024 * 16)

func main() {
	bytesCount := uint64(1)
	h := cbyte.InitHeader64(bytesCount, bytesCount)
	log.Printf("  alloc %12d bytes at addr #%x", bytesCount, h.Data)
	for bytesCount <= allocLimit {
		bytesCount *= 2
		h.Len, h.Cap = bytesCount, bytesCount
		h.Data = uintptr(cbyte.GrowHeader64(h))
		if h.Data == 0 {
			log.Fatalf("realloc %12d bytes failed", bytesCount)
			return
		}
		log.Printf("realloc %12d bytes at addr #%x", bytesCount, h.Data)
	}
	log.Println("done")
	cbyte.ReleaseHeader64(h)
}
