package sns

import "testing"

const testIndexerHost = "http://localhost:3000"
const testSafeHost = "http://localhost:8090"
const testRPC = "https://eth-sepolia.g.alchemy.com/v2/H43zK7UnIN2v7u2ZoTbizIPnXkylKIZl"
const testPublicResolverAddr = "0x4ffCfd37C362B415E4c4A607815f5dB6A297Ed8A"

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
