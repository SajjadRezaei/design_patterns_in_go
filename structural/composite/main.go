package composite

import "fmt"

type File struct {
	name string
}

type Folder struct {
	components []Component
	name       string
}

type Component interface {
	search(string)
}

//File methods

func (f *File) getName() string {
	return f.name
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s \n", keyword, f.name)
}

//Folder methods

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func Run() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		name: "Folder1",
	}

	folder1.add(file1)

	folder2 := &Folder{
		name: "Folder2",
	}

	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")
}
