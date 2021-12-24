package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
)

type path struct {
	Key, Path string
}

var (
	datasets = []int{1, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
)

func f(d int, wg *sync.WaitGroup, errs chan error, paths chan path) {
	defer wg.Done()
	if rand.Int31n(1000) < 50 {
		err := errors.New("some error")
		errs <- err
	} else {
		p := path{
			Key:  "foo",
			Path: "bar",
		}
		paths <- p
	}
}

func main() {
	paths := make(map[string]string, len(datasets))
	errs := make(chan error, len(datasets))
	pathsChan := make(chan path, len(datasets))

	var wg sync.WaitGroup
	for _, d := range datasets {
		wg.Add(1)
		go f(d, &wg, errs, pathsChan)
	}

	var hasErr bool
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(datasets); {
			select {
			case p := <-pathsChan:
				fmt.Println("received path", p.Key)
				paths[p.Key] = p.Path
				i++
			case err := <-errs:
				fmt.Println("error caught", err)
				hasErr = true
				i++
			}
		}
	}()

	wg.Wait()
	close(pathsChan)
	close(errs)

	if !hasErr {
		fmt.Println("start ProcessMessage handling")
	} else {
		fmt.Println("skip ProcessMessage handling due to errors")
	}
}
