package selfmock

import "fmt"

type Reader interface {
	Open() error
	Close()
	Read(size int) (string, error)
	Something() error
}

type FileReader struct {
	Path string
}

func NewFileReader(path string) *FileReader {
	instance := new(FileReader)
	instance.Path = path

	return instance
}

func (f *FileReader) Open() error {
	fmt.Printf("path: %s", f.Path)
	return nil
}

func (f *FileReader) Close() {
}
func (f *FileReader) Read(size int) (string, error) {
	return "abcde", nil
}
func (f *FileReader) Something() error {
	fmt.Println("Do something")
	return nil
}
