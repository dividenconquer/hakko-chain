package types

const (
	// ModuleName defines the module name
	ModuleName = "hakkochain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_hakkochain"
)

var (
	ParamsKey = []byte("p_hakkochain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
