package vcs_test

import (
	"flag"
	"log"
	"net"
	"os"
	"testing"

	"sourcegraph.com/sourcegraph/sourcegraph/cmd/gitserver/server"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/gitserver"
)

func TestMain(m *testing.M) {
	flag.Parse()

	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatalf("listen failed: %s", err)
	}
	go (&server.Server{InsecureSkipCheckVerifySSH: true}).ServeLegacy(l)

	gitserver.DefaultClient = &gitserver.Client{Addrs: []string{l.Addr().String()}}

	os.Exit(m.Run())
}
