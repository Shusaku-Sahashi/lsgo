package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

var options *opts

type opts struct {
	immediateDirs bool
}

var (
	cwdFiles    []fileInfo = make([]fileInfo, 0, 100)
	cwdUsed     uint       = 0
	cwdAlloced  uint       = 100
	sortedFiles []*fileInfo
)

type fileInfo struct {
	name string
	// filetype filetype
	stat os.FileInfo
}

func main() {
	opts := opts{}
	flag.BoolVar(&(opts.immediateDirs), "d", false, "")

	flag.Parse()

	exec(&opts, flag.Args())
}

func exec(opts *opts, files []string) {
	options = opts

	if len(files) <= 0 {
		if options.immediateDirs {
			globFiles(".")
		} else {
			panic("Not Implemmented")
		}
	} else {
		panic("Not Implemmented.")
	}

	if len(cwdFiles) > 0 {
		sortFiles()
	}

	for _, file := range sortedFiles {
		fmt.Fprintln(os.Stdout, file.stat.Name())
	}
}

func initilizeOrderSlice() {
	for i := range cwdFiles {
		sortedFiles = append(sortedFiles, &cwdFiles[i])
	}
}

func sortFiles() {
	initilizeOrderSlice()
	sort.SliceStable(sortedFiles, func(i, j int) bool {
		return sortedFiles[i].name < sortedFiles[j].name
	})
}

func globFiles(name string) {
	f := fileInfo{}
	f.name = name
	// f.filetype = type

	stat, err := os.Lstat(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%q: No such file or directory\n", name)
	}
	f.stat = stat

	cwdFiles = append(cwdFiles, f)
}
