package main

import (
	"flag"
	"nashrul-be/crm/app"
	"os"
	"path/filepath"
)

var (
	envPath string
)

func GetCurrentDir() (string, error) {
	program := os.Args[0]
	absPath, err := filepath.Abs(program)
	if err != nil {
		return "", err
	}
	return filepath.Dir(absPath), nil
}

func InitFlag() {
	cd, err := GetCurrentDir()
	if err != nil {
		panic("can't get current directory")
	}
	flag.StringVar(&envPath, "env", filepath.Join(cd, ".env"), "path to .env file")
}

func main() {
	InitFlag()
	flag.Parse()
	app.Init(envPath)
}
