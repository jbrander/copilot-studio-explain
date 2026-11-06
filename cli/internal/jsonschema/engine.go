package jsonschema

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"copilot-studio-explain/cli/internal/describe"
	"copilot-studio-explain/cli/internal/model"
)

// Options configures an Engine for one schema source.
type Options struct {
	// Source is the overlay section and display name: "bot" or "card".
	Source string
	// Discriminator is the property that selects a union member: "kind" or "type".
	Discriminator string
	// Overlay supplies curated descriptions.
	Overlay describe.Section
}

// Engine translates a schema document into model nodes, lazily and with caching.
type Engine struct {
	opts   Options
	defs   map[string]*Schema
	names  []string
	lower  map[string]string
	cache  map[string]*model.Node
	merged map[string]*merged
	owners map[string]bool // "Owner.field" keys found in placeholder tokens
	// shorthands records the non-object alternatives of "shorthand or object" definitions.
	shorthands map[string][]string
}

type merged struct {
	props    map[string]*Schema
	required map[string]bool
}

// New builds an engine over a parsed schema file.
func New(f *File, opts Options) *Engine {
	e := &Engine{
		opts:       opts,
		defs:       f.Definitions,
		lower:      make(map[string]string, len(f.Definitions)),
		cache:      map[string]*model.Node{},
		merged:     map[string]*merged{},
		owners:     map[string]bool{},
		shorthands: map[string][]string{},
	}
	for name := range f.Definitions {
		e.names = append(e.names, name)
		e.lower[strings.ToLower(name)] = name
	}
	sort.Strings(e.names)
	// Index the owners named by placeholder tokens, including properties of
	// inline object branches ("shorthand or object" definitions).
	for _, name := range e.names {
		for _, ps := range e.merge(name).props {
			for _, tok := range []string{ps.Description, ps.Title} {
				if owner, f, ok := describe.TokenOwner(tok); ok {
					e.owners[owner+"."+f] = true
					break
				}
			}
		}
	}
	return e
}

// Name is the source name ("bot" or "card").
func (e *Engine) Name() string { return e.opts.Source }

// Names lists every definition, sorted.
func (e *Engine) Names() []string { return e.names }

// Lookup finds a definition by name, case-insensitively.
func (e *Engine) Lookup(name string) (*model.Node, bool) {
	canonical, ok := e.canonical(name)
	if !ok {
		return nil, false
	}
	return e.node(canonical), true
}

// Definition returns the raw JSON of a definition.
func (e *Engine) Definition(name string) (json.RawMessage, bool) {
	d, ok := e.defs[name]
	if !ok {
		return nil, false
	}
	return d.raw, true
}

func (e *Engine) canonical(name string) (string, bool) {
	if _, ok := e.defs[name]; ok {
		return name, true
	}
	c, ok := e.lower[strings.ToLower(name)]
	return c, ok
}

// ---------------------------------------------------------------- structure

// merge collects a definition's properties, filling the card schema's `{}`
// placeholders (and, where additional properties are allowed, missing
// properties) from allOf parents.
func (e *Engine) merge(name string) *merged {
	if m, ok := e.merged[name]; ok {
		return m
	}
	m := &merged{props: map[string]*Schema{}, required: map[string]bool{}}
	e.merged[name] = m // break cycles
	s := e.defs[name]
	if s == nil {
		return m
	}
	own := s
	if len(s.Properties) == 0 {
		// "shorthand or object": oneOf [{type: string, ...}, {type: object, properties}]
		if obj, alts := inlineObjectBranch(s); obj != nil {
			own = obj
			for _, a := range alts {
				e.shorthands[name] = append(e.shorthands[name], shorthandLabel(a))
			}
		}
	}
	for k, v := range own.Properties {
		m.props[k] = v
	}
	for _, r := range own.Required.Names {
		m.required[r] = true
	}
	closed := own.AddProps != nil && !own.AddProps.Allowed
	for _, parent := range s.AllOf {
		pn := parent.RefName()
		if pn == "" {
			continue
		}
		pm := e.merge(pn)
		for k, v := range pm.props {
			if cur, ok := m.props[k]; ok && cur.IsEmpty() {
				m.props[k] = v
			} else if !ok && !closed {
				m.props[k] = v
			}
		}
		for r := range pm.required {
			if _, ok := m.props[r]; ok {
				m.required[r] = true
			}
		}
	}
	return m
}

