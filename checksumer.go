package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func gensum(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	sum_string := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Printf("%v : %v\n", path, sum_string)
}

func main() {
	l := log.New(os.Stderr, "", 0)
	anonfunc := func(file string, info fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("error at [%v] at path [%q]", err, file)
			return err
		}
		if info.IsDir() {
			l.Printf("%v: is a directory\n", file)
		} else {
			gensum(file)
		}
		return nil
	}
	path := os.Args[1]
	filepath.WalkDir(path, anonfunc)
}
