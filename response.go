package main

import (
	"bytes"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"path"
	"strings"
	"text/template"
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

	return strings.TrimRight(listing, "\r\n"), nil
}

func (server *GopherServer) createLink(itemType, text, path string) string {
	return fmt.Sprintf("%s%s\t%s\t%s\t%s\r\n", itemTypes[itemType], text, path, server.Domain, server.Port)
}

func (server *GopherServer) parseTemplate(templ string, data map[string]string) (string, error) {

	outHeader := bytes.NewBufferString("")
	outFile := bytes.NewBufferString("")
	outFooter := bytes.NewBufferString("")

	// Parse header template
	if headerTmpl, err := template.New("header").Parse(server.ServerRoot.HeaderTemplate()); err == nil {
		if err := headerTmpl.Execute(outHeader, data); err != nil {
			return "", errors.New("Could not parse header Template")
		}
	}

	// Parse footer template
	if footerTmpl, err := template.New("footer").Parse(server.ServerRoot.FooterTemplate()); err == nil {
		if err := footerTmpl.Execute(outFooter, data); err != nil {
			return "", errors.New("Could not parse footer template")
		}
	}

	// Parse request file as template
	if fileTmpl, err := template.New("request").Parse(templ); err == nil {
		if err := fileTmpl.Execute(outFile, data); err != nil {
			return "", errors.New("Could not parse file template")
		}
	}

	return strings.TrimRight(outHeader.String()+outFile.String()+outFooter.String(), "\n"), nil
}

func (server *GopherServer) isTemplate(path string) bool {
	//TODO implement
	return true
}
