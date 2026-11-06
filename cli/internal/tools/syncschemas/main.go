// Command syncschemas copies the vendored schema files into
// internal/schemas/data so they can be embedded (go:embed cannot reach
// outside the module). Run via `go generate` from the cli/ directory.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	src := flag.String("src", "../vendor/skills-for-copilot-studio/reference", "upstream reference/ directory")
	dst := flag.String("dst", "internal/schemas/data", "destination directory")
	flag.Parse()
	if err := run(*src, *dst); err != nil {
		fmt.Fprintln(os.Stderr, "syncschemas:", err)
		os.Exit(1)
	}
}

func run(src, dst string) error {
	if _, err := os.Stat(filepath.Join(src, "bot.schema.yaml-authoring.json")); err != nil {
		return fmt.Errorf("%s does not look like the vendored reference dir (is the submodule checked out? `git submodule update --init`): %w", src, err)
	}
	if err := os.RemoveAll(dst); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(dst, "connectors"), 0o755); err != nil {
		return err
	}
	files := []string{"bot.schema.yaml-authoring.json", "adaptive-card.schema.json"}
	conns, err := filepath.Glob(filepath.Join(src, "connectors", "*.yml"))
	if err != nil {
		return err
	}
	for _, c := range conns {
		files = append(files, filepath.Join("connectors", filepath.Base(c)))
	}
	for _, f := range files {
		b, err := os.ReadFile(filepath.Join(src, f))
		if err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(dst, f), b, 0o644); err != nil {
			return err
		}
	}
	version := "vendor/skills-for-copilot-studio @ " + commit(filepath.Dir(src))
	if err := os.WriteFile(filepath.Join(dst, "VERSION"), []byte(version+"\n"), 0o644); err != nil {
		return err
	}
	fmt.Printf("syncschemas: copied %d files from %s (%s)\n", len(files), src, version)
	return nil
}

func commit(repo string) string {
	out, err := exec.Command("git", "-C", repo, "rev-parse", "--short", "HEAD").Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(bytes.TrimSpace(out)))
}
