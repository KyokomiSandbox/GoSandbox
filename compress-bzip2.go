package main

import (
	"compress/bzip2"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func bzip2Exec() {
	bz2File, err := os.Open("text.tar.bz2")
	if err != nil {
		log.Fatal(err)
	}
	defer bz2File.Close()

	f := bzip2.NewReader(bz2File)

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}