package engine

import (
	"fmt"
	"github.com/didip/didip.github.io/libstring"
	"io/ioutil"
	"os"
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
