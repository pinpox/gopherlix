package main

import (
	"bytes"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"path"
)

func (server *GopherServer) createListing(reqPath string) (string, error) {
	log.Info("Creating listing for request: ", reqPath)

	var listing string

	// Handle directories
	if server.ServerRoot.DirExists(reqPath) {

		log.Info("Requested path ", reqPath, " exists")

		// Check if it contains a "index.gph" and serve it if it does
		if server.ServerRoot.FileExists(path.Join(reqPath, "index.gph")) {
			gopherMap, err := server.ServerRoot.GetServerFile(path.Join(reqPath, "index.gph"))

			if err != nil {
				return "", err
			}

			log.Info("Requested path", reqPath, " contains gophermap")
			return string(bytes.TrimRight(gopherMap, "\n")), nil
		}

		// If it is a directory without "index.gph", generate a menu from the contents
		log.Info("Requested path ", reqPath, " does not contain a gophermap, creating file list")
		files, err := server.ServerRoot.GetServerDir(reqPath)

		if err != nil {
			return "", err
		}

		for _, f := range files {
			link := ""
			link = server.createLink(server.ServerRoot.Type(path.Join(reqPath, f)), f, path.Join(reqPath, f))
			log.Info("Adding item: " + replaceCRLF(link))
			listing += link
		}

		// Add last dot
		listing += "."
		return listing, nil
	}

	// Handle files
	if server.ServerRoot.FileExists(reqPath) {
		gopherMap, err := server.ServerRoot.GetServerFile(reqPath)
		if err != nil {
			return "", err
		}

		return string(bytes.TrimRight(gopherMap, "\n")), nil
	}

	return "", errors.New("File or directory " + reqPath + " not found")
}

func (server *GopherServer) createLink(itemType, text, path string) string {
	return fmt.Sprintf("%s%s\t%s\t%s\t%s\r\n", itemTypes[itemType], text, path, server.Domain, server.Port)
}
