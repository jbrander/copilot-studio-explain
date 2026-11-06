// Package jsonschema turns the two JSON-schema sources (the Copilot Studio bot
// authoring schema and the Adaptive Card schema) into model nodes.
package jsonschema

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// Schema is the subset of JSON Schema used by the vendored schemas.
type Schema struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Type        typeList  `json:"type"`
	Ref         string    `json:"$ref"`
	Role        string    `json:"$role"`
	Const       any       `json:"const"`
	Default     any       `json:"default"`
	Enum        []any     `json:"enum"`
	Pattern     string    `json:"pattern"`
	Format      string    `json:"format"`
	Version     string    `json:"version"`
	Examples    []any     `json:"examples"`
	Properties  props     `json:"properties"`
	Required    required  `json:"required"`
	Items       *Schema   `json:"items"`
	AddProps    *addProps `json:"additionalProperties"`
	OneOf       []*Schema `json:"oneOf"`
	AnyOf       []*Schema `json:"anyOf"`
	AllOf       []*Schema `json:"allOf"`

	hasDefault bool
	hasConst   bool
	raw        json.RawMessage
}

// UnmarshalJSON keeps the raw bytes and records whether default/const were
// present (they may legitimately be false/0/null).
func (s *Schema) UnmarshalJSON(b []byte) error {
	type plain Schema
	var p plain
	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}
	*s = Schema(p)
	s.raw = bytes.Clone(b)
	var keys map[string]json.RawMessage
	if err := json.Unmarshal(b, &keys); err == nil {
		_, s.hasDefault = keys["default"]
		_, s.hasConst = keys["const"]
	}
	return nil
}

// IsEmpty reports a `{}` schema (used by the card schema for inherited props).
func (s *Schema) IsEmpty() bool {
	return s == nil || len(bytes.TrimSpace(s.raw)) <= 2
}

// Raw returns the source JSON fragment.
func (s *Schema) Raw() json.RawMessage { return s.raw }

// RefName returns the definition name of a "#/definitions/X" reference.
func (s *Schema) RefName() string {
	if s == nil || s.Ref == "" {
		return ""
	}
	return s.Ref[strings.LastIndex(s.Ref, "/")+1:]
}

// typeList accepts "type": "string" and "type": ["string", "null"].
type typeList []string

func (t *typeList) UnmarshalJSON(b []byte) error {
	var one string
	if err := json.Unmarshal(b, &one); err == nil {
		*t = typeList{one}
		return nil
	}
	var many []string
	if err := json.Unmarshal(b, &many); err != nil {
		return fmt.Errorf("type: %w", err)
	}
	*t = many
	return nil
}

func (t typeList) has(name string) bool {
	for _, x := range t {
		if x == name {
			return true
		}
	}
	return false
}

// required accepts the usual list of property names and the draft-3 form
// `"required": true` on a property itself (used twice by the card schema).
type required struct {
	Names []string
	Self  bool
}

func (r *required) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &r.Self); err == nil {
		return nil
	}
	return json.Unmarshal(b, &r.Names)
}

// props is an object of property schemas.
type props map[string]*Schema

// addProps accepts additionalProperties as a bool or a schema.
type addProps struct {
	Allowed bool
	Schema  *Schema
}

func (a *addProps) UnmarshalJSON(b []byte) error {
	var allowed bool
	if err := json.Unmarshal(b, &allowed); err == nil {
		a.Allowed = allowed
		return nil
	}
	a.Allowed = true
	a.Schema = &Schema{}
	return json.Unmarshal(b, a.Schema)
}

// File is a whole schema document.
type File struct {
	Definitions map[string]*Schema `json:"definitions"`
	OneOf       []*Schema          `json:"oneOf"`
	AnyOf       []*Schema          `json:"anyOf"`
}

// Parse decodes a schema document.
func Parse(b []byte) (*File, error) {
	var f File
	if err := json.Unmarshal(b, &f); err != nil {
		return nil, err
	}
	if f.Definitions == nil {
		return nil, fmt.Errorf("schema has no definitions")
	}
	return &f, nil
}

// branchRef unwraps a union branch: either {"$ref": X} or the card schema's
// {"required": [...], "allOf": [{"$ref": X}]}.
func branchRef(s *Schema) string {
	if s == nil {
		return ""
	}
	if s.Ref != "" {
		return s.RefName()
	}
	if len(s.AllOf) == 1 && s.AllOf[0].Ref != "" && len(s.Properties) == 0 {
		return s.AllOf[0].RefName()
	}
	return ""
}

// isExprBranch reports the bot schema's Power Fx branch: {"type":"string","pattern":"=.*"}.
func isExprBranch(s *Schema) bool {
	return s != nil && s.Pattern == "=.*" && s.Ref == ""
}

// isPatternOnly reports the card schema's case-insensitive enum mirror: {"pattern": "..."}.
func isPatternOnly(s *Schema) bool {
	return s != nil && s.Pattern != "" && s.Pattern != "=.*" && len(s.Type) == 0 && s.Ref == ""
}

// literalValues returns the const/enum literals of a branch.
func literalValues(s *Schema) []string {
	if s == nil {
		return nil
	}
	if s.hasConst {
		return []string{fmt.Sprint(s.Const)}
	}
	if len(s.Enum) > 0 && s.Ref == "" {
		out := make([]string, 0, len(s.Enum))
		for _, v := range s.Enum {
			out = append(out, fmt.Sprint(v))
		}
		return out
	}
	return nil
}

func formatValue(v any) string {
	switch x := v.(type) {
	case string:
		return x
	case nil:
		return "null"
	default:
		b, err := json.Marshal(x)
		if err != nil {
			return fmt.Sprint(x)
		}
		return string(b)
	}
}
