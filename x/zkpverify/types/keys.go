package types

const (
	// ModuleName defines the module name
	ModuleName = "zkpverify"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_zkpverify"
)

var (
	ParamsKey     = []byte{0x01} // key prefix for the parameters
	VerifyDataKey = []byte{0x02} // key prefix for the verify data

)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
