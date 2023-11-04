package sns

import namehash "github.com/Taoist-Labs/sns-go-namehash"
import safe "github.com/Taoist-Labs/sns-go-safe"
import api "github.com/Taoist-Labs/sns-go-api"

func Resolve(sns string) (addr string) {
	return ResolveWithRPC(sns, api.RPC)
}

func ResolveWithRPC(sns, rpc string) (addr string) {
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
