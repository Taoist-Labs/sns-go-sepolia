package sns

import (
	api "github.com/Taoist-Labs/sns-go-api"
	namehash "github.com/Taoist-Labs/sns-go-namehash"
	safe "github.com/Taoist-Labs/sns-go-safe"
)

// Resolve sns to address
// parameter 'sns' example: 'abc.seedao' 'sub.abc.seedao'
func Resolve(sns string) string {
	return ResolveWithRPC(sns, api.RPC)
}

func ResolveWithRPC(sns, rpc string) string {
	if len(sns) == 0 {
		return "0x0000000000000000000000000000000000000000" // sns is empty
	}

	ok, name := namehash.Normalize(sns)
	if !ok {
		return "0x0000000000000000000000000000000000000000" // sns is empty
	}

	if !safe.IsSafe(name) {
		return "0x0000000000000000000000000000000000000000" // sns is empty
	}

	return api.ResolveWithRPC(name, rpc)
}

func Resolves(sns []string) []string {
	return ResolvesWithRPC(sns, api.RPC)
}

func ResolvesWithRPC(sns []string, rpc string) []string {
	if len(sns) == 0 {
		return []string{}
	}

	var names []string
	for _, s := range sns {
		ok, n := namehash.Normalize(s)
		if ok && safe.IsSafe(n) {
			names = append(names, n)
		} else {
			names = append(names, "")
		}
	}

	return api.Resolves(names)
}

// Name address to sns
// return addr example: 'abc.seedao' 'sub.abc.seedao'
func Name(addr string) (sns string) {
	return NameWithRPC(addr, api.RPC)
}

func NameWithRPC(addr, rpc string) (sns string) {
	if len(addr) == 0 {
		return "" // address is empty
	}

	name := api.NameWithRPC(addr, rpc)
	if len(name) == 0 {
		return "" // address is empty
	}

	if !safe.IsSafe(name) {
		return "" // address is empty
	}

	return name
}

func Names(addr []string) []string {
	return NamesWithRPC(addr, api.RPC)
}

func NamesWithRPC(addr []string, rpc string) []string {
	if len(addr) == 0 {
		return []string{}
	}

	sns := api.Names(addr)

	return safe.Safe(sns)
}
