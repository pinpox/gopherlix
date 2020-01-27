package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path"
)

func createListing(reqPath string) (string, error) {
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
			return fmt.Sprint(string(gopherMap)), nil
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
				link = createLink("MENU", f.Name(), path.Clean(path.Join(f.Name())))
			} else {
				link = createLink("TEXT", f.Name(), path.Clean(path.Join(f.Name())))
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
		return fmt.Sprint(string(gopherMap)), nil
	}

	return listing, nil

}

func createLink(itemType, text, path string) string {
	return fmt.Sprintf("%s%s\t%s\t%s\t%s\r\n", itemTypes[itemType], text, path, CONN_DOMAIN, CONN_PORT)
}
