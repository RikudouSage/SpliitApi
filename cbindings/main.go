package main

import "C"

const (
	// SpliitSuccess indicates that a C API call completed successfully.
	SpliitSuccess C.int = iota
	// SpliitError indicates that a C API call failed and last error is available.
	SpliitError
)

func main() {}
