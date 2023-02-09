package ssh_conf

import (
	"fmt"
	"testing"

	"github.com/kevinburke/ssh_config"
)

func TestParseFile(t *testing.T) {
	cfg, err := parseFile("testdata/host1.conf")

	if err != nil {
		t.Error(fmt.Errorf("error parsing file: %w", err))
	}

	if cfg == nil {
		t.Error("cfg is nil")
	}

	hostname, err := cfg.Get("host1", "Hostname")
	if err != nil {
		t.Error(err)
	}

	if hostname != "1.2.3.1" {
		t.Errorf("hostname is %s, expected 1.2.3.1", hostname)
	}
}

func TestParseDir(t *testing.T) {
	cfg := new(ssh_config.Config)
	err := parseDir(cfg, "testdata/")

	if err != nil {
		t.Fatal(err)
	}

	if cfg == nil {
		t.Fatal("cfg is nil")
	}
	fmt.Println(cfg)

	if len(cfg.Hosts) != 4 {
		t.Fatalf("expected 4 hosts, got %d", len(cfg.Hosts))
	}

}
