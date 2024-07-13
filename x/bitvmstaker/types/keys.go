package types

const (
	// ModuleName defines the module name
	ModuleName = "bitvmstaker"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_bitvmstaker"
)

var (
	ParamsKey = []byte("p_bitvmstaker")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
