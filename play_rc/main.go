package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

var (
	weight = map[string]int{
		"regionDataset":       0,
		"provinceDataset":     1,
		"municipalityDataset": 2,
		"zipCodeDataset":      3,
		"roadDataset":         4,
		"addressDataset":      5,
		"unitDataset":         6,
		"entryAddressDataset": 7,
		"buildingDataset":     8,
	}

	// map source for example, just imagine you get file key and file path from somewhere else
	source = map[string]string{
		"regionDataset":       "path/to/file0",
		"provinceDataset":     "path/to/file1",
		"municipalityDataset": "path/to/file2",
		"zipCodeDataset":      "path/to/file3",
		"roadDataset":         "path/to/file4",
		"addressDataset":      "path/to/file5",
		"unitDataset":         "path/to/file6",
		"entryAddressDataset": "path/to/file7",
		"buildingDataset":     "path/to/file8",
	}

	files Files
)

type File struct {
	key, fn string
	weight  int
}

type Files []File

func (f *Files) Len() int {
	return len(*f)
}

func (f *Files) Less(i, j int) bool {
	return (*f)[i].weight < (*f)[j].weight
}

func (f *Files) Swap(i, j int) {
	(*f)[i], (*f)[j] = (*f)[j], (*f)[i]
}

func (f *Files) String() string {
	var (
		buf bytes.Buffer
		l   = len(*f)
	)
	_ = (*f)[l-1]
	for i := 0; i < l; i++ {
		x := (*f)[i]
		buf.WriteString(strconv.Itoa(x.weight))
		buf.WriteString(". ")
		buf.WriteString(x.key)
		buf.WriteString(": ")
		buf.WriteString(x.fn)
		buf.WriteByte('\n')

	}
	return buf.String()
}

func main() {
	// read files and keys, for example from map
	fmt.Println("files came in order:")
	for k, f := range source {
		fmt.Printf("%s: %s\n", k, f)
		files = append(files, File{
			key:    k,
			fn:     f,
			weight: weight[k],
		})
	}
	sort.Sort(&files)
	fmt.Println("\nfiles after sort")
	fmt.Println(&files)
}
