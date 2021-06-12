package design

import (
	"sort"
	"strings"
)

type Obj struct {
	isfile  bool
	content string
	sub     map[string]*Obj
}

func newObj() *Obj {
	return &Obj{
		isfile:  false,
		content: "",
		sub:     make(map[string]*Obj),
	}
}

type FileSystem struct {
	root *Obj
}

func Constructor() FileSystem {
	return FileSystem{newObj()}
}

func (f *FileSystem) Ls(path string) []string {
	res := []string{}
	cur := f.root

	if path != "/" {
		parsed := strings.Split(path, "/")
		for i := 1; i < len(parsed); i++ {
			cur = cur.sub[parsed[i]]
		}
		if cur.isfile {
			res = append(res, parsed[len(parsed)-1])
			return res
		}
	}
	list := []string{}
	for k := range cur.sub {
		list = append(list, k)
	}
	sort.Strings(list)
	return list
}

func (f *FileSystem) Mkdir(path string) {
	cur := f.root
	parsed := strings.Split(path, "/")
	n := len(parsed)
	for i := 1; i < n; i++ {
		if _, ok := cur.sub[parsed[i]]; !ok {
			cur.sub[parsed[i]] = newObj()
		}
		cur = cur.sub[parsed[i]]
	}
}

func (f *FileSystem) AddContentToFile(filePath string, content string) {
	cur := f.root
	parsed := strings.Split(filePath, "/")
	n := len(parsed)
	for i := 1; i < n-1; i++ {
		cur = cur.sub[parsed[i]]
	}
	if _, ok := cur.sub[parsed[len(parsed)-1]]; !ok {
		cur.sub[parsed[len(parsed)-1]] = newObj()
	}
	cur = cur.sub[parsed[len(parsed)-1]]
	cur.isfile = true
	cur.content += content
}

func (f *FileSystem) ReadContentFromFile(filePath string) string {
	cur := f.root
	parsed := strings.Split(filePath, "/")
	n := len(parsed)
	for i := 1; i < n-1; i++ {
		cur = cur.sub[parsed[i]]
	}
	return cur.sub[parsed[len(parsed)-1]].content
}
