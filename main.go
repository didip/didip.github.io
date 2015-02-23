package main

import (
	"github.com/codegangsta/cli"
	"github.com/didip/didip.github.io/engine"
	"github.com/didip/didip.github.io/engine/dirwatcher"
	"github.com/go-fsnotify/fsnotify"
	"net/http"
	"os"
	"path"
)

func main() {
	blog := engine.NewBlog("Didip's Tech Mind")

	app := cli.NewApp()
	app.Name = "didip.github.io"
	app.Usage = "builds Didip's blog."
	app.Author = "Didip Kerabat"
	app.Email = "didipk@gmail.com"
	app.Version = "9001.0.0"

	app.Commands = []cli.Command{
		{
			Name:      "server",
			ShortName: "serv",
			Usage:     "Run HTTP server",
			Action: func(c *cli.Context) {
				// Watches markdown/posts/ directory and auto-generate HTML
				go dirwatcher.WatchDir(path.Join(blog.CurrentDir, "markdown", "posts"), func(event fsnotify.Event) {
					blog.GenerateAllHTML()
				})

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

				} else if crud == "generate" {
					path := c.Args().Get(1)

					if path == "" {
						blog.GenerateAllHTML()
					} else {
						blog.GeneratePostHTML(path)
					}
				}
			},
		},
	}

	app.Run(os.Args)
}