// inlineObjectBranch returns the single inline object branch of a union with
// no $ref branches, plus the other (shorthand) branches.
func inlineObjectBranch(s *Schema) (*Schema, []*Schema) {
	var obj *Schema
	var alts []*Schema
	for _, b := range branches(s) {
		switch {
		case branchRef(b) != "":
			return nil, nil
		case len(b.Properties) > 0:
			if obj != nil {
				return nil, nil
			}
			obj = b
		default:
			alts = append(alts, b)
		}
	}
	return obj, alts
}

func shorthandLabel(b *Schema) string {
	if b.hasConst {
		return fmt.Sprintf("%q", formatValue(b.Const))
	}
	if len(b.Type) > 0 {
		return "any " + strings.Join(b.Type, " | ")
	}
	return "other"
}

func branches(s *Schema) []*Schema {
	out := make([]*Schema, 0, len(s.OneOf)+len(s.AnyOf))
	out = append(out, s.OneOf...)
	return append(out, s.AnyOf...)
}

// shapeOf classifies a definition without building its node.
func (e *Engine) shapeOf(name string) model.Shape {
	s := e.defs[name]
	if s == nil {
		return model.Scalar
	}
	if len(e.merge(name).props) > 0 {
		return model.Object
	}
	bs := branches(s)
	if len(bs) > 0 {
		hasRef, hasExpr, allLiteral, anyLiteral := false, false, true, false
		for _, b := range bs {
			switch {
			case branchRef(b) != "":
				hasRef = true
				allLiteral = false
			case isExprBranch(b):
				hasExpr = true
				allLiteral = false
			case isPatternOnly(b):
			case literalValues(b) != nil:
				anyLiteral = true
			default:
				allLiteral = false
			}
		}
		switch {
		case hasRef && !hasExpr:
			return model.Union
		case allLiteral && anyLiteral:
			return model.Enum
		default:
			return model.Scalar
		}
	}
	if len(s.Enum) > 1 {
		return model.Enum
	}
	return model.Scalar
}

func (e *Engine) discValue(name string) string {
	ps := e.merge(name).props[e.opts.Discriminator]
	if ps == nil {
		return ""
	}
	if vals := literalValues(ps); len(vals) == 1 {
		return vals[0]
	}
	return ""
}

func (e *Engine) enumValues(name string) []model.EnumValue {
	s := e.defs[name]
	if s == nil {
		return nil
	}
	var out []model.EnumValue
	seen := map[string]bool{}
	add := func(v string, b *Schema) {
		if seen[v] {
			return
		}
		seen[v] = true
		out = append(out, model.EnumValue{Value: v, Description: e.enumDesc(name, v, b)})
	}
	for _, v := range literalValues(&Schema{Enum: s.Enum}) {
		add(v, nil)
	}
	for _, b := range branches(s) {
		for _, v := range literalValues(b) {
			add(v, b)
		}
	}
	return out
}

func enumStrings(vals []model.EnumValue) []string {
	out := make([]string, len(vals))
	for i, v := range vals {
		out[i] = v.Value
	}
	return out
}

// ---------------------------------------------------------------- nodes

