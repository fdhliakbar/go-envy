package envy

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	os.Setenv("FOO", "bar")
	if val := Get("FOO", "baz"); val != "bar" {
		t.Errorf("expected 'bar', got '%s'", val)
	}
	if val := Get("NOT_SET", "baz"); val != "baz" {
		t.Errorf("expected default 'baz', got '%s'", val)
	}
}

func TestRequire(t *testing.T) {
	os.Setenv("REQ", "val")
	if val := Require("REQ"); val != "val" {
		t.Errorf("expected 'val', got '%s'", val)
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Require should panic if env not set")
		}
	}()
	Require("NOT_EXIST")
}

func TestGetInt(t *testing.T) {
	os.Setenv("INT", "42")
	if val := GetInt("INT", 10); val != 42 {
		t.Errorf("expected 42, got %d", val)
	}
	if val := GetInt("NOT_INT", 10); val != 10 {
		t.Errorf("expected default 10, got %d", val)
	}
	os.Setenv("BAD_INT", "abc")
	if val := GetInt("BAD_INT", 7); val != 7 {
		t.Errorf("expected default 7, got %d", val)
	}
}

func TestGetBool(t *testing.T) {
	os.Setenv("BOOL_TRUE", "true")
	os.Setenv("BOOL_FALSE", "false")
	if !GetBool("BOOL_TRUE", false) {
		t.Errorf("expected true")
	}
	if GetBool("BOOL_FALSE", true) {
		t.Errorf("expected false")
	}
	if !GetBool("NOT_BOOL", true) {
		t.Errorf("expected default true")
	}
	os.Setenv("BOOL_ONE", "1")
	if !GetBool("BOOL_ONE", false) {
		t.Errorf("expected true for '1'")
	}
	os.Setenv("BOOL_ZERO", "0")
	if GetBool("BOOL_ZERO", true) {
		t.Errorf("expected false for '0'")
	}
}

func TestLoad(t *testing.T) {
	f, err := os.CreateTemp("", "testenv*.env")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(f.Name())
	f.WriteString("FOO=bar\nBAR=baz\n# comment\nEMPTY=\n")
	f.Close()

	if err := Load(f.Name()); err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	if val := os.Getenv("FOO"); val != "bar" {
		t.Errorf("expected 'bar', got '%s'", val)
	}
	if val := os.Getenv("BAR"); val != "baz" {
		t.Errorf("expected 'baz', got '%s'", val)
	}
}
