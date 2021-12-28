package cbm

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// todo print of bookmarks (flat, nested?)

type ChromeBookmarks struct {
	Roots struct {
		BookmarkBar Child `json:"bookmark_bar"`
	} `json:"roots"`
}

type Child struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Url      string  `json:"url"`
	Children []Child `json:"children"`
	Path     string
}

func unmarshalBookmarks(path string) (*ChromeBookmarks, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var bm ChromeBookmarks
	err = json.Unmarshal(file, &bm)
	if err != nil {
		return nil, err
	}

	return &bm, nil
}

func FlatList(fn string) error {
	bm, err := unmarshalBookmarks(fn)
	if err != nil {
		return err
	}

	return bm.Roots.BookmarkBar.flatList("")
}

func (c *Child) flatList(path string) error {
	c.Path = fmt.Sprintf("%s/%s", path, c.Name)
	if len(c.Children) == 0 {
		fmt.Printf("%s: %s\n", c.Id, c.Path)
	} else {
		for i, _ := range c.Children {
			c.Children[i].flatList(c.Path)
		}
	}

	return nil
}

func ScrubPath(s string) string {
	// todo dedeupe "_"
	re, err := regexp.Compile(`[^\w/]`)
	if err != nil {
		log.Fatal(err)
	}
	return strings.ToLower(re.ReplaceAllString(s, "_"))
}
