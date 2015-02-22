package engine

import (
	"fmt"
	"github.com/didip/didip.github.io/libstring"
	"github.com/go-fsnotify/fsnotify"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var nonWordMatcher = regexp.MustCompile(`\W`)

func NewBlog(title string) *Engine {
	e := &Engine{}
	e.Title = title
	e.CurrentDir = "$GOPATH/src/github.com/didip/didip.github.io"

	return e
}

type Engine struct {
	Title      string
	CurrentDir string
}

func (e *Engine) ServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}
	return port
}

func (e *Engine) NewPostFilename(title string) string {
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

func (e *Engine) NewBlankPost(title string) error {
	postFilename := e.NewPostFilename(title)

	titleMd := []byte(fmt.Sprintf("## %v\n", title))

	err := ioutil.WriteFile(
		filepath.Join(
			libstring.ExpandTildeAndEnv(e.CurrentDir),
			"markdown",
			"posts",
			postFilename), titleMd, 0644)

	return err
}

func (e *Engine) GeneratePostHTML(inpath string) error {
	inpath = libstring.ExpandTildeAndEnv(inpath)

	filename := path.Base(inpath)
	filename = strings.Replace(filename, ".md", ".html", -1)

	markdownData, err := ioutil.ReadFile(inpath)
	if err != nil {
		return err
	}

	html := blackfriday.MarkdownCommon(markdownData)

	err = ioutil.WriteFile(
		filepath.Join(
			libstring.ExpandTildeAndEnv(e.CurrentDir),
			"posts",
			filename), html, 0644)

	return err
}

func (e *Engine) GenerateAllPostsHTML() error {
	mdDir := filepath.Join(
		libstring.ExpandTildeAndEnv(e.CurrentDir),
		"markdown",
		"posts")

	manyFileInfo, err := ioutil.ReadDir(mdDir)

	for _, fileinfo := range manyFileInfo {
		println(fileinfo.Name())
		e.GeneratePostHTML(filepath.Join(mdDir, fileinfo.Name()))
	}

	return err
}

func (e *Engine) WatchDir(path string, callback func(fsnotify.Event)) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					println(event.Name)
					callback(event)
				}

			case err := <-watcher.Errors:
				if err != nil {
					log.Println("Error: ", err)
				}
			}
		}
	}()

	err = watcher.Add(libstring.ExpandTildeAndEnv(path))
	if err != nil {
		return err
	}

	<-done
	return nil
}
