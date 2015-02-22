package main

import (
	"github.com/codegangsta/cli"
	"github.com/didip/didip.github.io/engine"
	"net/http"
	"os"
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
					blog.GeneratePostHTML(path)
				}
			},
		},
	}

	app.Run(os.Args)
}
