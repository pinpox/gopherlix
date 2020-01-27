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

func getSavePath(subPath string) (string, error) {
	//TODO check for directory traversal
	savePath := path.Clean(path.Join(SERVER_ROOT, subPath))
	if !strings.HasPrefix(savePath, SERVER_ROOT) {
		return "", errors.New("Path outside of server root")
	}
	return savePath, nil
}
