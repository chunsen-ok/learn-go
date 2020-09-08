/*
go standard package
	builtin V
	fmt V
	errors V
	strings
	strconv
	math
	time
*/
package main

import (
	"archive/zip"
	_ "bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var src = flag.String("src", "", "src path to zip")

func main() {
	flag.Parse()

	if len(*src) == 0 {
		log.Fatal("No input.")
	}

	var files []string
	filepath.Walk(*src, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fmt.Println("find: ", path)
			files = append(files, path)
		}
		return nil
	})

	fmt.Println(files)

	// compress("zz.zip", files)
	// decompress("zz.zip")
}

func compress(dst string, files []string) {
	ar, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	w := zip.NewWriter(ar)
	defer w.Close()

	for _, file := range files {
		out, err := w.Create(file)
		if err != nil {
			log.Fatal(err)
		}

		b, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		_, err = out.Write(b)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func decompress(zFile string) {
	r, err := zip.OpenReader(zFile)
	if err != nil {
		log.Fatal(err)
	}

	d := filepath.Base(zFile)
	db := strings.Split(d, ".")
	os.Mkdir(db[0], os.ModeDir)

	for _, f := range r.File {
		log.Println("read file: ", f.Name)

		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}

		fd := filepath.Dir(f.Name)
		os.Mkdir(db[0]+"/"+fd, os.ModeDir)

		fmt.Println("dst:", db[0]+"/"+f.Name)
		of, err := os.Create(db[0] + "/" + f.Name)
		if err != nil {
			log.Fatal(err)
		}
		defer of.Close()

		_, err = io.Copy(of, rc)
		if err != nil {
			log.Fatal(err)
		}
	}
}
