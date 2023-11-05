package sns

import "testing"

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

	for _, tt := range tests {
		got := Resolve(tt.sns)
		if got != tt.want {
			t.Errorf("Resolve(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
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
	got := Resolves(sns)

	if len(got) != len(want) {
		t.Errorf("Resolves(%s)'s result: %v, want: %v", sns, got, want)
	}
	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("Resolves(%s)'s result: %v, want: %v", sns, got, want)
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

	for _, tt := range tests {
		got := Name(tt.addr)
		if got != tt.want {
			t.Errorf("Resolve(%s)'s result: %v, want: %v", tt.addr, got, tt.want)
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
	got := Names(addr)

	if len(got) != len(want) {
		t.Errorf("Names(%s)'s result: %v, want: %v", addr, got, want)
	}
	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("Names(%s)'s result: %v, want: %v", addr, got, want)
		}
	}
}