func (e *Engine) node(name string) *model.Node {
	if n, ok := e.cache[name]; ok {
		return n
	}
	s := e.defs[name]
	n := &model.Node{Name: name, Label: name, Raw: s.raw}
	e.cache[name] = n
	n.Description, n.Docs = e.defDesc(name)
	n.Shape = e.shapeOf(name)
	switch n.Shape {
	case model.Object:
		m := e.merge(name)
		n.Discriminator = e.opts.Discriminator
		n.DiscriminatorValue = e.discValue(name)
		n.Fields = e.fields(name, m.props, m.required)
		n.Shorthands = e.shorthands[name]
	case model.Union:
		n.Discriminator = e.opts.Discriminator
		for _, b := range branches(s) {
			ref := branchRef(b)
			if ref == "" && len(b.Properties) > 0 {
				req := map[string]bool{}
				for _, r := range b.Required.Names {
					req[r] = true
				}
				value := ""
				if ps := b.Properties[e.opts.Discriminator]; ps != nil {
					if v := literalValues(ps); len(v) == 1 {
						value = v[0]
					}
				}
				n.Members = append(n.Members, model.Member{Name: "object", Value: value, Node: e.inline(b, name)})
				continue
			}
			if ref == "" {
				if isExprBranch(b) || isPatternOnly(b) {
					continue
				}
				n.MemberNotes = append(n.MemberNotes, e.typeOf(b, name).label)
				continue
			}
			ref = e.resolveName(ref)
			desc, _ := e.defDesc(ref)
			n.Members = append(n.Members, model.Member{
				Name:        ref,
				Value:       e.discValue(ref),
				Description: desc,
				Node:        e.lazy(ref),
			})
		}
	case model.Enum:
		n.EnumValues = e.enumValues(name)
	case model.Scalar:
		t := e.typeOf(s, name)
		n.Label = t.label
		if t.target != nil {
			// e.g. an expression wrapper around an enum: expose the enum values
			if inner := t.target(); inner != nil && inner.Shape == model.Enum {
				n.EnumValues = inner.EnumValues
			}
		}
	}
	return n
}

func (e *Engine) resolveName(ref string) string {
	if c, ok := e.canonical(ref); ok {
		return c
	}
	return ref
}

func (e *Engine) lazy(name string) func() *model.Node {
	return func() *model.Node {
		if _, ok := e.defs[name]; !ok {
			return nil
		}
		return e.node(name)
	}
}

// fields builds an object's fields: discriminator first, then alphabetical.
func (e *Engine) fields(owner string, props map[string]*Schema, required map[string]bool) []*model.Field {
	names := make([]string, 0, len(props))
	for k := range props {
		names = append(names, k)
	}
	disc := e.opts.Discriminator
	sort.Slice(names, func(i, j int) bool {
		if (names[i] == disc) != (names[j] == disc) {
			return names[i] == disc
		}
		return names[i] < names[j]
	})
	out := make([]*model.Field, 0, len(names))
	for _, p := range names {
		ps := props[p]
		t := e.typeOf(ps, owner)
		label := t.label
		if ps.Role == "expression" && !strings.Contains(label, "=expr") {
			label += " | =expr"
		}
		f := &model.Field{
			Name:        p,
			Label:       label,
			Required:    required[p] || ps.Required.Self,
			Enum:        t.enum,
			Target:      t.target,
			Description: e.fieldDesc(owner, p, ps),
			Raw:         ps.raw,
		}
		if t.hasConst {
			f.Const = t.constVal
			f.Enum = nil
		}
		if ps.hasDefault {
			f.Default, f.HasDefault = ps.Default, true
		}
		if ps.Version != "" {
			f.Notes = append(f.Notes, "since Adaptive Cards v"+ps.Version)
		}
		if ps.Format != "" {
			f.Notes = append(f.Notes, "format: "+ps.Format)
		}
		if ps.Pattern != "" && ps.Pattern != "=.*" {
			f.Notes = append(f.Notes, "pattern: "+ps.Pattern)
		}
		if len(ps.Examples) > 0 {
			ex := make([]string, len(ps.Examples))
			for i, x := range ps.Examples {
				ex[i] = formatValue(x)
			}
			f.Notes = append(f.Notes, "examples: "+strings.Join(ex, ", "))
		}
		out = append(out, f)
	}
	return out
}

// typeInfo is the rendered type of a schema fragment.
type typeInfo struct {
	label    string
	target   func() *model.Node
	enum     []string
	constVal string
	hasConst bool
}

func jsonType(v any) string {
	switch v.(type) {
	case bool:
		return "boolean"
	case float64, int:
		return "number"
	case nil:
		return "null"
	default:
		return "string"
	}
}

