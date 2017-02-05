package main

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
)

type Picture struct {
	Name string
	URL  string
}

func storeFile(userId, fileName string, f multipart.File) error {
	os.MkdirAll(STORAGE+userId, os.ModePerm)
	outF, e := os.OpenFile(filepath.Join(STORAGE, userId, fileName), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer outF.Close()
	if e != nil {
		return e
	}
	if _, e = io.Copy(outF, f); e != nil {
		return e
	}
	return nil
}

func listFiles(path string) []Picture {
	files, err := ioutil.ReadDir("./storage/" + path)
	if err != nil {
		return make([]Picture, 0)
	}

	paths := make([]Picture, len(files))
	for idx, file := range files {
		paths[idx].Name = file.Name()
		paths[idx].URL = filepath.Join("/image", path, file.Name())
	}
	return paths
}
