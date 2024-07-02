package verifiers

import (
	"unsafe"
)

/*
#cgo linux,amd64  LDFLAGS: ${SRCDIR}/lib/libverifier_linux_amd64.a -ldl -lrt -lm -lssl -lcrypto -Wl,--allow-multiple-definition
#cgo darwin,amd64  LDFLAGS: -L./lib -lverifier_darwin_amd64 -framework SystemConfiguration
#cgo darwin,arm64  LDFLAGS: -L./lib -lverifier_darwin_arm64 -framework SystemConfiguration

#include "lib/lib.h"
*/
import "C"

func VerifySp1Proof(proofBuffer []byte, proofLen uint32, elfBuffer []byte, elfLen uint32) bool {
	proofPtr := (*C.uchar)(unsafe.Pointer(&proofBuffer[0]))
	elfPtr := (*C.uchar)(unsafe.Pointer(&elfBuffer[0]))

	return (bool)(C.verify_sp1_proof_ffi(proofPtr, (C.uint32_t)(proofLen), elfPtr, (C.uint32_t)(elfLen)))
}

func VerifyBitvmProof(vkBuffer []byte, vkLen uint32, proofBuffer []byte, proofLen uint32, publicInputBuffer []byte, publicInputLen uint32) (bool, []byte) {
	vkPtr := (*C.uchar)(unsafe.Pointer(&vkBuffer[0]))
	proofPtr := (*C.uchar)(unsafe.Pointer(&proofBuffer[0]))
	piPtr := (*C.uchar)(unsafe.Pointer(&publicInputBuffer[0]))

	res := C.verify_proof_groth16_ffi(vkPtr, (C.uint32_t)(vkLen), proofPtr, (C.uint32_t)(proofLen), piPtr, (C.uint32_t)(publicInputLen))
	success := bool(res.success)
	length := int(res.len)
	data := C.GoBytes(unsafe.Pointer(res.data), C.int(length))
	C.free_verify_witness(res)

	return success, data
}
