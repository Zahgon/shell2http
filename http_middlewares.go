package main

import (
	"net/http"
)

// mwMultiMethod - produce handler for several http methods
func mwMultiMethod(in map[string]http.HandlerFunc) (http.HandlerFunc, error) {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc), nil
}

// not matched http method

// mwMethodOnly - allow one HTTP method only
func mwMethodOnly(handler http.HandlerFunc, method string) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// mwBasicAuth - add HTTP Basic Authentication
func mwBasicAuth(handler http.HandlerFunc, users authUsers) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// mwLogging - add logging for handler
func mwLogging(handler http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// mwCommonHeaders - set common headers
func mwCommonHeaders(handler http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// mwOneThread - run handler in one thread
func mwOneThread(handler http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// responseWriterLogger - wrapper around http.ResponseWriter
type responseWriterLogger struct {
	srcRW      http.ResponseWriter
	statusCode int
	size       int
}

func (rwl *responseWriterLogger) Header() http.Header {
	_ = "STUB: not implemented"
	return *new(http.Header)
}

func (rwl *responseWriterLogger) Write(data []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (rwl *responseWriterLogger) WriteHeader(statusCode int) { _ = "STUB: not implemented"; return }

func (rwl *responseWriterLogger) StatusCode() int { _ = "STUB: not implemented"; return 0 }

func (rwl *responseWriterLogger) Size() int { _ = "STUB: not implemented"; return 0 }
