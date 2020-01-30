package content

import (
	"fmt"
	"strings"
)

const (
	githubFmt = "https://raw.githubusercontent.com/Jesse0Michael/go-test-fiiles/master/content/%s"
)

//Content is a collection of all test files
type Content struct {
	Files []File
}

// Filter a list of Content Files to those that match a given keyword
func (c *Content) Filter(keyword string) {
	matching := []File{}
	for _, f := range c.Files {
		if f.Matches(keyword) {
			matching = append(matching, f)
		}
	}

	c.Files = matching
}

//File is an individual test file with identifying information
type File struct {
	Path     string
	Keywords []string
}

// Matches determines whether a given keyword matches this ContentFile
func (c *File) Matches(keyword string) bool {
	sanitizer := strings.NewReplacer("/", "", "-", "")
	key := sanitizer.Replace(keyword)

	if strings.Contains(sanitizer.Replace(c.Path), key) {
		return true
	}

	for _, k := range c.Keywords {
		if strings.Contains(sanitizer.Replace(k), key) {
			return true
		}
	}

	return false
}

// URL returns the absolute url to the raw content file
func (c *File) URL() string {
	return fmt.Sprintf(githubFmt, c.Path)
}
