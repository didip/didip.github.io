package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/didip/didip.github.io/libstring"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var nonWordMatcher = regexp.MustCompile(`\W`)

type Blog struct {
	Title      string
	CurrentDir string
}

func (b *Blog) ServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}
	return port
}

func (b *Blog) NewPostFilename(title string) string {
	title = strings.ToLower(title)
	title = strings.Replace(title, " ", "-", -1)

	for _, blankable := range []string{",", ".", "'"} {
		title = strings.Replace(title, blankable, "", -1)
	}

	title = fmt.Sprintf("%v-%v", time.Now().UnixNano(), title)

	nonWordMatcher.ReplaceAllString(title, "")
	title = title + ".md"

	return title
}

func (b *Blog) NewBlankPost(title string) error {
	postFilename := b.NewPostFilename(title)

	titleMd := []byte(fmt.Sprintf("## %v\n", title))

	err := ioutil.WriteFile(
		filepath.Join(
			libstring.ExpandTildeAndEnv(b.CurrentDir),
			"markdown",
			"posts",
			postFilename), titleMd, 0644)

	return err
}

func main() {
	blog := &Blog{}
	blog.Title = "Didip's Tech Mind"
	blog.CurrentDir = "$GOPATH/src/github.com/didip/didip.github.io"

	app := cli.NewApp()
	app.Name = "didip-blog"
	app.Usage = "It builds Didip's blog."
	app.Author = "Didip Kerabat"
	app.Email = "didipk@gmail.com"

	app.Commands = []cli.Command{
		{
			Name:      "server",
			ShortName: "serv",
			Usage:     "Run HTTP server",
			Action: func(c *cli.Context) {
				http.Handle("/", http.FileServer(http.Dir(".")))
				http.ListenAndServe(blog.ServerPort(), nil)
			},
		},
		{
			Name:  "posts",
			Usage: "Posts CRUD operations",
			Action: func(c *cli.Context) {
				crud := c.Args().First()

				if crud == "create" {
					title := c.Args().Get(1)
					blog.NewBlankPost(title)
				}
			},
		},
	}

	app.Run(os.Args)
}
