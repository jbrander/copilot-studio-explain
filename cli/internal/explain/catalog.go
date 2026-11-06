// Package explain resolves `kubectl explain`-style resource paths
// (e.g. "AdaptiveDialog.beginDialog.OnRecognizedIntent.actions") against the
// three schema sources.
package explain

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"strings"
	"sync"

	"copilot-studio-explain/cli/internal/connector"
	"copilot-studio-explain/cli/internal/describe"
	"copilot-studio-explain/cli/internal/jsonschema"
	"copilot-studio-explain/cli/internal/model"
	"copilot-studio-explain/cli/internal/schemas"
)

// Source names.
const (
	Bot       = "bot"
	Card      = "card"
	Connector = "connector"
)

// SourceNames lists the sources in display order.
var SourceNames = []string{Bot, Card, Connector}

// Source is one schema source.
type Source interface {
	Name() string
	Names() []string
	Lookup(name string) (*model.Node, bool)
	ValidKey(key string) bool
}

// RawSource is implemented by sources that can return a definition's raw JSON.
type RawSource interface {
	Definition(name string) (json.RawMessage, bool)
}

// Catalog loads sources lazily from a schema bundle.
type Catalog struct {
	bundle  *schemas.Bundle
	overlay *describe.Overlay

	mu      sync.Mutex
	sources map[string]Source
}

// Open creates a catalog over the embedded schemas (dir == "") or a directory.
func Open(dir string) (*Catalog, error) {
	b, err := schemas.Open(dir)
	if err != nil {
		return nil, err
	}
	ov, err := describe.Load()
	if err != nil {
		return nil, err
	}
	return &Catalog{bundle: b, overlay: ov, sources: map[string]Source{}}, nil
}

// Version describes where the schemas came from.
func (c *Catalog) Version() string { return c.bundle.Version }

// Overlay returns the curated descriptions.
func (c *Catalog) Overlay() *describe.Overlay { return c.overlay }

// Source returns (loading on first use) a schema source by name.
func (c *Catalog) Source(name string) (Source, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if s, ok := c.sources[name]; ok {
		return s, nil
	}
	var (
		s   Source
		err error
	)
	switch name {
	case Bot:
		s, err = c.loadJSON(schemas.BotSchema, jsonschema.Options{Source: Bot, Discriminator: "kind", Overlay: c.overlay.Section(Bot)})
	case Card:
		s, err = c.loadJSON(schemas.CardSchema, jsonschema.Options{Source: Card, Discriminator: "type", Overlay: c.overlay.Section(Card)})
	case Connector:
		s, err = connector.New(c.bundle.FS, schemas.ConnectorsDir, c.overlay.Section(Connector))
	default:
		return nil, fmt.Errorf("unknown schema %q (want one of: %s)", name, strings.Join(SourceNames, ", "))
	}
	if err != nil {
		return nil, err
	}
	c.sources[name] = s
	return s, nil
}

func (c *Catalog) loadJSON(file string, opts jsonschema.Options) (Source, error) {
	b, err := fs.ReadFile(c.bundle.FS, file)
	if err != nil {
		return nil, fmt.Errorf("%s schema: %w", opts.Source, err)
	}
	f, err := jsonschema.Parse(b)
	if err != nil {
		return nil, fmt.Errorf("%s schema %s: %w", opts.Source, file, err)
	}
	return jsonschema.New(f, opts), nil
}

// SplitSource splits "card:Input.Text" into ("card", "Input.Text"). Without a
// recognised prefix the source is the bot schema.
func SplitSource(expr string) (source, rest string) {
	if before, after, ok := strings.Cut(expr, ":"); ok {
		for _, s := range SourceNames {
			if strings.EqualFold(before, s) || strings.EqualFold(before, s+"s") {
				return s, after
			}
		}
	}
	return Bot, expr
}
