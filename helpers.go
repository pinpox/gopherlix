package main

import (
	"errors"
	"os"
	"path"
	"strings"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func dirExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func (server *GopherServer) getSavePath(subPath string) (string, error) {
	savePath := path.Clean(path.Join(server.RootDir, subPath))
	if !strings.HasPrefix(savePath, server.RootDir) {
		return "", errors.New("Path outside of server root: " + savePath)
	}
	return savePath, nil
}
