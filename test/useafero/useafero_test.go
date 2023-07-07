package useafero_test

import (
	"errors"
	"os"
	"play-with-go-lang/test/useafero"

	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "afero test suite")
}

var _ = Describe("afero test", func() {
	var handler useafero.FileHandler

	BeforeEach(func() {
		handler = useafero.FileHandler{
			FileSystem: afero.NewMemMapFs(),
		}
	})

	Describe("Create", func() {
		It("create a file", func() {
			err := handler.Create("/unknown11a/tmp/abc.txt", "a\nb\nc\n")
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("create a backup file if the file already exists", func() {
			exist, err := afero.Exists(handler.FileSystem, "/unknown11a/tmp/abc.txt_backup")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(exist).Should(BeFalse())

			err = handler.Create("/unknown11a/tmp/abc.txt", "a\nb\nc\n")
			Expect(err).ShouldNot(HaveOccurred())

			err = handler.Create("/unknown11a/tmp/abc.txt", "a\nb\nc\n")
			Expect(err).ShouldNot(HaveOccurred())

			exist, err = afero.Exists(handler.FileSystem, "/unknown11a/tmp/abc.txt_backup")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(exist).Should(BeTrue())
		})
	})

	Describe("ReadToGetSum", func() {
		It("error when a file doesn't exist", func() {
			_, err := handler.ReadToGetSum("/unknown11a/tmp/data.txt")
			Expect(errors.Is(err, os.ErrNotExist)).Should(BeTrue())
		})

		It("error when file format is unexpected", func() {
			file, err := handler.FileSystem.Create("/unknown11a/tmp/data.txt")
			Expect(err).ShouldNot(HaveOccurred())

			file.WriteString("1 1\n 2\n")

			_, err = handler.ReadToGetSum("/unknown11a/tmp/data.txt")
			Expect(err.Error()).Should(ContainSubstring("unexpected format"))
		})

		FIt("succeeds", func() {
			file, err := handler.FileSystem.Create("/unknown11a/tmp/data.txt")
			Expect(err).ShouldNot(HaveOccurred())

			file.WriteString("1\n2\n3")

			sum, err := handler.ReadToGetSum("/unknown11a/tmp/data.txt")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(sum).Should(Equal(6))
		})
	})
})
