package main

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path"
)

func (server *GopherServer) createListing(reqPath string) (string, error) {
	log.Info("Creating listing for request: ", reqPath)

	var listing string

	// Handle directories
	if dirExists(reqPath) {

		log.Info("Requested path ", reqPath, " exists")

		// Check if it contains a "index.gph" and serve it if it does
		if fileExists(path.Join(reqPath, "index.gph")) {
			gopherMap, err := ioutil.ReadFile(path.Join(reqPath, "index.gph"))

			if err != nil {
				log.Fatal(err)
			}

			log.Info("Requested path", reqPath, " contains gophermap")
			return string(bytes.TrimRight(gopherMap, "\n")), nil
		}

		// If it is a directory without "index.gph", generate a menu from the contents
		log.Info("Requested path ", reqPath, " does not contain a gophermap, creating file list")
		files, err := ioutil.ReadDir(reqPath)

		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			link := ""
			if f.IsDir() {
				link = server.createLink("MENU", f.Name(), path.Clean(path.Join(f.Name())))
			} else {
				link = server.createLink("TEXT", f.Name(), path.Clean(path.Join(f.Name())))
			}
			log.Info("Adding item: " + link)
			listing += link
		}

		// Add last dot
		listing += "."
	}

	// Handle files
	if fileExists(reqPath) {
		gopherMap, err := ioutil.ReadFile(reqPath)
		if err != nil {
			log.Fatal(err)
		}

		log.Info("Requested path", reqPath, " contains gophermap")
		return string(bytes.TrimRight(gopherMap, "\n")), nil
	}

	return listing, nil

}

func (server *GopherServer) createLink(itemType, text, path string) string {
	return fmt.Sprintf("%s%s\t%s\t%s\t%s\r\n", itemTypes[itemType], text, path, server.Domain, server.Port)
}
