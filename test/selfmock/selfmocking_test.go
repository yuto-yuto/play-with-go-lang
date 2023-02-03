package selfmock_test

import (
	"errors"
	"play-with-go-lang/test/selfmock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Provider", func() {
	var readerMock *FileReaderMock

	BeforeEach(func() {
		readerMock = NewFileReaderMock()
	})

	Describe("ProvideData", func() {
		It("should return data", func() {
			instance := selfmock.NewProvider(readerMock)
			readerMock.FakeRead = func(size int) (string, error) {
				return "inject fake data", nil
			}

			data, err := instance.ProvideData()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal("inject fake data"))
		})

		It("should return data (call the function twice)", func() {
			instance := selfmock.NewProvider(readerMock)
			readerMock.FakeRead = func(size int) (string, error) {
				return "inject fake data", nil
			}

			instance.ProvideData()
			instance.ProvideData()
			Expect(readerMock.spy.CallCount[readKey]).Should(Equal(2))
			Expect(readerMock.spy.Args[readKey][0][0]).Should(Equal(99))
			Expect(readerMock.spy.Args[readKey][1][0]).Should(Equal(99))
		})

		It("should call Close method when data is read", func() {
			instance := selfmock.NewProvider(readerMock)
			readerMock.FakeRead = func(size int) (string, error) {
				return "inject fake data", nil
			}

			_, err := instance.ProvideData()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(readerMock.spy.CallCount[closeKey]).Should(Equal(1))
		})

		It("should throw an error when Open method returns error", func() {
			instance := selfmock.NewProvider(readerMock)
			readerMock.FakeOpen = func() error {
				return errors.New("fake error")
			}

			_, err := instance.ProvideData()
			Expect(err).Should(HaveOccurred())
		})

		It("should not call Close method when Open method returns error", func() {
			instance := selfmock.NewProvider(readerMock)
			readerMock.FakeOpen = func() error {
				return errors.New("fake error")
			}

			instance.ProvideData()
			_, prs := readerMock.spy.CallCount[closeKey]
			Expect(prs).Should(BeFalse())
		})

		It("should return error when Read returns error", func() {
			instance := selfmock.NewProvider(readerMock)
			readerMock.FakeRead = func(size int) (string, error) {
				return "", errors.New("fake error")
			}

			_, err := instance.ProvideData()
			Expect(err).Should(HaveOccurred())
		})
	})
})
