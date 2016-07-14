// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lzfse

/*
#cgo CFLAGS: -I. -I./internal/lzfse/src
#cgo !windows LDFLAGS: -lm

#include "./capi.h"

#include <stdlib.h>
*/
import "C"
import "unsafe"

func Main(args []string) int {
	argc := C.int(len(args))
	argv := make([]*C.char, len(args))
	for i := 0; i < len(args); i++ {
		argv[i] = C.CString(args[i])
		defer C.free(unsafe.Pointer(argv[i]))
	}
	rv := C.lzfse_main(argc, (**C.char)(unsafe.Pointer(&argv[0])))
	return int(rv)
}
