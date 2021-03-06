package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestReplaceAndAppendEnvVars(t *testing.T) {
	var (
		d = []string{"HOME=/"}
		o = []string{"HOME=/root", "TERM=xterm"}
	)

	env := ReplaceOrAppendEnvValues(d, o)
	if len(env) != 2 {
		t.Fatalf("expected len of 2 got %d", len(env))
	}
	if env[0] != "HOME=/root" {
		t.Fatalf("expected HOME=/root got '%s'", env[0])
	}
	if env[1] != "TERM=xterm" {
		t.Fatalf("expected TERM=xterm got '%s'", env[1])
	}
}

func TestReadDockerIgnore(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "dockerignore-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	di, err := ReadDockerIgnore(nil)
	if err != nil {
		t.Fatalf("Expected not to have error, got %v", err)
	}

	if diLen := len(di); diLen != 0 {
		t.Fatalf("Expected to have zero dockerignore entry, got %d", diLen)
	}

	diName := filepath.Join(tmpDir, ".dockerignore")
	content := fmt.Sprintf("test1\n/test2\n/a/file/here\n\nlastfile")
	err = ioutil.WriteFile(diName, []byte(content), 0777)
	if err != nil {
		t.Fatal(err)
	}

	diFd, err := os.Open(diName)
	if err != nil {
		t.Fatal(err)
	}
	di, err = ReadDockerIgnore(diFd)
	if err != nil {
		t.Fatal(err)
	}

	if di[0] != "test1" {
		t.Fatalf("First element is not test1")
	}
	if di[1] != "/test2" {
		t.Fatalf("Second element is not /test2")
	}
	if di[2] != "/a/file/here" {
		t.Fatalf("Third element is not /a/file/here")
	}
	if di[3] != "lastfile" {
		t.Fatalf("Fourth element is not lastfile")
	}
}
