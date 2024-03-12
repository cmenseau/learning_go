package main

import (
	"fmt"
	"sort"
	"time"
)

type base struct {
	num int
}

type container struct {
	base
	str string
}

func (b base) String() string {
	return fmt.Sprintf("base=%d", b.num)
}

func (c container) String() string {
	return fmt.Sprintf("container={%s, %s}", c.base, c.str)
}

func incrBase(b base) int {
	return b.num + 1
}

func (b base) incrBase() int {
	return b.num + 1
}

type incrementer interface {
	incrBase() int
}

func container_base() {
	var c container = container{base{41}, "wtf"}

	fmt.Println(c) // uses base.String() if container.String() doesn't exist
	fmt.Println(c.base.num)
	fmt.Println(c.num)
	// base=41
	// 41
	// 41

	// not working !!! struct composition is not substitution!!!!
	// incrBase(c)
	fmt.Println(incrBase(c.base)) // 42

	fmt.Println(c.base.incrBase()) // 42
	fmt.Println(c.incrBase())      // 42

	i := incrementer(c)
	j := incrementer(c.base)
	fmt.Println(i) //container={base=41, wtf}
	fmt.Println(j) //base=41
}

////////////////////////////////////

type StringList []string

type EmbedingIntAndNamedType struct {
	int
	StringList
	str string
}

func run_embedding_int_named_type() {
	var v EmbedingIntAndNamedType = EmbedingIntAndNamedType{12, StringList{"a", "b"}, "hey"}
	fmt.Println(v)
	var i int = 1
	i = i + v.int
	fmt.Println(i)
}

////////////////////////////////////

type ContainerPtrBase struct {
	*base
	str string
}

func (b *base) incrBasePtrSafe() int {
	if b == nil {
		return 0
	}
	return b.num + 1
}

func container_ptr_base() {
	var c ContainerPtrBase = ContainerPtrBase{}

	//fmt.Println(c.incrBase())
	// panic segmentation violation
	// nothings prevents user from calling method on nil value
	fmt.Println(c.incrBasePtrSafe())

}

////////////////////////////////////

type MyFile struct {
	name       string
	size       int
	lastUpdate time.Time
}

type MyFiles []MyFile

func (files MyFiles) Len() int {
	return len(files)
}

func (files MyFiles) Swap(a, b int) {
	files[a], files[b] = files[b], files[a]
}

type ByName struct {
	MyFiles
}

func (n ByName) Less(a, b int) bool {
	return n.MyFiles[a].name < n.MyFiles[b].name
}

type BySize struct {
	MyFiles
}

func (n BySize) Less(a, b int) bool {
	return n.MyFiles[a].size < n.MyFiles[b].size
}

func file_composition_sorting() {
	var files = MyFiles{
		MyFile{name: "example.txt", size: 1234, lastUpdate: time.Now()},
		MyFile{"prog.go", 540, time.Date(2024, time.January, 1, 0, 0, 0, 0, time.Now().Local().Location())},
	}

	sort.Sort(ByName{files})
	fmt.Println(files)
	sort.Sort(BySize{files})
	fmt.Println(files)
}

func main() {

	// container_base()
	//run_embedding_int_named_type()
	//container_ptr_base()
	file_composition_sorting()
}
