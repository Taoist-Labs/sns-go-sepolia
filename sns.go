package sns

import (
	api "github.com/Taoist-Labs/sns-go-api"
	namehash "github.com/Taoist-Labs/sns-go-namehash"
	safe "github.com/Taoist-Labs/sns-go-safe"
)

// Resolve sns to address
// parameter 'sns' example: 'abc.seedao' 'sub.abc.seedao'
func Resolve(sns string) string {
	return resolve(sns, safeHost, indexerHost, rpc, publicResolverAddr)
}

// ResolveWithRPC sns to address with custom rpc
// parameter 'sns' example: 'abc.seedao' 'sub.abc.seedao'
func ResolveWithRPC(sns, rpc string) string {
	return resolve(sns, safeHost, indexerHost, rpc, publicResolverAddr)
}

func resolve(sns, safeHost, indexerHost, rpc, publicResolverAddr string) (addr string) {
	if len(sns) == 0 {
		return "0x0000000000000000000000000000000000000000" // sns is empty
	}

	ok, name := namehash.Normalize(sns)
	if !ok {
		return "0x0000000000000000000000000000000000000000" // sns is empty
	}

	if !safe.IsSafe(name, safeHost) {
		return "0x0000000000000000000000000000000000000000" // sns is empty
	}

	return api.Resolve(name, indexerHost, rpc, publicResolverAddr)
}

// Resolves sns to address batch
// parameter 'sns' example: ['abc.seedao', 'sub.abc.seedao']
func Resolves(sns []string) []string {
	return resolves(sns, safeHost, indexerHost, rpc, publicResolverAddr)
}

// ResolvesWithRPC sns to address batch with custom rpc
// parameter 'sns' example: ['abc.seedao', 'sub.abc.seedao']
func ResolvesWithRPC(sns []string, rpc string) []string {
	return resolves(sns, safeHost, indexerHost, rpc, publicResolverAddr)
}

func resolves(sns []string, safeHost, indexerHost, rpc, publicResolverAddr string) []string {
	if len(sns) == 0 {
		return []string{}
	}

	var names []string
	for _, s := range sns {
		ok, n := namehash.Normalize(s)
		if ok && safe.IsSafe(n, safeHost) {
			names = append(names, n)
		} else {
			names = append(names, "")
		}
	}

	return api.Resolves(names, indexerHost, rpc, publicResolverAddr)
}

// Name parse address from sns
// return addr example: '0x8C913aEc7443FE2018639133398955e0E17FB0C1' '0xc1eE7cB74583D1509362467443C44f1FCa981283'
func Name(addr string) (sns string) {
	return name(addr, safeHost, indexerHost, rpc, publicResolverAddr)
}

// NameWithRPC parse address from sns with custom rpc
// return addr example: '0x8C913aEc7443FE2018639133398955e0E17FB0C1' '0xc1eE7cB74583D1509362467443C44f1FCa981283'
func NameWithRPC(addr, rpc string) (sns string) {
	return name(addr, safeHost, indexerHost, rpc, publicResolverAddr)
}

func name(addr, safeHost, indexerHost, rpc, publicResolverAddr string) (sns string) {
	if len(addr) == 0 {
		return "" // address is empty
	}

	name := api.Name(addr, indexerHost, rpc, publicResolverAddr)
	if len(name) == 0 {
		return "" // address is empty
	}

	if !safe.IsSafe(name, safeHost) {
		return "" // address is empty
	}

	return name
}

// Names parse address from sns batch
// return addr example: ['0x8C913aEc7443FE2018639133398955e0E17FB0C1', '0xc1eE7cB74583D1509362467443C44f1FCa981283']
func Names(addr []string) []string {
	return names(addr, safeHost, indexerHost, rpc, publicResolverAddr)
}

// NamesWithRPC parse address from sns batch with custom rpc
// return addr example: ['0x8C913aEc7443FE2018639133398955e0E17FB0C1', '0xc1eE7cB74583D1509362467443C44f1FCa981283']
func NamesWithRPC(addr []string, rpc string) []string {
	return names(addr, safeHost, indexerHost, rpc, publicResolverAddr)
}

func names(addr []string, safeHost, indexerHost, rpc, publicResolverAddr string) []string {
	if len(addr) == 0 {
		return []string{}
	}

	sns := api.Names(addr, indexerHost, rpc, publicResolverAddr)

	return safe.Safe(sns, safeHost)
}

// TokenId get sns's ERC721 token id
// parameter 'sns' example: 'abc.seedao' 'sub.abc.seedao'
func TokenId(sns string) string {
	return tokenId(sns, safeHost, indexerHost, rpc, baseRegistrarAddr)
}

// TokenIdWithRPC get sns's ERC721 token id with custom rpc
// parameter 'sns' example: 'abc.seedao' 'sub.abc.seedao'
func TokenIdWithRPC(sns, rpc string) string {
	return tokenId(sns, safeHost, indexerHost, rpc, baseRegistrarAddr)
}

func tokenId(sns, safeHost, indexerHost, rpc, baseRegistrarAddr string) (addr string) {
	if len(sns) == 0 {
		return "" // sns is empty
	}

	ok, name := namehash.Normalize(sns)
	if !ok {
		return "" // sns is empty
	}

	if !safe.IsSafe(name, safeHost) {
		return "" // sns is not safe
	}

	return api.TokenId(name, indexerHost, rpc, baseRegistrarAddr)
}