func (e *Engine) typeOf(s *Schema, owner string) typeInfo {
	if s == nil || s.IsEmpty() {
		return typeInfo{label: "any"}
	}
	if r := branchRef(s); r != "" {
		return e.refType(e.resolveName(r))
	}
	if s.hasConst {
		return typeInfo{label: jsonType(s.Const), constVal: formatValue(s.Const), hasConst: true}
	}
	if vals := literalValues(s); len(vals) > 0 {
		if len(vals) == 1 {
			return typeInfo{label: jsonType(s.Enum[0]), constVal: vals[0], hasConst: true}
		}
		return typeInfo{label: jsonType(s.Enum[0]), enum: vals}
	}
	if bs := branches(s); len(bs) > 0 {
		return e.unionType(bs, owner)
	}
	switch {
	case s.Type.has("array"):
		it := e.typeOf(s.Items, owner)
		return typeInfo{label: "[]" + it.label, target: it.target, enum: it.enum}
	case len(s.Properties) > 0:
		return typeInfo{label: "object", target: e.inline(s, owner)}
	case s.Type.has("object"):
		if s.AddProps != nil && s.AddProps.Schema != nil {
			it := e.typeOf(s.AddProps.Schema, owner)
			return typeInfo{label: "map[string]" + it.label, target: it.target}
		}
		return typeInfo{label: "object"}
	case len(s.Type) > 0:
		var ts []string
		for _, t := range s.Type {
			if t != "null" {
				ts = append(ts, t)
			}
		}
		if s.Pattern == "=.*" {
			return typeInfo{label: "=expr"}
		}
		return typeInfo{label: strings.Join(ts, " | ")}
	case s.Pattern == "=.*":
		return typeInfo{label: "=expr"}
	}
	return typeInfo{label: "any"}
}

func (e *Engine) refType(name string) typeInfo {
	if _, ok := e.defs[name]; !ok {
		return typeInfo{label: name}
	}
	switch e.shapeOf(name) {
	case model.Object, model.Union:
		return typeInfo{label: name, target: e.lazy(name)}
	case model.Enum:
		return typeInfo{label: name, target: e.lazy(name), enum: enumStrings(e.enumValues(name))}
	default:
		return e.typeOf(e.defs[name], name)
	}
}

func (e *Engine) unionType(bs []*Schema, owner string) typeInfo {
	var (
		labels  []string
		seen    = map[string]bool{}
		hasExpr bool
		lits    []string
		litType string
		out     typeInfo
	)
	for _, b := range bs {
		switch {
		case isExprBranch(b):
			hasExpr = true
			continue
		case isPatternOnly(b):
			continue
		case len(b.Type) == 1 && b.Type[0] == "null" && branchRef(b) == "":
			continue
		}
		if v := literalValues(b); v != nil && branchRef(b) == "" {
			lits = append(lits, v...)
			if litType == "" {
				if b.hasConst {
					litType = jsonType(b.Const)
				} else {
					litType = jsonType(b.Enum[0])
				}
			}
			continue
		}
		t := e.typeOf(b, owner)
		if !seen[t.label] {
			seen[t.label] = true
			labels = append(labels, t.label)
		}
		if out.target == nil && t.target != nil {
			out.target = t.target
		}
		if out.enum == nil {
			out.enum = t.enum
		}
	}
	if len(labels) == 0 && len(lits) > 0 {
		labels = []string{litType}
		out.enum = lits
	}
	out.label = strings.Join(labels, " | ")
	if hasExpr {
		if out.label == "" {
			out.label = "=expr"
		} else if !strings.Contains(out.label, "=expr") {
			out.label += " | =expr"
		}
	}
	if out.label == "" {
		out.label = "any"
	}
	return out
}

// inline wraps an anonymous object schema in a node.
func (e *Engine) inline(s *Schema, owner string) func() *model.Node {
	var n *model.Node
	return func() *model.Node {
		if n != nil {
			return n
		}
		req := map[string]bool{}
		for _, r := range s.Required.Names {
			req[r] = true
		}
		n = &model.Node{Label: "object", Shape: model.Object, Raw: s.raw}
		if !describe.IsPlaceholder(s.Description) {
			n.Description = s.Description
		}
		n.Fields = e.fields(owner, s.Properties, req)
		return n
	}
}

// ---------------------------------------------------------------- descriptions

func realText(s string) string {
	if describe.IsPlaceholder(s) {
		return ""
	}
	return strings.TrimSpace(s)
}

func (e *Engine) defDesc(name string) (desc, docs string) {
	if en, ok := e.opts.Overlay.Get(name, describe.StripNoKind(name)); ok {
		return en.Description, en.Docs
	}
	if s := e.defs[name]; s != nil {
		return realText(s.Description), ""
	}
	return "", ""
}

