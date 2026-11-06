// Package describe holds the curated descriptions overlay. The Copilot Studio
// bot schema ships only placeholder descriptions (BotSchema:$$X_y_title$$), so
// human-written text lives in descriptions.yaml, keyed by:
//
//	Def              a definition (kind)
//	Def.field        a field of Def
//	Owner.field      a field shared by every definition whose placeholder token names Owner
//	enum.Enum.Value  one value of an enum
//
// A separate top-level `roots:` key lists, per source, the kinds that can be
// the root of a file; `cse list` shows those by default.
package describe

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed descriptions.yaml
var embedded []byte

// Entry is one curated description.
type Entry struct {
	Description string `yaml:"description"`
	Docs        string `yaml:"docs"`
}

// UnmarshalYAML accepts either a plain string or {description, docs}.
func (e *Entry) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind == yaml.ScalarNode {
		e.Description = n.Value
		return nil
	}
	type plain Entry
	var p plain
	if err := n.Decode(&p); err != nil {
		return err
	}
	*e = Entry(p)
	return nil
}

// Root is a kind that can be the root of a file, and the file(s) it lives in.
type Root struct {
	Kind string `yaml:"kind"`
	File string `yaml:"file"`
}

// Overlay is the parsed descriptions file, one map per schema source.
type Overlay struct {
	sections map[string]map[string]Entry
	roots    map[string][]Root
}

// Load parses the embedded overlay.
func Load() (*Overlay, error) { return Parse(embedded) }

// Parse parses an overlay document.
func Parse(b []byte) (*Overlay, error) {
	var doc struct {
		Roots    map[string][]Root           `yaml:"roots"`
		Sections map[string]map[string]Entry `yaml:",inline"`
	}
	if err := yaml.Unmarshal(b, &doc); err != nil {
		return nil, fmt.Errorf("descriptions overlay: %w", err)
	}
	o := &Overlay{sections: doc.Sections, roots: doc.Roots}
	if o.sections == nil {
		o.sections = map[string]map[string]Entry{}
	}
	return o, nil
}

// Section returns the overlay for one schema source ("bot", "card", "connector").
func (o *Overlay) Section(name string) Section {
	if o == nil {
		return Section{}
	}
	return Section{entries: o.sections[name]}
}

// Roots returns the curated file-root kinds of a source, in display order
// (nil when the source has none, i.e. every definition is listed).
func (o *Overlay) Roots(source string) []Root {
	if o == nil {
		return nil
	}
	return o.roots[source]
}

// Section is the overlay for one schema source.
type Section struct {
	entries map[string]Entry
}

// Get returns the first non-empty entry among keys.
func (s Section) Get(keys ...string) (Entry, bool) {
	for _, k := range keys {
		if e, ok := s.entries[k]; ok && e.Description != "" {
			return e, true
		}
	}
	return Entry{}, false
}

// Keys returns all keys, sorted.
func (s Section) Keys() []string {
	keys := make([]string, 0, len(s.entries))
	for k := range s.entries {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// Len is the number of entries.
func (s Section) Len() int { return len(s.entries) }

var tokenRE = regexp.MustCompile(`^BotSchema:\$\$([A-Za-z0-9]+)_([A-Za-z0-9]+)_(?:title|description)\$\$$`)

// TokenOwner parses a field placeholder token such as
// "BotSchema:$$DialogAction_id_description$$" into ("DialogAction", "id").
func TokenOwner(token string) (owner, field string, ok bool) {
	m := tokenRE.FindStringSubmatch(token)
	if m == nil || m[1] == "enum" {
		return "", "", false
	}
	return m[1], m[2], true
}

// IsPlaceholder reports whether a schema title/description is an untranslated token.
func IsPlaceholder(s string) bool {
	return strings.Contains(s, "BotSchema:$$")
}

// StripNoKind maps "FooNoKind" to "Foo"; NoKind variants share their kind's docs.
func StripNoKind(name string) string {
	return strings.TrimSuffix(name, "NoKind")
}
