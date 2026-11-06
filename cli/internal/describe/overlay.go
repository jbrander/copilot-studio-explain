// Package describe holds the curated descriptions overlay. The Copilot Studio
// bot schema ships only placeholder descriptions (BotSchema:$$X_y_title$$), so
// human-written text lives in descriptions.yaml, keyed by:
//
//	Def              a definition (kind)
//	Def.field        a field of Def
//	Owner.field      a field shared by every definition whose placeholder token names Owner
//	enum.Enum.Value  one value of an enum
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

// Overlay is the parsed descriptions file, one map per schema source.
type Overlay struct {
	sections map[string]map[string]Entry
}

// Load parses the embedded overlay.
func Load() (*Overlay, error) { return Parse(embedded) }

// Parse parses an overlay document.
func Parse(b []byte) (*Overlay, error) {
	var raw map[string]map[string]Entry
	if err := yaml.Unmarshal(b, &raw); err != nil {
		return nil, fmt.Errorf("descriptions overlay: %w", err)
	}
	o := &Overlay{sections: map[string]map[string]Entry{}}
	for k, v := range raw {
		o.sections[k] = v
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
