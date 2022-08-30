package server

import (
	"reflect"
	"testing"
)

func TestNewServer(t *testing.T) {
	want := &Server{cfg: defaultConfig}
	got := NewServer()
	if !reflect.DeepEqual(want, got) {
		t.Fatal(want, got)
	}

	want = &Server{cfg: defaultConfig}
	want.cfg.port = 5555
	got = NewServer(PortOpt(5555))
	if !reflect.DeepEqual(want, got) {
		t.Fatal(want, got)
	}

	want = &Server{cfg: Config{
		hostname: "127.0.0.1",
		port:     5555,
		protocol: "https",
	}}
	got = NewServer(HostnameOpt("127.0.0.1"), ProtocolOpt("https"), PortOpt(5555))
	if !reflect.DeepEqual(want, got) {
		t.Fatal(want, got)
	}

	want = &Server{cfg: defaultConfig}
	want.cfg.port = 9090
	got = NewServer(PortOpt(8888), PortOpt(9090))
	if !reflect.DeepEqual(want, got) {
		t.Fatal(want, got)
	}
}
