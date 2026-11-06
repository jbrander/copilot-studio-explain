// Package model is the source-neutral tree that every schema source (bot JSON
// schema, Adaptive Card JSON schema, connector YAML) is translated into, and
// that the renderers print.
package model

import "encoding/json"

// Shape is the structural category of a Node.
type Shape int

const (
	// Object has named fields.
	Object Shape = iota
	// Union is "one of" several member kinds, chosen by a discriminator
	// (`kind:` in bot YAML, `type:` in Adaptive Cards).
	Union
	// Enum is a closed set of literal values.
	Enum
	// Scalar is a leaf value (string, boolean, expression, ...).
	Scalar
)

func (s Shape) String() string {
	switch s {
	case Object:
		return "object"
	case Union:
		return "union"
	case Enum:
		return "enum"
	default:
		return "scalar"
	}
}

// Node is one explainable type: a schema definition, an inline object, a
// connector, a connector operation or a connector record.
type Node struct {
	Name        string // definition / connector / operation name; "" for anonymous inline objects
	Label       string // type label without angle brackets, e.g. "SendActivity", "boolean | =expr"
	Shape       Shape
	Description string
	Docs        string // repo-relative docs link from the curated overlay

	// Discriminator is the property that selects a union member ("kind" or
	// "type"). For an Object it names the property whose constant value is
	// DiscriminatorValue.
	Discriminator      string
	DiscriminatorValue string

	Members     []Member    // Union
	MemberNotes []string    // Union: non-navigable branches, e.g. "string shorthand"
	EnumValues  []EnumValue // Enum
	Fields      []*Field    // Object
	// Shorthands are non-object spellings an Object also accepts, e.g. a
	// prebuilt entity written as just `"BooleanPrebuiltEntity"`.
	Shorthands []string

	// Raw is the source fragment (json.RawMessage for JSON schemas, a decoded
	// YAML value or a func() any producing one for connectors), used by `-o json`.
	Raw any
}

// Member is one choice of a Union.
type Member struct {
	Name        string // definition name
	Value       string // discriminator value (e.g. kind: SendActivity); "" if the member has none
	Description string
	Node        func() *Node
}

// EnumValue is one allowed literal of an Enum.
type EnumValue struct {
	Value       string
	Description string
}

// Field is one property of an Object.
type Field struct {
	Name        string
	Label       string // type label without angle brackets
	Required    bool
	Default     any
	HasDefault  bool
	Const       string   // the only allowed value (e.g. the `kind:` property)
	Enum        []string // allowed literal values, when the type is an enum
	Description string
	Notes       []string // extra facts: "since v1.3", "display name: Foo", ...
	Target      func() *Node
	Raw         any
}

// Child returns the node a field points at, or nil for leaves.
func (f *Field) Child() *Node {
	if f == nil || f.Target == nil {
		return nil
	}
	return f.Target()
}

// RawJSON marshals a Raw value to JSON.
func RawJSON(raw any) json.RawMessage {
	if f, ok := raw.(func() any); ok {
		raw = f()
	}
	switch v := raw.(type) {
	case nil:
		return json.RawMessage("null")
	case json.RawMessage:
		return v
	default:
		b, err := json.Marshal(v)
		if err != nil {
			return json.RawMessage("null")
		}
		return b
	}
}
