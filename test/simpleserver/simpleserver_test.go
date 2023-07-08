package simpleserver_test

import (
	// "errors"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "simple server test suite")
}

const (
	serverHost = "localhost:9999"
	dataPath   = "/api/v1/data&id=1"
	countPath  = "/api/v1/count"
)

var _ = Describe("Simple server test", func() {
	Describe("random port", func() {
		It("", func() {

		})
	})
})
