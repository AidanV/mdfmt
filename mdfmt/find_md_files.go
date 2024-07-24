package mdfmt

import "os"

// take in file path return md files
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

func GetAllPathsInPaths(dirs []string) []string {
	return []string{"input/All.md"}
}

func getMdFilePathsInDir(dir string) []string {
	paths := make(PathSet)
	// create a set of all .md file paths
	for _, filePath := range os.Args[1:] {
		paths.Add(filePath)
	}
	return []string{"input/All.md"}
}
