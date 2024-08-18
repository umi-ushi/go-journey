package main

import (
	"errors"
	"fmt"
)

func compare() {
	compareByIs()
	compareByAs()
	compareOfAggregatedErrors()

}

func compareByIs() {
	wrpErr := fmt.Errorf("wrap: %w", errNotFound)
	if errors.Is(wrpErr, errNotFound) {
		fmt.Println("wrapped error is errNotFound")
	}
}

func compareByAs() {
	err := &customError{Err: errors.New("original error")}
	var customErr *customError
	if errors.As(err, &customErr) {
		fmt.Println("Error is type customError")
	}
}

func compareOfAggregatedErrors() {
	aggErr := errors.Join(errNotFound, errNotFound)
	errs, err := discreteErrors(aggErr)
	if err != nil {
		fmt.Println("discreteErrors error:", err)
	}
	for _, err = range errs {
		fmt.Printf("errors match %s\n", errors.Is(err, errNotFound))
	}
}

var errNotFound = errors.New("not found")

type customError struct {
	Err error
}

func (e *customError) Error() string {
	return e.Err.Error()
}

func (e *customError) Unwrap() error {
	return e.Err
}

// 集約されたエラーをバラす
func discreteErrors(errs error) ([]error, error) {
	if errs, ok := errs.(interface{ Unwrap() []error }); ok {
		return errs.Unwrap(), nil
	}

	return nil, errors.New("not aggregated")
}
