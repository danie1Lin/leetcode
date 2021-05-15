package main

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

// ls 要按照字母排序
type File struct {
	name    string
	content string
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetNode(path string) Node {
	fmt.Println(f.name, path)
	if f.name == path {
		return f
	}
	return nil
}

func (f *File) String() string {
	return f.name
}

func (f *File) Ls(path string) []Node {
	if path == "." {
		return []Node{f}
	}
	return nil
}

func (f *File) AddNode(path string) {

}

type Dir struct {
	name  string
	nodes []Node
}

// "/a/b/c"
// d.name = "/a"
//
func (d *Dir) Ls(path string) []Node {
	if path == "." || path == "" {
		return d.nodes
	}
	node := d.GetNode(path)
	if node == nil {
		return nil
	}
	return node.Ls(".")
}

func (d *Dir) AddNode(path string) {
	if path == "" {
		return
	}
	childDirName, nextPath := getHeadDirAndNextPath(path)
	childDir := d.GetNode(childDirName)
	if childDir == nil {
		childDir = &Dir{name: childDirName}
		d.nodes = append(d.nodes, childDir)
	}
	childDir.AddNode(nextPath)
	fmt.Println(d, ":", d.nodes)
}

func (d *Dir) GetName() string {
	return d.name
}
func (d *Dir) String() string {
	return d.GetName()
}

func (this *Dir) GetNode(path string) Node {
	if path == "" || path == "." {
		return this
	}
	childDir, nextPath := getHeadDirAndNextPath(path)
	for i := range this.nodes {
		if this.nodes[i].GetName() == childDir {
			fmt.Println(this.name, "found", this.nodes[i].GetName())
			if nextPath == "." || nextPath == "" {
				return this.nodes[i]
			} else {
				return this.nodes[i].GetNode(nextPath)
			}
		}
	}
	return nil
}

type Node interface {
	Ls(path string) []Node
	GetName() string
	AddNode(path string)
	GetNode(path string) Node
	fmt.Stringer
}

type FileSystem struct {
	*Dir
}

func Constructor() FileSystem {
	return FileSystem{
		&Dir{name: "/"},
	}
}

func (this *FileSystem) Ls(path string) []string {
	if path[0] != '/' {
		return []string{}
	}
	nodes := this.Dir.Ls(path[1:])
	result := []string{}
	for i := range nodes {
		result = append(result, nodes[i].GetName())
	}
	sort.Strings(result)
	return result
}

func (this *FileSystem) Mkdir(path string) {
	if path[0] == '/' {
		this.AddNode(path[1:])
	}
}

func (this *FileSystem) GetNode(path string) Node {
	if path[0] == '/' {
		return this.Dir.GetNode(path[1:])
	}
	return this.Dir.GetNode(path)
}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
	fileName := filepath.Base(filePath)
	this.Mkdir(filepath.Dir(filePath))
	f := this.GetNode(filePath)
	if f == nil {
		dir := filepath.Dir(filePath)
		node := this.GetNode(dir)
		f = &File{
			name: fileName,
		}
		switch v := node.(type) {
		case *Dir:
			v.nodes = append(v.nodes, f)
		}
	}
	f.(*File).content += content
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	node := this.GetNode(filePath)
	switch v := node.(type) {
	case *File:
		return v.content
	}
	return ""
}

func getHeadDirAndNextPath(path string) (string, string) {
	dirs := strings.Split(path, "/")
	return dirs[0], strings.Join(dirs[1:], "/")
}

func main() {
	fs := Constructor()
	fs.AddContentToFile("/ab", "1aaaaaa")
	fmt.Println(fs.Ls("/"))
	fmt.Println(fs.ReadContentFromFile("/ab"))
}
