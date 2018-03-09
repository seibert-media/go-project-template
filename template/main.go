package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/alecthomas/template"
)

var (
	targetDir = flag.String("targetDir", "", "absolute path where the result should be put")
)

func main() {
	err := filepath.Walk("./template/files", walkFunc)
	if err != nil {
		panic(err)
	}
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	templ := template.New(filepath.Base(path)).Funcs(template.FuncMap{
		"env": env,
	})
	_, err = templ.ParseFiles(path)
	if err != nil {
		return err
	}
	regex := regexp.MustCompile(".*template/files")
	path = regex.ReplaceAllString(path, *targetDir)
	//path = strings.Replace(path, "template/files/", "", -1)
	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	if err := templ.Execute(f, nil); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return err
}

func env(val interface{}) (interface{}, error) {
	if val == nil {
		return "", nil
	}
	value, ok := os.LookupEnv(val.(string))
	if !ok {
		return nil, fmt.Errorf("var not defined", val)
	}
	return value, nil
}
