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
	e.MaxSummaryParagraphs = 3

	return e
}

type Engine struct {
	Title                string
	CurrentDir           string
	MaxSummaryParagraphs int
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
	if err != nil {
		return err
	}

	for _, fileinfo := range manyFileInfo {
		e.GeneratePostHTML(filepath.Join(mdDir, fileinfo.Name()))
	}

	return err
}

func (e *Engine) GenerateAllHTML() error {
	err := e.GenerateAllPostsHTML()
	if err != nil {
		return err
	}

	return e.GenerateIndexHTML()
}

func (e *Engine) MarkdownSummary(inpath string) (map[string]string, error) {
	data := make(map[string]string)

	markdownData, err := ioutil.ReadFile(libstring.ExpandTildeAndEnv(inpath))
	if err != nil {
		return data, err
	}

	text := make([]string, 0)
	newParagraphCounter := 0

	for i, line := range strings.Split(string(markdownData), "\n") {
		if i == 0 {
			data["Title"] = strings.Replace(line, "## ", "", 1)
			continue
		}

		text = append(text, line)

		if strings.TrimSpace(line) == "" {
			println(line)
			newParagraphCounter++
		}
		if newParagraphCounter > e.MaxSummaryParagraphs {
			break
		}
	}

	if len(text) > 0 {
		data["Text"] = strings.Join(text, "\n")
	}

	return data, err
}

func (e *Engine) GenerateIndexMarkdown() error {
	postsMarkdownDir := filepath.Join(
		libstring.ExpandTildeAndEnv(e.CurrentDir),
		"markdown",
		"posts")

	manyFileInfo, err := ioutil.ReadDir(postsMarkdownDir)
	if err != nil {
		return err
	}

	content := make([]string, 0)

	for _, fileinfo := range manyFileInfo {
		postLink := "/posts/" + fileinfo.Name()
		postLink = strings.Replace(postLink, ".md", ".html", -1)

		summary, err := e.MarkdownSummary(filepath.Join(postsMarkdownDir, fileinfo.Name()))
		if err != nil {
			continue
		}

		// Generate Link
		content = append(content, fmt.Sprintf("[<h2>%v</h2>](%v)\n", summary["Title"], postLink))

		// Generate Summary
		if _, ok := summary["Text"]; ok {
			content = append(content, fmt.Sprintf("%v\n", summary["Text"]))
		}
	}

	indexMarkdownFilename := filepath.Join(
		libstring.ExpandTildeAndEnv(e.CurrentDir),
		"markdown",
		"index.md")

	err = ioutil.WriteFile(indexMarkdownFilename, []byte(strings.Join(content, "\n")), 0644)

	return err
}

func (e *Engine) GenerateIndexHTML() error {
	err := e.GenerateIndexMarkdown()
	if err != nil {
		return err
	}

	inpath := filepath.Join(
		libstring.ExpandTildeAndEnv(e.CurrentDir),
		"markdown",
		"index.md")

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
			filename), html, 0644)

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
