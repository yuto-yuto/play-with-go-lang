package selfmock_test

import (
	"errors"
	"play-with-go-lang/test/selfmock"
)

type Spy struct {
	CallCount map[string]int
	Args      map[string][][]any
}

func (s *Spy) Init() {
	s.CallCount = make(map[string]int)
	s.Args = make(map[string][][]any)
}

func (s *Spy) Register(funcName string, args ...any) {
	val := s.CallCount[funcName]
	val++

	s.CallCount[funcName] = val
	s.Args[funcName] = append(s.Args[funcName], args)
}

type FileReaderMock struct {
	selfmock.Reader
	spy       Spy
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
	instance.spy.Init()

	instance.FakeOpen = func() error { return nil }
	instance.FakeClose = func() {}
	instance.FakeRead = func(size int) (string, error) {
		return "", errors.New("define the behavior")
	}

	return instance
}

func (f FileReaderMock) Open() error {
	f.spy.Register(openKey)
	return f.FakeOpen()
}

func (f FileReaderMock) Close() {
	f.spy.Register(closeKey)
	f.FakeClose()
}

func (f FileReaderMock) Read(size int) (string, error) {
	f.spy.Register(readKey, size)
	return f.FakeRead(size)
}
