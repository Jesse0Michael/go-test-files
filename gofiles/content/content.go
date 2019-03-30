package content

import (
	"fmt"
	"strings"
)

const (
	githubFmt = "https://github.com/Jesse0Michael/test-files/blob/master/content/%s?raw=1"
)

//Content is a collection of all test files
type Content struct {
	Files []ContentFile
}

// Filter a list of Content Files to those that match a given keyword
func (c *Content) Filter(keyword string) {
	matching := []ContentFile{}
	for _, f := range c.Files {
		if f.Matches(keyword) {
			matching = append(matching, f)
		}
	}

	c.Files = matching
}

//ContentFile is an individual test file with identifying information
type ContentFile struct {
	Path     string
	Keywords []string
}

// Matches determines whether a given keyword matches this ContentFile
func (c *ContentFile) Matches(keyword string) bool {
	sanitizer := strings.NewReplacer("/", "", "-", "")
	key := sanitizer.Replace(keyword)

	if strings.Contains(sanitizer.Replace(c.Path), key) {
		return true
	}

	for _, k := range c.Keywords {
		if strings.Contains(sanitizer.Replace(c.Path), k) {
			return true
		}
	}

	return false
}

// URL returns the absolute url to the raw content file
func (c *ContentFile) URL() string {
	return fmt.Sprintf(githubFmt, c.Path)
}
