package cairo_platinum

/*
#cgo darwin LDFLAGS: ${SRCDIR}/lib/libcairo_platinum.a
#cgo linux LDFLAGS: ${SRCDIR}/lib/libcairo_platinum.a -ldl -lrt -lm -Wl,--allow-multiple-definition


#include "lib/cairo_platinum.h"
*/
import "C"

import (
	"unsafe"
)

const MAX_PROOF_SIZE = 1024 * 1024

func VerifyCairoProof100Bits(proofBuffer [MAX_PROOF_SIZE]byte, proofLen uint) bool {
	proofPtr := (*C.uchar)(unsafe.Pointer(&proofBuffer[0]))
	return (bool)(C.verify_cairo_proof_ffi_100_bits(proofPtr, (C.uint)(proofLen)))
}
