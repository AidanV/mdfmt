package main

import (
	"fmt"
	"mdfmt/mdfmt"
	"os"
	"sync"
)

type PathSet map[string]struct{}

func (p PathSet) Add(s string) {
	p[s] = struct{}{}
}

func (p PathSet) List() []string {
	paths := make([]string, 0, len(p))
	for path := range p {
		paths = append(paths, path)
	}
	return paths
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run mdfmt.go <file>")
		return
	}

	paths := make(PathSet)
	// create a set of all .md file paths
	for _, filePath := range os.Args[1:] {
		paths.Add(filePath)
	}

	var wg sync.WaitGroup

	handleReformat := func(path string) {
		defer wg.Done()
		in, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}
		fmt.Println(mdfmt.Reformat(string(in)))
	}

	for _, path := range paths.List() {
		wg.Add(1)
		go handleReformat(path)
	}
	wg.Wait()
}
