// Package schemas provides the vendored schema files, either embedded in the
// binary (copied from vendor/skills-for-copilot-studio/reference by
// `go generate`) or read from a directory on disk.
package schemas

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

//go:embed data
var data embed.FS

// File names, relative to the schema root (mirrors the upstream reference/ dir).
const (
	BotSchema     = "bot.schema.yaml-authoring.json"
	CardSchema    = "adaptive-card.schema.json"
	ConnectorsDir = "connectors"
	VersionFile   = "VERSION"
)

// Bundle is a schema root plus a human-readable version string.
type Bundle struct {
	FS      fs.FS
	Version string // e.g. "vendor/skills-for-copilot-studio @ 9c42d3a"
}

// Open returns the embedded schemas when dir is empty, else the schemas in dir
// (an upstream reference/ directory).
func Open(dir string) (*Bundle, error) {
	if dir == "" {
		sub, err := fs.Sub(data, "data")
		if err != nil {
			return nil, err
		}
		return &Bundle{FS: sub, Version: version(sub, "embedded")}, nil
	}
	st, err := os.Stat(dir)
	if err != nil {
		return nil, fmt.Errorf("schema dir: %w", err)
	}
	if !st.IsDir() {
		return nil, fmt.Errorf("schema dir %s is not a directory", dir)
	}
	fsys := os.DirFS(dir)
	if _, err := fs.Stat(fsys, BotSchema); err != nil {
		return nil, fmt.Errorf("schema dir %s has no %s", dir, BotSchema)
	}
	return &Bundle{FS: fsys, Version: version(fsys, dir)}, nil
}

func version(fsys fs.FS, fallback string) string {
	b, err := fs.ReadFile(fsys, VersionFile)
	if err != nil {
		return fallback
	}
	return strings.TrimSpace(string(b))
}
