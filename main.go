package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()

	dir := flag.Arg(0)
	if dir == "" {
		fmt.Printf("Usage: project-doc <parent-directory-of-document-sources>\n\n")
		return
	}

	// Get absolute path
	absDir, err := filepath.Abs(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	src := absDir + "/src/"
	out := absDir + "/out/"

	// check path
	outDirInfo, err := os.Stat(out)
	if err != nil {
		if err := os.Mkdir(out, 0777); err != nil {
			fmt.Println(err)
			return
		}
	} else {
		if !outDirInfo.IsDir() {
			fmt.Printf("%s is not directory.", out)
			return
		}
	}

	fmt.Printf("ABSOLUTE PATH: %s\n", absDir)

	if err := generateDoc(absDir, src, out); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Congratulation, the documents generated.\n\n")
}
