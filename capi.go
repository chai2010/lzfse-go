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

func lzfseMain(args []string) int {
	argc := C.int(len(args))
	argv := make([]*C.char, len(args))
	for i := 0; i < len(args); i++ {
		argv[i] = C.CString(args[i])
		defer C.free(unsafe.Pointer(argv[i]))
	}
	rv := C.lzfse_main(argc, (**C.char)(unsafe.Pointer(&argv[0])))
	return int(rv)
}

/*! @abstract Get the required scratch buffer size to compress using LZFSE.   */
//size_t lzfse_encode_scratch_size() LZFSE_LIB_API;

/*! @abstract Compress a buffer using LZFSE.
 *
 *  @param dst_buffer
 *  Pointer to the first byte of the destination buffer.
 *
 *  @param dst_size
 *  Size of the destination buffer in bytes.
 *
 *  @param src_buffer
 *  Pointer to the first byte of the source buffer.
 *
 *  @param src_size
 *  Size of the source buffer in bytes.
 *
 *  @param scratch_buffer
 *  If non-NULL, a pointer to scratch space for the routine to use as workspace;
 *  the routine may use up to lzfse_encode_scratch_size( ) bytes of workspace
 *  during its operation, and will not perform any internal allocations. If
 *  NULL, the routine may allocate its own memory to use during operation via
 *  a single call to malloc( ), and will release it by calling free( ) prior
 *  to returning. For most use, passing NULL is perfectly satisfactory, but if
 *  you require strict control over allocation, you will want to pass an
 *  explicit scratch buffer.
 *
 *  @return
 *  The number of bytes written to the destination buffer if the input is
 *  successfully compressed. If the input cannot be compressed to fit into
 *  the provided buffer, or an error occurs, zero is returned, and the
 *  contents of dst_buffer are unspecified.                                   */
//size_t lzfse_encode_buffer(uint8_t *__restrict dst_buffer,
//                           size_t dst_size,
//                           const uint8_t *__restrict src_buffer,
//                           size_t src_size,
//                           void *__restrict scratch_buffer) LZFSE_LIB_API;

/*! @abstract Get the required scratch buffer size to decompress using LZFSE. */
//size_t lzfse_decode_scratch_size() LZFSE_LIB_API;

/*! @abstract Decompress a buffer using LZFSE.
 *
 *  @param dst_buffer
 *  Pointer to the first byte of the destination buffer.
 *
 *  @param dst_size
 *  Size of the destination buffer in bytes.
 *
 *  @param src_buffer
 *  Pointer to the first byte of the source buffer.
 *
 *  @param src_size
 *  Size of the source buffer in bytes.
 *
 *  @param scratch_buffer
 *  If non-NULL, a pointer to scratch space for the routine to use as workspace;
 *  the routine may use up to lzfse_decode_scratch_size( ) bytes of workspace
 *  during its operation, and will not perform any internal allocations. If
 *  NULL, the routine may allocate its own memory to use during operation via
 *  a single call to malloc( ), and will release it by calling free( ) prior
 *  to returning. For most use, passing NULL is perfectly satisfactory, but if
 *  you require strict control over allocation, you will want to pass an
 *  explicit scratch buffer.
 *
 *  @return
 *  The number of bytes written to the destination buffer if the input is
 *  successfully decompressed. If there is not enough space in the destination
 *  buffer to hold the entire expanded output, only the first dst_size bytes
 *  will be written to the buffer and dst_size is returned. Note that this
 *  behavior differs from that of lzfse_encode_buffer.                        */
//size_t lzfse_decode_buffer(uint8_t *__restrict dst_buffer,
//                           size_t dst_size,
//                           const uint8_t *__restrict src_buffer,
//                           size_t src_size,
//                           void *__restrict scratch_buffer) LZFSE_LIB_API;
