package list

import (
	"fmt"
	"html"
	"net/url"
	"os"
	"path"
	"strings"
)

// Item documents list item
type Item struct {
	DisplayName string
	SrcName     string
	OutName     string
	Sub         []Item
}

// List documents list
type List []Item

// New return new List
func New(files []os.FileInfo) List {
	list := List{}

	for _, f := range files {
		fname := f.Name()

		// skip directories
		if f.IsDir() {
			fmt.Printf("%s is Directory, skipped.\n", fname)
			continue
		}

		// check if markdown file
		ext := strings.ToLower(path.Ext(fname))
		switch ext {
		case ".md", ".mdown", ".markdown":
		default:
			fmt.Printf("%s is not markdown file, skipped.\n", fname)
			continue
		}

		clearName := fname[0:strings.LastIndex(fname, ext)]
		outName := clearName + ".html"

		// process doc level
		names := strings.Split(clearName, "_")
		namesLength := len(names)
		if namesLength == 1 {
			if i := list.Index(names[0]); i != -1 {
				list[i].SrcName = fname
				list[i].OutName = outName
			} else {
				list = append(list, Item{
					DisplayName: names[0],
					SrcName:     fname,
					OutName:     outName,
				})
			}
		} else {
			if namesLength > 2 {
				names = []string{names[0], strings.Join(names[1:namesLength], "_")}
			}

			item := Item{
				DisplayName: names[1],
				SrcName:     fname,
				OutName:     outName,
			}

			if i := list.Index(names[0]); i != -1 {
				list[i].Sub = append(list[i].Sub, item)
			} else {
				list = append(list, Item{
					DisplayName: names[0],
					Sub:         []Item{item},
				})
			}
		}
	}

	return list
}

// Index return -1 if not found
func (l List) Index(displayName string) int {
	for k, v := range l {
		if v.DisplayName == displayName {
			return k
		}
	}

	return -1
}

// SetHomePage check index.html
func (l List) SetHomePage() {
	for _, v := range l {
		if v.DisplayName == "index" {
			return
		}
	}

	for k, v := range l {
		if strings.ToLower(v.DisplayName) == "readme" {
			l[k].OutName = "index.html"
			return
		}
	}
}

// ToHTML format list to html tags
func (l List) ToHTML() string {
	out := "<ul>"

	for _, v := range l {
		out += "<li>"

		if v.SrcName == "" {
			out += fmt.Sprintf("<a class=\"h\">%s</a>", html.EscapeString(v.DisplayName))
		} else {
			out += fmt.Sprintf("<a class=\"h\" href=\"%s\">%s</a>",
				url.QueryEscape(v.OutName),
				html.EscapeString(v.DisplayName),
			)
		}

		if len(v.Sub) > 0 {
			out += "<ul>"
			for _, w := range v.Sub {
				out += fmt.Sprintf("<li><a href=\"%s\">%s</a></li>",
					url.QueryEscape(w.OutName),
					html.EscapeString(w.DisplayName),
				)
			}
			out += "</ul>"
		}

		out += "</li>"
	}
	out += "</ul>"

	return out
}
