package main

import (
	"filepath"
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type Blog struct {
	Title string
}

func (b *Blog) ServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}
	return port
}

func (b *Blog) CurrentDir() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

func (b *Blog) NewBlankPost(title string) error {
	dasherizedTitle := strings.ToLower(title)
	dasherizedTitle = strings.Replace(title, " ", "-", -1)

	titleMd := []byte(fmt.Sprintf("## %v\n", title))
	err := ioutil.WriteFile(filepath.Join(), titleMd, 0644)

}

func main() {
	blog := &Blog{}
	blog.Title = "Didip's Tech Mind"

	app = cli.NewApp()
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
					dasherizedTitle := strings.ToLower(title)
					dasherizedTitle = strings.Replace(title, " ", "-", -1)

					os.Create(filepath.Join())
				}
			},
		},
	}

}
