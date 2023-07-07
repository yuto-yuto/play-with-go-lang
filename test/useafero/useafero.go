package useafero

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/afero"
)

type FileHandler struct {
	FileSystem afero.Fs
}

func (f *FileHandler) Create(path string, content string) error {
	exist, err := afero.Exists(f.FileSystem, path)
	if err != nil {
		return fmt.Errorf("failed to check path existence: %w", err)
	}

	if exist {
		if err = f.FileSystem.Rename(path, path+"_backup"); err != nil {
			return fmt.Errorf("failed to rename a file: %w", err)
		}
	}

	file, err := f.FileSystem.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create a file: %w", err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write content: %w", err)
	}

	return nil
}

func (f *FileHandler) ReadToGetSum(path string) (int, error) {
	exist, err := afero.Exists(f.FileSystem, path)
	if err != nil {
		return 0, fmt.Errorf("failed to check path existence: %w", err)
	}

	if !exist {
		return 0, os.ErrNotExist
	}

	file, err := f.FileSystem.Open(path)
	if err != nil {
		return 0, fmt.Errorf("failed to open a file: %w", err)
	}

	buffer := make([]byte, 10)
	content := ""

	for {
		size, err := file.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return 0, fmt.Errorf("failed to read content: %w", err)
		}

		content += string(buffer[0:size])
	}

	lines := strings.Split(content, "\n")
	sum := 0
	for _, value := range lines {
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("unexpected format: %w", err)
		}
		sum += int(intValue)
	}

	return sum, nil
}
