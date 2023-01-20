package utils

import (
	"archive/zip"
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"
)

func CreateFile() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("unable to get the current filename")
		return
	}
	dirname := filepath.Dir(filename)
	fmt.Printf("Current dir: %s\n", dirname)

	buf := new(bytes.Buffer)
	writer := zip.NewWriter(buf)
	filePath := filepath.Join(dirname, "file.txt")
	// filePath := filepath.Join(dirname, "temp", "file.txt")
	fmt.Printf("File path : %s\n", filePath)

	file, err := writer.Create(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if _, err := file.Write([]byte("something")); err != nil {
		fmt.Println(err.Error())
		return
	}

	if err = writer.Flush(); err != nil {
		fmt.Println(err.Error())
		return
	}

	writer.Close()
}
