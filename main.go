// Package main deals with the CLI interface and overall program control flow.
package main

import (
	"log"
	"os"
	"io/ioutil"
	"fmt"
	"strconv"
	// "net/http"

	// "github.com/GeertJohan/go.rice"
	"github.com/cdkini/clocviz/src/utils"
	"github.com/cdkini/clocviz/src/web"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("clocviz: Usage 'clocviz [src]'")
	}

	// Run cloc to generate file system and related stats
	/*
	raw, err := utils.RunCloc(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	*/
	b, err := ioutil.ReadFile(os.Args[1]) // just pass the file name
    if err != nil {
        fmt.Print(err)
    }
    raw := string(b)


	// Parse data and aggregate into object to be fed into template
	data := utils.ParseResults(raw)
	byLang := utils.GetLinesByLang(data)
	byFile := utils.GetLinesByFile(data)
	content := web.NewContent("Test", byLang, byFile)

	// Feed data into HTML/CSS/JS, start server, and render to browser
	port := 8080
	if len(os.Args) > 2 {
		port, err = strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
		fmt.Println(" you passed in a port value")
	}

	web.Serve(content, port)

	os.Exit(0)
}
