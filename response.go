package main

import (
	// "bytes"
	// "errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"path"
)

func (server *GopherServer) createListing(reqPath string) (string, error) {
	log.Info("Creating listing for request: ", reqPath)

	var listing string

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

	return listing, nil
}

func (server *GopherServer) createLink(itemType, text, path string) string {
	return fmt.Sprintf("%s%s\t%s\t%s\t%s\r\n", itemTypes[itemType], text, path, server.Domain, server.Port)
}
