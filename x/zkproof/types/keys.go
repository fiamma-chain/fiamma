package types

const (
	// ModuleName defines the module name
	ModuleName = "zkproof"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_zkproof"
)

var (
	ParamsKey = []byte("p_zkproof")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
