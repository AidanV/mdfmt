package mdfmt

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

func getMdFilePathsInDir(dir string) []string {
	return []string{"/home/aidan/Projects/mdfmt/testdata/input/All.md"}
}
