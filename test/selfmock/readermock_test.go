package selfmock_test

import (
	"errors"
	"play-with-go-lang/test/selfmock"
)

type FileReaderMock struct {
	selfmock.Reader
	callCount map[string]int
	FakeOpen  func() error
	FakeClose func()
	FakeRead  func(size int) (string, error)
}

const (
	openKey  = "Open"
	closeKey = "Close"
	readKey  = "Read"
)

func NewFileReaderMock() *FileReaderMock {
	instance := new(FileReaderMock)
	instance.callCount = make(map[string]int)

	instance.FakeOpen = func() error { return nil }
	instance.FakeClose = func() {}
	instance.FakeRead = func(size int) (string, error) {
		return "", errors.New("define the behavior")
	}

	return instance
}

func (f *FileReaderMock) incrementCallCount(key string) {
	val, prs := f.callCount[key]
	if prs {
		f.callCount[key] = val + 1
	} else {
		f.callCount[key] = 1
	}
}

func (f FileReaderMock) Open() error {
	f.incrementCallCount(openKey)
	return f.FakeOpen()
}

func (f FileReaderMock) Close() {
	f.incrementCallCount(closeKey)
	f.FakeClose()
}

func (f FileReaderMock) Read(size int) (string, error) {
	f.incrementCallCount(readKey)
	return f.FakeRead(size)
}
