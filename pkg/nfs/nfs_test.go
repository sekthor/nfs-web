package nfs

import (
	"strings"
	"testing"
)

func TestMarshalHost(t *testing.T) {
	cases := []Host{
		Host{
			Addr: "127.0.0.1/32",
			Options: []string{
				"no_root_squash",
			},
		},
	}

	want := "127.0.0.1/32(no_root_squash)"
	got := cases[0].Marshal()

	if want != got {
		t.Fatalf("wanted '%s'; got: '%s'", want, got)
	}
}

func TestMarshalShare(t *testing.T) {
	test := NfsShare{
		Directory: "/mnt/nfs/share1",
		Hosts: []Host{
			Host{
				Addr: "127.0.0.1/32",
				Options: []string{
					"rw", "sync", "no_root_squash",
				},
			},
			Host{
				Addr: "127.0.0.11/32",
				Options: []string{
					"rw", "sync",
				},
			},
		},
	}
	want := "/mnt/nfs/share1 127.0.0.1/32(rw,sync,no_root_squash) 127.0.0.11/32(rw,sync)"
	got := test.Marshal()

	if want != got {
		t.Fatalf("wanted '%s'; got: '%s'", want, got)
	}
}

func TestUnmarshalShare(t *testing.T) {
	test := "/mnt/nfs/share1 127.0.0.1/32(rw,sync,no_root_squash) 127.0.0.11/32(rw,sync)"

	want := NfsShare{
		Directory: "/mnt/nfs/share1",
		Hosts: []Host{
			Host{
				Addr: "127.0.0.1/32",
				Options: []string{
					"rw", "sync", "no_root_squash",
				},
			},
			Host{
				Addr: "127.0.0.11/32",
				Options: []string{
					"rw", "sync",
				},
			},
		},
	}

	var got NfsShare
	err := UnmarshalShare(test, &got)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if want.Directory != got.Directory {
		t.Fatalf("directory did not match")
	}

	for i := 0; i < len(want.Hosts); i++ {
		if strings.Join(want.Hosts[i].Options, ",") != strings.Join(got.Hosts[i].Options, ",") {
			t.Fatal("wanted", want.Hosts[i].Options, "; got", got.Hosts[i].Options)
		}
	}
}
