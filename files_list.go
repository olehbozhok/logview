package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type dirsAndFiles struct {
	Dirs  []string `json:"dirs"`
	Files []string `json:"files"`
}

func filesList(c *gin.Context) {
	path := c.Query("path")
	showHidedStr := c.Query("showHided")
	showHided := showHidedStr == "1"

	if path == "" {
		http.Error(c.Writer, "path is empty", http.StatusInternalServerError)
		return
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var dAndF dirsAndFiles
	dAndF.Dirs = make([]string, 0)
	dAndF.Files = make([]string, 0)
	for _, ent := range files {
		if !showHided {
			if strings.HasPrefix(ent.Name(), ".") {
				continue
			}
		}
		if ent.IsDir() {
			dAndF.Dirs = append(dAndF.Dirs, ent.Name())
			continue
		}
		dAndF.Files = append(dAndF.Files, ent.Name())
	}

	b, _ := json.Marshal(dAndF)
	c.Writer.Write(b)
}
