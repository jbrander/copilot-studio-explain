package schemas

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

// TestEmbeddedMatchesVendor guards against forgetting `go generate` after the
// submodule is bumped.
func TestEmbeddedMatchesVendor(t *testing.T) {
	vendor := filepath.Join("..", "..", "..", "vendor", "skills-for-copilot-studio", "reference")
	if _, err := os.Stat(filepath.Join(vendor, BotSchema)); err != nil {
		t.Skip("vendor submodule not checked out")
	}
	b, err := Open("")
	if err != nil {
		t.Fatal(err)
	}
	vfs := os.DirFS(vendor)
	files := []string{BotSchema, CardSchema}
	conns, _ := fs.Glob(vfs, ConnectorsDir+"/*.yml")
	files = append(files, conns...)
	embedded, _ := fs.Glob(b.FS, ConnectorsDir+"/*.yml")
	if len(embedded) != len(conns) {
		t.Errorf("embedded has %d connectors, vendor has %d; run `go generate` in cli/", len(embedded), len(conns))
	}
	for _, f := range files {
		want, err := fs.ReadFile(vfs, f)
		if err != nil {
			t.Fatal(err)
		}
		got, err := fs.ReadFile(b.FS, f)
		if err != nil {
			t.Errorf("%s missing from embedded data; run `go generate` in cli/", f)
			continue
		}
		if !bytes.Equal(got, want) {
			t.Errorf("%s differs from vendor; run `go generate` in cli/", f)
		}
	}
}

func TestOpenDir(t *testing.T) {
	if _, err := Open(t.TempDir()); err == nil {
		t.Error("expected error for a dir without schemas")
	}
	if _, err := Open(filepath.Join(t.TempDir(), "missing")); err == nil {
		t.Error("expected error for a missing dir")
	}
}
