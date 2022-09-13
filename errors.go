package main

import "fmt"

type HttpError struct {
	Code int
	Message string
}

func (err *HttpError) Error() string {
	return fmt.Sprintf("HTTP Error with status code %d and with message %s", err.Code, err.Message)
}

func SomeFunc(code int) ([]string, *HttpError){
	if (code < 400) {
		return []string{"OK"}, nil
	}
	return []string{}, &HttpError{404, "Not Found"}
}