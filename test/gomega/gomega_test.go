package gomega_test

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

var errFoo error = errors.New("foo error")

type ErrorWithoutProp struct{}

func (e ErrorWithoutProp) Error() string {
	return "error message from ErrorWithoutProp"
}

type ErrorWithProp struct {
	Name string
}

func (e ErrorWithProp) Error() string {
	return "error message from ErrorWithProp"
}

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "gomega matcher test suite")
}

var _ = Describe("Provider", func() {
	Describe("error validation", func() {
		It("predefined error", func() {
			Expect(errFoo).Should(HaveOccurred())
			Expect(errFoo.Error()).Should(ContainSubstring("foo err"))

			Expect(errors.Is(errFoo, errFoo)).Should(BeTrue())
			Expect(errFoo).Should(MatchError(errFoo))

			wrappedErr := fmt.Errorf("additional error info here: %w", errFoo)
			Expect(errors.Is(wrappedErr, errFoo)).Should(BeTrue())
		})

		It("without prep", func() {
			err := &ErrorWithoutProp{}
			Expect(errors.Is(err, &ErrorWithoutProp{})).Should(BeTrue())

			var expectedErr *ErrorWithoutProp
			Expect(errors.As(err, &expectedErr)).Should(BeTrue())

			Expect(err).Should(MatchError(&ErrorWithoutProp{}))

			wrappedErr := fmt.Errorf("additional error info here: %w", err)
			Expect(errors.Is(wrappedErr, &ErrorWithoutProp{})).Should(BeTrue())
			Expect(errors.As(wrappedErr, &expectedErr)).Should(BeTrue())
		})

		It("with prep", func() {
			err := &ErrorWithProp{Name: "Yuto"}
			Expect(errors.Is(err, &ErrorWithProp{Name: "Yuto"})).Should(BeFalse())

			var expectedErr *ErrorWithProp
			Expect(errors.As(err, &expectedErr)).Should(BeTrue())

			Expect(err).Should(MatchError(&ErrorWithProp{Name: "Yuto"}))

			wrappedErr := fmt.Errorf("additional error info here: %w", err)
			Expect(errors.As(wrappedErr, &expectedErr)).Should(BeTrue())

			Expect(wrappedErr).ShouldNot(MatchError(&ErrorWithProp{Name: "Yuto"}))
		})
	})

	Describe("validation", func() {
		It("number", func() {
			Expect(int64(1)).ShouldNot(Equal(1))
			Expect(int64(1)).Should(Equal(int64(1)))

			floatingValue := 1.0
			Expect(floatingValue).ShouldNot(Equal(1))
			Expect(floatingValue).Should(Equal(1.0))
			Expect(floatingValue).Should(Equal(float64(1)))
			Expect(floatingValue).Should(BeNumerically("==", 1))

			Expect(5).Should(BeNumerically(">", 2))
			Expect(5).Should(BeNumerically(">=", 5))
			Expect(5).Should(BeNumerically("<", 50))

			Expect(1.001).Should(BeNumerically("~", 1.0, 0.01))
			Expect(1.001).ShouldNot(BeNumerically("~", 1.0, 0.0001))

		})

		It("bool", func() {
			Expect(true).Should(Equal(true))
			Expect(true).Should(BeTrue())
			Expect(false).Should(BeFalse())
		})

		It("nil", func() {
			var nilValue *string
			Expect(nilValue).ShouldNot(Equal(nil))
			Expect(nilValue).Should(BeNil())
		})

		It("array", func() {
			emptyArray := []int{}
			Expect(emptyArray).ShouldNot(BeNil())
			Expect(emptyArray).Should(BeEmpty())
			Expect(emptyArray).Should(HaveLen(0))
			Expect(len(emptyArray)).Should(BeZero())

			var nilArray []int
			Expect(len(nilArray)).Should(BeZero())
			Expect(nilArray).Should(BeEmpty())
			Expect(nilArray).Should(BeNil())
			Expect(nilArray).Should(HaveLen(0))

			array := []int{5, 4, 3, 2, 1}
			Expect(array).Should(HaveLen(5))
			Expect(array).Should(Equal([]int{5, 4, 3, 2, 1}))
			Expect(array).ShouldNot(Equal([]int{1, 2, 3, 4, 5}))
			Expect(array).Should(ContainElements([]int{4, 2, 5}))
			Expect(array).Should(ContainElements(1, 3, 2))
		})

		It("map", func() {
			emptyMap := map[string]int{}
			Expect(emptyMap).ShouldNot(BeNil())
			Expect(emptyMap).Should(BeEmpty())
			Expect(emptyMap).Should(HaveLen(0))
			Expect(len(emptyMap)).Should(BeZero())

			var nilMap map[string]int
			Expect(len(nilMap)).Should(BeZero())
			Expect(nilMap).Should(BeEmpty())
			Expect(nilMap).Should(BeNil())
			Expect(nilMap).Should(HaveLen(0))

			mapValue := map[string]int{"first": 1, "second": 2, "third": 3}
			Expect(mapValue).Should(HaveLen(3))
			Expect(mapValue).Should(Equal(map[string]int{"first": 1, "second": 2, "third": 3}))
			Expect(mapValue).Should(Equal(map[string]int{"second": 2, "first": 1, "third": 3}))
			Expect(mapValue).Should(HaveKey("second"))
			Expect(mapValue).Should(HaveKeyWithValue("second", 2))
		})
	})
})