func (e *Engine) fieldKeys(owner, field string, ps *Schema) []string {
	keys := []string{owner + "." + field}
	for _, tok := range []string{ps.Description, ps.Title} {
		if o, f, ok := describe.TokenOwner(tok); ok {
			keys = append(keys, o+"."+f)
			break
		}
	}
	if base := describe.StripNoKind(owner); base != owner {
		keys = append(keys, base+"."+field)
	}
	return keys
}

func (e *Engine) fieldDesc(owner, field string, ps *Schema) string {
	if en, ok := e.opts.Overlay.Get(e.fieldKeys(owner, field, ps)...); ok {
		return en.Description
	}
	return realText(ps.Description)
}

func (e *Engine) enumDesc(enum, value string, b *Schema) string {
	if en, ok := e.opts.Overlay.Get("enum."+enum+"."+value, "enum."+describe.StripNoKind(enum)+"."+value); ok {
		return en.Description
	}
	if b != nil {
		return realText(b.Description)
	}
	return ""
}

// ---------------------------------------------------------------- overlay coverage

// longestDef splits key at the longest definition-name prefix followed by '.'.
func (e *Engine) longestDef(key string) (def, rest string, ok bool) {
	for i := len(key) - 1; i > 0; i-- {
		if key[i] != '.' {
			continue
		}
		if _, found := e.defs[key[:i]]; found {
			return key[:i], key[i+1:], true
		}
	}
	return "", "", false
}

// ValidKey reports whether an overlay key matches something in this schema.
func (e *Engine) ValidKey(key string) bool {
	if rest, ok := strings.CutPrefix(key, "enum."); ok {
		def, value, ok := e.longestDef(rest)
		if !ok {
			return false
		}
		for _, v := range e.enumValues(def) {
			if v.Value == value {
				return true
			}
		}
		return false
	}
	if _, ok := e.defs[key]; ok {
		return true
	}
	if e.owners[key] {
		return true
	}
	def, field, ok := e.longestDef(key)
	if !ok {
		return false
	}
	_, ok = e.merge(def).props[field]
	return ok
}

// Coverage summarises how much of the schema has descriptions.
type Coverage struct {
	Source                        string
	Definitions, DescribedDefs    int
	Kinds, DescribedKinds         int
	FieldKeys, DescribedFieldKeys int
	Stale                         []string
	UndescribedKinds              []string
}

// Coverage computes description coverage and stale overlay keys.
func (e *Engine) Coverage() Coverage {
	c := Coverage{Source: e.opts.Source}
	fieldSeen := map[string]bool{}
	for _, name := range e.names {
		desc, _ := e.defDesc(name)
		c.Definitions++
		if desc != "" {
			c.DescribedDefs++
		}
		if e.discValue(name) != "" && !strings.HasSuffix(name, "NoKind") {
			c.Kinds++
			if desc != "" {
				c.DescribedKinds++
			} else {
				c.UndescribedKinds = append(c.UndescribedKinds, name)
			}
		}
		for field, ps := range e.merge(name).props {
			if field == e.opts.Discriminator {
				continue
			}
			keys := e.fieldKeys(name, field, ps)
			key := keys[len(keys)-1]
			if len(keys) > 1 {
				key = keys[1] // the token owner (or NoKind base) is the shared key
			}
			if fieldSeen[key] {
				continue
			}
			fieldSeen[key] = true
			c.FieldKeys++
			if e.fieldDesc(name, field, ps) != "" {
				c.DescribedFieldKeys++
			}
		}
	}
	for _, k := range e.opts.Overlay.Keys() {
		if !e.ValidKey(k) {
			c.Stale = append(c.Stale, k)
		}
	}
	return c
}

// String renders a coverage line.
func (c Coverage) String() string {
	pct := func(a, b int) string {
		if b == 0 {
			return "n/a"
		}
		return fmt.Sprintf("%d/%d (%.0f%%)", a, b, 100*float64(a)/float64(b))
	}
	return fmt.Sprintf("%-6s kinds %s · definitions %s · fields %s",
		c.Source, pct(c.DescribedKinds, c.Kinds), pct(c.DescribedDefs, c.Definitions), pct(c.DescribedFieldKeys, c.FieldKeys))
}
