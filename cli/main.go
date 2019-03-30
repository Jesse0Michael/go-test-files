package cli

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/atotto/clipboard"
)

const (
	// EnvDefaultKeyword sets a keyword to filter by when no arguments are given
	EnvDefaultKeyword = "TEST_FILES_DEFAULT"
)

func main() {
	rand.Seed(time.Now().Unix())
	files := StaticFiles()
	keywords := os.Args

	if d, ok := os.LookupEnv(EnvDefaultKeyword); ok && len(os.Args) == 0 {
		keywords = []string{d}
	}

	for _, key := range keywords {
		files.Filter(key)
	}

	file := files.Files[rand.Intn(len(files.Files))].URL()
	fmt.Println(file)
	clipboard.WriteAll(file)
}
