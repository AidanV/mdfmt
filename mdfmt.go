package main

import (
	"fmt"
	"mdfmt/mdfmt"
	"os"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run mdfmt.go <file>")
		return
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

	for _, path := range mdfmt.GetAllPathsInPaths(os.Args[1:]) {
		wg.Add(1)
		go handleReformat(path)
	}
	wg.Wait()
}
