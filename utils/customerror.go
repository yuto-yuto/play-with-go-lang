package utils

import (
	"errors"
	"fmt"
)

var ErrPredefinedError = errors.New("predefined error message")

type ErrorWithoutPrep struct{}

func (e *ErrorWithoutPrep) Error() string {
	return "error without prep"
}

type ErrorWithPrep struct {
	Name string
}

func (e *ErrorWithPrep) Error() string {
	return fmt.Sprintf("error with prep, name: %s", e.Name)
}

type ErrorIsImpl struct {
	Name string
	ID   int
}

func (e *ErrorIsImpl) Error() string {
	return fmt.Sprintf("error is impl, name: %s", e.Name)
}

func (e *ErrorIsImpl) Is(err error) bool {
	other, ok := err.(*ErrorIsImpl)
	if !ok {
		return false
	}

	return other.Name == e.Name && other.ID == e.ID
}

type WrappedError struct {
	OriginalError error
}

func (e *WrappedError) Error() string {
	return fmt.Sprintf("wrapped error, original: %w", e.OriginalError)
}

func (e *WrappedError) Unwrap() error {
	return e.OriginalError
}

func runIs(original error, expected error) {
	result1 := errors.Is(original, expected)
	wrapped := fmt.Errorf("wrapped error: %w", original)
	result2 := errors.Is(wrapped, expected)
	fmt.Printf("%v\t %v\n", result1, result2)
}

func RunCustomError() {
	fmt.Println("------- errors.Unwrap ------")
	wrapped := fmt.Errorf("wrapped error: %w", ErrPredefinedError)
	fmt.Println(ErrPredefinedError)     // predefined error message
	fmt.Println(wrapped)                // wrapped error: predefined error message
	fmt.Println(errors.Unwrap(wrapped)) // predefined error message

	fmt.Println("------- errors.Is ------")

	runIs(ErrPredefinedError, ErrPredefinedError)                               // true     true
	runIs(&ErrorWithoutPrep{}, &ErrorWithoutPrep{})                             // true     true
	runIs(&ErrorWithPrep{}, &ErrorWithPrep{})                                   // false    false
	runIs(&ErrorWithPrep{Name: "Yuto"}, &ErrorWithPrep{Name: "Yuto"})           // false    false
	runIs(&ErrorIsImpl{Name: "Yuto"}, &ErrorIsImpl{Name: "Yuto", ID: 2})        // false     false
	runIs(&ErrorIsImpl{Name: "Yuto", ID: 1}, &ErrorIsImpl{Name: "Yuto", ID: 1}) // true     true
	runIs(&ErrorIsImpl{Name: "Yuto"}, &ErrorIsImpl{Name: "NN"})                 // false    false

	fmt.Println("------- errors.As compare same type------")
	var withoutPrep *ErrorWithoutPrep
	wrapped = fmt.Errorf("wrapped error: %w", &ErrorWithoutPrep{})
	result1 := errors.As(&ErrorWithoutPrep{}, &withoutPrep)
	result2 := errors.As(wrapped, &withoutPrep)
	fmt.Printf("%v\t%v\n", result1, result2) // true		true

	var withPrep *ErrorWithPrep
	wrapped = fmt.Errorf("wrapped error: %w", &ErrorWithPrep{Name: "Yuto"})
	result1 = errors.As(&ErrorWithPrep{Name: "Yuto"}, &withPrep)
	result2 = errors.As(wrapped, &withPrep)
	fmt.Printf("%v\t%v\n", result1, result2) // true		true

	var isImpl *ErrorIsImpl
	wrapped = fmt.Errorf("wrapped error: %w", &ErrorIsImpl{Name: "Yuto"})
	result1 = errors.As(&ErrorIsImpl{Name: "Yuto"}, &isImpl)
	result2 = errors.As(wrapped, &isImpl)
	fmt.Printf("%v\t%v\n", result1, result2) // true		true

	fmt.Println("------- errors.As compare different types------")

	fmt.Println(errors.As(&ErrorIsImpl{Name: "Yuto"}, &withPrep)) // false
	original := ErrorWithPrep{Name: "Yuto"}
	wrappedError := &WrappedError{OriginalError: &original}
	fmt.Println(errors.As(wrappedError, &withPrep)) // true
	fmt.Println(errors.Is(wrappedError, withPrep))  // true

}
