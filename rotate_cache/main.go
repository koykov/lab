package main

import (
	"log"
	"time"
)

var c rotateCache

func init() {
	for i := 0; i < 2; i++ {
		c.set(0, "foo")
		c.set(1, "bar")
		c.set(2, "qwe")
		c.set(3, "asd")
		c.rotate()
	}
}

func main() {
	done := make(chan struct{})
	go func() {
		var s string
		for {
			select {
			case <-done:
				log.Println("reader stop")
				return
			default:
				if s = c.get(0); s != "foo" {
					log.Println("foo mismatch")
				}
				if s = c.get(1); s != "bar" {
					log.Println("bar mismatch")
				}
				if s = c.get(2); s != "qwe" {
					log.Println("qwe mismatch")
				}
				if s = c.get(3); s != "asd" {
					log.Println("asd mismatch")
				}
			}
		}
	}()

	t := time.NewTicker(time.Second)
	for i := 0; i < 10; i++ {
		select {
		case <-t.C:
			c.resetBuf()
			c.set(0, "foo")
			c.set(1, "bar")
			c.set(2, "qwe")
			c.set(3, "asd")
			c.rotate()
			log.Println("rotate cache")
		}
	}
	done <- struct{}{}
}
