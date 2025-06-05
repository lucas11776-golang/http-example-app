package prompt

import (
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/lucas11776-golang/http/utils/path"
	"github.com/open2b/scriggo"
	"github.com/open2b/scriggo/native"
)

type writer struct {
	parsed []byte
}

// Comment
func (ctx *writer) Write(p []byte) (n int, err error) {
	ctx.parsed = append(ctx.parsed, p...)

	return len(ctx.parsed), nil
}

// Comment
func (ctx *writer) Parsed() []byte {
	return ctx.parsed
}

type Prompt struct {
	fs fs.FS
}

type PromptData map[string]interface{}

type filesystem struct {
	dir       string
	extension string
	// cache     scriggo.Files // TODO: Thinking of cache but if restart what to change job description.
}

// Comment
func (ctx *filesystem) Open(name string) (fs.File, error) {
	return os.Open(fmt.Sprintf("%s/%s", ctx.dir, path.FileRealPath(name, ctx.extension)))
}

// Comment
func NewPrompt(dir string, extension string) *Prompt {
	return &Prompt{
		fs: &filesystem{
			dir:       strings.TrimRight(dir, "/"),
			extension: extension,
			// cache:     make(scriggo.Files),
		},
	}
}

// Comment
func (ctx *Prompt) Generate(name string, data PromptData) (string, error) {
	globals := native.Declarations{
		"formatTime": func(time time.Time) string {
			return strings.Join([]string{strconv.Itoa(time.Day()), time.Month().String(), strconv.Itoa(time.Year())}, " ")
		},
	}

	for key, value := range data {
		globals[key] = value
	}

	template, err := scriggo.BuildTemplate(ctx.fs, name, &scriggo.BuildOptions{Globals: globals})

	if err != nil {
		return "", err
	}

	writer := &writer{}

	if err := template.Run(writer, nil, nil); err != nil {
		return "", err
	}

	return string(writer.Parsed()), nil
}
