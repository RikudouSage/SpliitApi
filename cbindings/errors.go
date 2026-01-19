package main

/*
#ifndef SPLIIT_ERRORS_CGO_PREAMBLE
#define SPLIIT_ERRORS_CGO_PREAMBLE 1

#include <stdlib.h>
#include <string.h>

#define SPLIIT_SUCCESS 0
#define SPLIIT_ERROR 1

static __thread char* spliit_last_error;

static void spliit_clear_last_error(void) {
	if (spliit_last_error) {
		free(spliit_last_error);
		spliit_last_error = NULL;
	}
}

static void spliit_set_last_error_copy(const char* msg) {
	spliit_clear_last_error();
	if (!msg) {
		return;
	}

	size_t len = strlen(msg);
	char* copy = (char*)malloc(len + 1);
	if (!copy) {
		return;
	}
	memcpy(copy, msg, len + 1);
	spliit_last_error = copy;
}

static size_t spliit_get_last_error(char* buf, size_t buf_len) {
	if (!spliit_last_error) {
		if (buf && buf_len > 0) {
			buf[0] = '\0';
		}
		return 1;
	}

	size_t len = strlen(spliit_last_error) + 1;
	if (buf && buf_len > 0) {
		size_t to_copy = len <= buf_len ? len : buf_len - 1;
		memcpy(buf, spliit_last_error, to_copy);
		buf[to_copy] = '\0';
	}
	return len;
}

#endif
*/
import "C"
import (
	"errors"
	"unsafe"
)

func clearLastError() {
	C.spliit_clear_last_error()
}

func setLastError(err error) {
	if err == nil {
		clearLastError()
		return
	}
	setLastErrorMessage(err.Error())
}

func setLastErrorMessage(msg string) {
	if msg == "" {
		clearLastError()
		return
	}
	cstr := C.CString(msg)
	C.spliit_set_last_error_copy(cstr)
	C.free(unsafe.Pointer(cstr))
}

func nullPointerError(name string) error {
	return errors.New(name + " is NULL")
}

//revive:disable-next-line:var-naming
//export Spliit_GetLastError
func Spliit_GetLastError(buf *C.char, bufLen C.size_t) C.size_t {
	return C.spliit_get_last_error(buf, bufLen)
}
