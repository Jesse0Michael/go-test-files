package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/atotto/clipboard"
	"github.com/jesse0michael/go-test-files/gofiles/content"
)

const (
	// EnvDefaultKeyword sets a keyword to filter by when no arguments are given
	EnvDefaultKeyword = "GOFILES_DEFAULT"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: testfiles [keywords...]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	rand.Seed(time.Now().Unix())
	files := content.StaticFiles()
	flag.Usage = usage
	flag.Parse()
	keywords := flag.Args()

	if d, ok := os.LookupEnv(EnvDefaultKeyword); ok && len(keywords) == 0 {
		keywords = []string{d}
	}

	for _, key := range keywords {
		files.Filter(key)
	}

	if len(files.Files) == 0 {
		fmt.Printf("Unable to find files matching: %v\n", keywords)
		os.Exit(1)
	}

	file := files.Files[rand.Intn(len(files.Files))].URL()
	fmt.Println(file)
	clipboard.WriteAll(file)
}
