package main

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
)

type showLogContext struct {
	Filepath   string
	Filters    []string
	ResultStrs []string
}

func showLog(box *packr.Box) func(*gin.Context) {

	s, err := box.FindString("templates/show_log.html")
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("show_log").Parse(s)
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		var tplCtx showLogContext

		filePath := c.Query("fp")
		filterJSON := c.Query("filter")

		tplCtx.Filepath = filePath

		var filterList []string
		err := json.Unmarshal([]byte(filterJSON), &filterList)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}
		tplCtx.Filters = filterList

		res, err := readLinesFillter(filePath, filterList)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		tplCtx.ResultStrs = res

		err = tmpl.Execute(c.Writer, tplCtx)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func readLinesFillter(filepath string, filter []string) (res []string, err error) {
	clarifFilrer := make([]string, 0, len(filter))
	for _, f := range filter {
		if f != "" {
			clarifFilrer = append(clarifFilrer, f)
		}
	}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 1024*1024)
	scanner.Buffer(buf, 10*1024*1024)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		for _, f := range clarifFilrer {
			if !strings.Contains(text, f) {
				continue
			}
			res = append(res, text)
		}
	}
	return res, nil
}
