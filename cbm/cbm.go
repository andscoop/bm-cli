package cbm

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type ChromeBookmarks struct {
	Roots struct {
		BookmarkBar Child `json:"bookmark_bar"`
	} `json:"roots"`
}

type Child struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Url      string  `json:"url"`
	Children []Child `json:"children"`
	Path     string
}

func ScrubPath(s string) string {
	// todo dedeupe "_"
	re, err := regexp.Compile(`[^\w/]`)
	if err != nil {
		log.Fatal(err)
	}
	return strings.ToLower(re.ReplaceAllString(s, "_"))
}

func (c *Child) Walk(path string) error {
	c.Path = ScrubPath(fmt.Sprintf("%s/%s", c.Path, c.Name))
	if len(c.Children) == 0 {
		return nil
	} else {
		for i, _ := range c.Children {
			c.Children[i].Walk(c.Path)
		}
	}

	return nil
}
