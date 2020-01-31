package main

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// GopherServerRoot abstracts the filesystem and limits access to the server's
// root folder
type GopherServerRoot struct {
	ServerRootDir string
	TemplatesDir  string
}

// NewGopherServerRoot returns a new server root object
func NewGopherServerRoot(root, templates string) (*GopherServerRoot, error) {

	// Check if root path exists
	info, err := os.Stat(root)
	if os.IsNotExist(err) {
		return nil, err
	}

	// Check if it is a directory
	if !info.IsDir() {
		return nil, err
	}

	// Check if templates path exists
	info, err = os.Stat(templates)
	if os.IsNotExist(err) {
		return nil, err
	}

	// Check if it is a directory
	if !info.IsDir() {
		return nil, err
	}

	return &GopherServerRoot{ServerRootDir: root, TemplatesDir: templates}, nil
}

// FileExists checks if a file exists relative to the server root
func (sr *GopherServerRoot) FileExists(path string) bool {

	savePath, err := sr.getSavePath(path)

	// If the path is not save, return as if it does not exist
	if err != nil {
		return false
	}

	info, err := os.Stat(savePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()

}

// DirExists checks if a directory exists relative to the server root
func (sr *GopherServerRoot) DirExists(path string) bool {

	savePath, err := sr.getSavePath(path)

	// If the path is not save, return as if it does not exist
	if err != nil {
		return false
	}

	info, err := os.Stat(savePath)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()

}

// getSavePath returns a path that is inside the server root to prevent
// directry traversal. It does *not* check if the requested file or directory
// actually exist!
func (sr *GopherServerRoot) getSavePath(subPath string) (string, error) {
	savePath := path.Clean(path.Join(sr.ServerRootDir, subPath))
	if !strings.HasPrefix(savePath, sr.ServerRootDir) {
		return "", errors.New("Path outside of server root: " + savePath)
	}
	return savePath, nil
}

// Type returns the gopher-specific type for a given item relative to the server root
func (sr *GopherServerRoot) Type(reqPath string) string {

	savePath, err := sr.getSavePath(reqPath)
	log.Info("Determining type for: ", savePath)

	// If the path is not save, return as if it does not exist
	if err != nil {
		log.Errorf("Error reading path: %s (%s)", reqPath, err)
		return ""
	}

	info, err := os.Stat(savePath)
	if os.IsNotExist(err) {
		log.Errorf("Error reading path: %s (%s)", reqPath, err)
		return ""
	}

	// Check if directory exists and does not contain gophermap
	if info.IsDir() && !sr.FileExists(path.Join(savePath, "index.gph")) {
		return "MENU"
	}

	return "TEXT"
}
func (sr *GopherServerRoot) HeaderTemplate() string {

	path := path.Join(sr.TemplatesDir, "header.gph")
	file, err := ioutil.ReadFile(path)

	if err != nil {
		log.Errorf("Error reading template: %s (%s)", path, err)
		return ""
	}
	return string(file)
}

func (sr *GopherServerRoot) FooterTemplate() string {

	path := path.Join(sr.TemplatesDir, "footer.gph")
	file, err := ioutil.ReadFile(path)

	if err != nil {
		log.Errorf("Error reading template: %s (%s)", path, err)
		return ""
	}
	return string(file)
}

// GetServerFile returns the contents of the requested file relative to the server root
func (sr *GopherServerRoot) GetServerFile(subpath string) ([]byte, error) {

	log.Info("Opening: ", subpath)

	path, err := sr.getSavePath(subpath)
	if err != nil {
		log.Errorf("Error reading file: %s (%s)", path, err)
		return nil, err
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Errorf("Error reading file: %s (%s)", path, err)
		return nil, err
	}
	return file, nil

}

// GetServerDir returns the names of all items in a directory relative to the server root
func (sr *GopherServerRoot) GetServerDir(subpath string) ([]string, error) {

	path, err := sr.getSavePath(subpath)
	if err != nil {
		log.Errorf("Error reading directory: %s (%s)", path, err)
		return nil, err
	}

	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Errorf("Error reading directory: %s (%s)", path, err)
		return nil, err
	}

	var list []string

	for _, f := range files {
		list = append(list, f.Name())
	}

	return list, nil
}
