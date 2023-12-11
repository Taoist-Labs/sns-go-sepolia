package sns

import "testing"

const testIndexerHost = "https://test-spp-indexer.seedao.tech"
const testSafeHost = "https://test-sns-api.seedao.tech"
const testRPC = "https://eth-goerli.g.alchemy.com/v2/MATWeLJN1bEGTjSmtyLedn0i34o1ISLD"
const testPublicResolverAddr = "0x6A80eA63cFfc6B10B764e1f26348832835520646"
const testBaseRegistrarAddr = "0x4C53Ff1A6a47E7089e1E727f83e7b7aEFCC9796B"

func TestResolve(t *testing.T) {
	tests := []struct {
		name string
		sns  string
		want string
	}{
		{
			name: "normal",
			sns:  "baiyu.seedao",
			want: "0x8C913aEc7443FE2018639133398955e0E17FB0C1",
		},
		{
			name: "not exist",
			sns:  "notexist.seedao",
			want: "0x0000000000000000000000000000000000000000",
		},
		{
			name: "empty",
			sns:  "",
			want: "0x0000000000000000000000000000000000000000",
		},
		{
			name: "contain sensitive word",
			sns:  "vitalik.seedao",
			want: "0x0000000000000000000000000000000000000000",
		},
		{
			name: "special char $",
			sns:  "$abc.seedao",
			want: "0x0000000000000000000000000000000000000000",
		},
		{
			name: "special char <",
			sns:  "<abc.seedao",
			want: "0x0000000000000000000000000000000000000000",
		},
		{
			name: "special char #",
			sns:  "#abc.seedao",
			want: "0x0000000000000000000000000000000000000000",
		},
	}

	// Query Indexer Success
	for _, tt := range tests {
		got := resolve(tt.sns, testSafeHost, testIndexerHost, "", "")
		if got != tt.want {
			t.Errorf("Query Indexer Success: resolve(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
		}
	}
	// Query Contract Success
	for _, tt := range tests {
		got := resolve(tt.sns, testSafeHost, "", testRPC, testPublicResolverAddr)
		if got != tt.want {
			t.Errorf("Query Contract Success: resolve(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
		}
	}
}

func TestResolves(t *testing.T) {
	sns := []string{
		"baiyu.seedao",
		"",
		"notexists.seedao",
		"vitalik.seedao",
		"$abc.seedao",
		"#abc.seedao",
	}
	want := []string{
		"0x8C913aEc7443FE2018639133398955e0E17FB0C1",
		"0x0000000000000000000000000000000000000000",
		"0x0000000000000000000000000000000000000000",
		"0x0000000000000000000000000000000000000000",
		"0x0000000000000000000000000000000000000000",
		"0x0000000000000000000000000000000000000000",
	}

	// Query Indexer Success
	got := resolves(sns, testSafeHost, testIndexerHost, "", "")
	if len(got) != len(want) {
		t.Errorf("Query Indexer Success: resolves(%s)'s result: %v, want: %v", sns, got, want)
	}
	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("Query Indexer Success: resolves(%s)'s result: %v, want: %v", sns, got, want)
		}
	}
	//Query Contract Success
	got = resolves(sns, testSafeHost, "", testRPC, testPublicResolverAddr)
	if len(got) != len(want) {
		t.Errorf("Query Contract Success: resolves(%s)'s result: %v, want: %v", sns, got, want)
	}
	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("Query Contract Success: resolves(%s)'s result: %v, want: %v", sns, got, want)
		}
	}
}

func TestName(t *testing.T) {
	tests := []struct {
		name string
		addr string
		want string
	}{
		{
			name: "normal",
			addr: "0x8C913aEc7443FE2018639133398955e0E17FB0C1",
			want: "baiyu.seedao",
		},
		{
			name: "not exist",
			addr: "0x0000000000000000000000000000000000000000",
			want: "",
		},
		{
			name: "contain sensitive word",
			addr: "0xc1eE7cB74583D1509362467443C44f1FCa981283",
			want: "",
		},
	}

	// Query Indexer Success
	for _, tt := range tests {
		got := name(tt.addr, testSafeHost, testIndexerHost, "", "")
		if got != tt.want {
			t.Errorf("Query Indexer Success: resolve(%s)'s result: %v, want: %v", tt.addr, got, tt.want)
		}
	}
	// Query Contract Success
	for _, tt := range tests {
		got := name(tt.addr, testSafeHost, "", testRPC, testPublicResolverAddr)
		if got != tt.want {
			t.Errorf("Query Contract Success: resolve(%s)'s result: %v, want: %v", tt.addr, got, tt.want)
		}
	}
}

func TestNames(t *testing.T) {
	addr := []string{
		"0x8C913aEc7443FE2018639133398955e0E17FB0C1",
		"0x0000000000000000000000000000000000000000",
		"0xc1eE7cB74583D1509362467443C44f1FCa981283",
	}

	want := []string{
		"baiyu.seedao", "", "",
	}

	// Query Indexer Success
	got := names(addr, testSafeHost, testIndexerHost, "", "")
	if len(got) != len(want) {
		t.Errorf("Query Indexer Success: names(%s)'s result: %v, want: %v", addr, got, want)
	}
	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("Query Indexer Success: names(%s)'s result: %v, want: %v", addr, got, want)
		}
	}
	// Query Contract Success
	got = names(addr, testSafeHost, "", testRPC, testPublicResolverAddr)
	if len(got) != len(want) {
		t.Errorf("Query Contract Success: names(%s)'s result: %v, want: %v", addr, got, want)
	}
	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("Query Contract Success: names(%s)'s result: %v, want: %v", addr, got, want)
		}
	}
}

func TestTokenId(t *testing.T) {
	tests := []struct {
		name string
		sns  string
		want string
	}{
		{
			name: "normal",
			sns:  "baiyu.seedao",
			want: "53",
		},
		//{
		//	name: "not exist",
		//	sns:  "baiyu2.seedao",
		//	want: "", // TODO Indexer return "", but contract return "0"
		//},
	}

	// Query Indexer Success
	for _, tt := range tests {
		got := tokenId(tt.sns, testSafeHost, testIndexerHost, "", "")
		if got != tt.want {
			t.Errorf("Query Indexer Success: tokenId(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
		}
	}
	// Query Contract Success
	for _, tt := range tests {
		got := tokenId(tt.sns, testSafeHost, "", testRPC, testBaseRegistrarAddr)
		if got != tt.want {
			t.Errorf("Query Contract Success: tokenId(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
		}
	}

	// ---->  <-------
	// Indexer return "", but contract return "0"
	got := tokenId("notexists.seedao", testSafeHost, testIndexerHost, "", "")
	if got != "" {
		t.Errorf("Query Indexer Success: tokenId(%s)'s result: %v, want: %v", "notexists.seedao", got, "")
	}
	got = tokenId("notexists.seedao", testSafeHost, "", testRPC, testBaseRegistrarAddr)
	if got != "0" {
		t.Errorf("Query Contract Success: tokenId(%s)'s result: %v, want: %v", "notexists.seedao", got, "0")
	}
	// ---->  <-------
}
