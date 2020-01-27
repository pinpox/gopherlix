package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func createListing(reqPath string) (string, error) {

	listing := ""

	// Check if requested path is a directory
	if dirExists(reqPath) {

		// Check if it contains a "index.gph" and serve it if it does
		if fileExists(path.Join(reqPath, "index.gph")) {
			gopherMap, err := os.Open(path.Join(reqPath, "index.gph"))
			if err != nil {
				log.Fatal(err)
			}
			return fmt.Sprint(gopherMap), nil
		}

		// If it is a directory without "index.gph", generate a menu from the contents
		files, err := ioutil.ReadDir(reqPath)

		if err != nil {
			return "", err
		}

		for _, f := range files {
			if f.IsDir() {
				listing += createLink("MENU", f.Name(), path.Join(SERVER_ROOT, reqPath, f.Name()))
			} else {
				listing += createLink("TEXT", f.Name(), path.Join(SERVER_ROOT, reqPath, f.Name()))
			}
		}

	}

	return listing, nil
}

func createLink(itemType, text, path string) string {
	return fmt.Sprintf("%s%s\t%s\t%s%s\r\n", itemTypes[itemType], text, path, CONN_DOMAIN, CONN_PORT)
}
