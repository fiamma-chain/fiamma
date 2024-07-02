package bitvm

/*
#cgo linux LDFLAGS: ${SRCDIR}/lib/libbitvm.so -ldl -lrt -lm -lssl -lcrypto -Wl,--allow-multiple-definition
#cgo darwin LDFLAGS: -L./lib -lbitvm

#include "lib/bitvm.h"
*/
import "C"
import "unsafe"

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
