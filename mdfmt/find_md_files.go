package mdfmt

import (
	"os"
)

// take in file path return md files
type PathSet map[string]struct{}

func (p PathSet) Add(s string) {
	p[s] = struct{}{}
}

func (p PathSet) AddList(s []string) {
	for _, path := range s {
		p.Add(path)
	}
}

func (p PathSet) List() []string {
	paths := make([]string, 0, len(p))
	for path := range p {
		paths = append(paths, path)
	}
	return paths
}

func GetAllPathsInPaths(dirs []string) []string {
	p := make(PathSet)
	for _, dir := range dirs {
		p.AddList(getMdFilePathsInDir(dir))
	}
	return p.List()
}

func extendPathWithFile(dir string, entry string) string {
	if dir == "" {
		return entry
	} else if dir[len(dir)-1] == '/' {
		return dir + entry
	} else {
		return dir + "/" + entry
	}
}

func getMdFilePathsInDir(dir string) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return []string{}
	}
	ret := []string{}
	for _, entry := range entries {
		if entry.IsDir() {
			ret = append(ret, getMdFilePathsInDir(extendPathWithFile(dir, entry.Name()))...)
		} else if entry.Name()[len(entry.Name())-3:] == ".md" {
			ret = append(ret, extendPathWithFile(dir, entry.Name()))
		}
	}
	return ret
}
