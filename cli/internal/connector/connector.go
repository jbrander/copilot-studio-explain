// Package connector explains the vendored connector definitions
// (reference/connectors/*.yml): connector -> operations/triggers -> inputs/outputs.
package connector

import (
	"fmt"
	"io/fs"
	"path"
	"sort"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"

	"copilot-studio-explain/cli/internal/describe"
	"copilot-studio-explain/cli/internal/model"
)

// Source is the connector catalog. Files are parsed lazily, one at a time.
type Source struct {
	fsys    fs.FS
	dir     string
	overlay describe.Section
	names   []string
	files   map[string]*lazyFile
}

type lazyFile struct {
	once sync.Once
	node *model.Node
	err  error
}

// New lists the connector files under dir in fsys.
func New(fsys fs.FS, dir string, overlay describe.Section) (*Source, error) {
	entries, err := fs.ReadDir(fsys, dir)
	if err != nil {
		return nil, fmt.Errorf("connectors: %w", err)
	}
	s := &Source{fsys: fsys, dir: dir, overlay: overlay, files: map[string]*lazyFile{}}
	for _, en := range entries {
		ext := path.Ext(en.Name())
		if en.IsDir() || (ext != ".yml" && ext != ".yaml") {
			continue
		}
		name := strings.TrimSuffix(en.Name(), ext)
		s.names = append(s.names, name)
		s.files[name] = &lazyFile{}
	}
	sort.Strings(s.names)
	return s, nil
}

// Name is the source name.
func (s *Source) Name() string { return "connector" }

// Names lists connector names (file stems such as shared_office365).
func (s *Source) Names() []string { return s.names }

// Lookup finds a connector by file stem, case-insensitively; the "shared_"
// prefix is optional.
func (s *Source) Lookup(name string) (*model.Node, bool) {
	for _, n := range s.names {
		if strings.EqualFold(n, name) || strings.EqualFold(strings.TrimPrefix(n, "shared_"), name) {
			node, err := s.load(n)
			if err != nil {
				return &model.Node{Name: n, Label: "ConnectorDefinition", Shape: model.Object,
					Description: "Failed to load connector: " + err.Error()}, true
			}
			return node, true
		}
	}
	return nil, false
}

func (s *Source) load(name string) (*model.Node, error) {
	lf := s.files[name]
	lf.once.Do(func() {
		var b []byte
		for _, ext := range []string{".yml", ".yaml"} {
			if b, lf.err = fs.ReadFile(s.fsys, path.Join(s.dir, name+ext)); lf.err == nil {
				break
			}
		}
		if lf.err != nil {
			return
		}
		var f file
		if lf.err = yaml.Unmarshal(b, &f); lf.err != nil {
			lf.err = fmt.Errorf("%s: %w", name, lf.err)
			return
		}
		lf.node = s.connectorNode(name, &f)
	})
	return lf.node, lf.err
}

// ---------------------------------------------------------------- YAML shapes

type file struct {
	Kind          string      `yaml:"kind"`
	ConnectorID   string      `yaml:"connectorId"`
	DisplayName   string      `yaml:"displayName"`
	Description   string      `yaml:"description"`
	ConnectorType string      `yaml:"connectorType"`
	IsSSO         *bool       `yaml:"isSSOSupported"`
	Operations    []operation `yaml:"operations"`
	Triggers      []operation `yaml:"triggers"`
}

type operation struct {
	OperationID string    `yaml:"operationId"`
	DisplayName string    `yaml:"displayName"`
	Description string    `yaml:"description"`
	InputType   *dataType `yaml:"inputType"`
	OutputType  *dataType `yaml:"outputType"`
	raw         yaml.Node
}

func (o *operation) UnmarshalYAML(n *yaml.Node) error {
	type plain operation
	var p plain
	if err := n.Decode(&p); err != nil {
		return err
	}
	*o = operation(p)
	o.raw = *n
	return nil
}

// dataType is either a scalar ("String") or {kind: Record|Table, properties: {...}}.
type dataType struct {
	Scalar     string
	Kind       string
	Properties map[string]*property
	raw        yaml.Node
}

func (d *dataType) UnmarshalYAML(n *yaml.Node) error {
	d.raw = *n
	if n.Kind == yaml.ScalarNode {
		d.Scalar = n.Value
		return nil
	}
	var p struct {
		Kind       string               `yaml:"kind"`
		Properties map[string]*property `yaml:"properties"`
	}
	if err := n.Decode(&p); err != nil {
		return err
	}
	d.Kind, d.Properties = p.Kind, p.Properties
	return nil
}

type property struct {
	DisplayName      string    `yaml:"displayName"`
	Description      string    `yaml:"description"`
	IsRequired       bool      `yaml:"isRequired"`
	Order            int       `yaml:"order"`
	Type             *dataType `yaml:"type"`
	Visibility       string    `yaml:"visibility"`
	AISensitivity    string    `yaml:"aiSensitivity"`
	DynamicValues    *struct{} `yaml:"dynamicValuesConfig"`
	EnumValues       []any     `yaml:"enumValues"`
	EnumDisplayNames []string  `yaml:"enumDisplayNames"`
	raw              yaml.Node
}

// UnmarshalYAML accepts the shorthand `name: String`.
func (p *property) UnmarshalYAML(n *yaml.Node) error {
	p.raw = *n
	if n.Kind == yaml.ScalarNode {
		p.Type = &dataType{Scalar: n.Value, raw: *n}
		return nil
	}
	type plain property
	var q plain
	if err := n.Decode(&q); err != nil {
		return err
	}
	raw := p.raw
	*p = property(q)
	p.raw = raw
	return nil
}

// decodeRaw returns a lazy decoder for a YAML subtree (used only by -o json).
func decodeRaw(n *yaml.Node) func() any {
	return func() any {
		if n == nil || n.Kind == 0 {
			return nil
		}
		var v any
		if err := n.Decode(&v); err != nil {
			return nil
		}
		return v
	}
}

// ---------------------------------------------------------------- nodes

func (s *Source) connectorNode(name string, f *file) *model.Node {
	n := &model.Node{
		Name:               name,
		Label:              "ConnectorDefinition",
		Shape:              model.Object,
		Discriminator:      "kind",
		DiscriminatorValue: f.Kind,
	}
	desc := f.DisplayName
	if f.Description != "" {
		desc += ": " + f.Description
	}
	if en, ok := s.overlay.Get(name); ok {
		desc, n.Docs = en.Description, en.Docs
	}
	facts := []string{"connectorId: " + f.ConnectorID}
	if f.ConnectorType != "" {
		facts = append(facts, "connectorType: "+f.ConnectorType)
	}
	if f.IsSSO != nil {
		facts = append(facts, fmt.Sprintf("isSSOSupported: %v", *f.IsSSO))
	}
	facts = append(facts, fmt.Sprintf("%d operations, %d triggers", len(f.Operations), len(f.Triggers)))
	n.Description = desc + "\n\n" + strings.Join(facts, "\n")
	n.Raw = map[string]any{
		"kind": f.Kind, "connectorId": f.ConnectorID, "displayName": f.DisplayName,
		"description": f.Description, "connectorType": f.ConnectorType,
		"operations": opIDs(f.Operations), "triggers": opIDs(f.Triggers),
	}
	add := func(ops []operation, label string) {
		for i := range ops {
			op := &ops[i]
			fd := &model.Field{
				Name:        op.OperationID,
				Label:       label,
				Description: s.opDesc(name, op),
				Raw:         decodeRaw(&op.raw),
			}
			if op.DisplayName != "" {
				fd.Notes = append(fd.Notes, "display name: "+op.DisplayName)
			}
			fd.Target = s.lazyOp(name, op, label)
			n.Fields = append(n.Fields, fd)
		}
	}
	add(f.Operations, "Operation")
	add(f.Triggers, "Trigger")
	return n
}

func opIDs(ops []operation) []string {
	out := make([]string, len(ops))
	for i, o := range ops {
		out[i] = o.OperationID
	}
	return out
}

func (s *Source) opDesc(connector string, op *operation) string {
	if en, ok := s.overlay.Get(connector + "." + op.OperationID); ok {
		return en.Description
	}
	switch {
	case op.Description != "" && op.DisplayName != "" && !strings.EqualFold(strings.TrimSuffix(op.Description, "."), op.DisplayName):
		return op.DisplayName + ": " + op.Description
	case op.Description != "":
		return op.Description
	default:
		return op.DisplayName
	}
}

func (s *Source) lazyOp(connector string, op *operation, label string) func() *model.Node {
	var n *model.Node
	return func() *model.Node {
		if n != nil {
			return n
		}
		n = &model.Node{
			Name:  op.OperationID,
			Label: label,
			Shape: model.Object,
			Description: s.opDesc(connector, op) + "\n\n" + fmt.Sprintf(
				"Call it from a topic node (InvokeConnectorAction) or a tool (InvokeConnectorTaskAction) with `operationId: %s`; "+
					"the inputs below are bound by name.", op.OperationID),
			Raw: decodeRaw(&op.raw),
		}
		if op.InputType != nil {
			n.Fields = append(n.Fields, s.typeField("inputs", op.InputType, "The operation's input parameters."))
		}
		if op.OutputType != nil {
			n.Fields = append(n.Fields, s.typeField("outputs", op.OutputType, "What the operation returns."))
		}
		return n
	}
}

func (s *Source) typeField(name string, t *dataType, desc string) *model.Field {
	label, target := s.typeOf(t)
	return &model.Field{Name: name, Label: label, Description: desc, Target: target, Raw: decodeRaw(&t.raw)}
}

func scalarLabel(s string) string {
	switch s {
	case "String", "Boolean", "Number", "Integer", "Date", "DateTime", "File", "Any":
		return strings.ToLower(s)
	default:
		return s
	}
}

func (s *Source) typeOf(t *dataType) (string, func() *model.Node) {
	if t == nil {
		return "any", nil
	}
	if t.Scalar != "" {
		return scalarLabel(t.Scalar), nil
	}
	label := "Record"
	if t.Kind == "Table" {
		label = "[]Record"
	}
	if len(t.Properties) == 0 {
		return label, nil
	}
	var n *model.Node
	return label, func() *model.Node {
		if n == nil {
			n = s.recordNode(t)
		}
		return n
	}
}

func (s *Source) recordNode(t *dataType) *model.Node {
	n := &model.Node{Label: "Record", Shape: model.Object, Raw: decodeRaw(&t.raw)}
	if t.Kind == "Table" {
		n.Description = "Each row of the table has these columns."
	}
	names := make([]string, 0, len(t.Properties))
	for k := range t.Properties {
		names = append(names, k)
	}
	sort.Slice(names, func(i, j int) bool {
		pi, pj := t.Properties[names[i]], t.Properties[names[j]]
		if pi.Order != pj.Order {
			return pi.Order < pj.Order
		}
		return names[i] < names[j]
	})
	for _, k := range names {
		p := t.Properties[k]
		label, target := s.typeOf(p.Type)
		f := &model.Field{
			Name:        k,
			Label:       label,
			Required:    p.IsRequired,
			Description: strings.TrimSpace(p.Description),
			Target:      target,
			Raw:         decodeRaw(&p.raw),
		}
		if p.DisplayName != "" {
			f.Notes = append(f.Notes, "display name: "+p.DisplayName)
		}
		if p.Visibility != "" {
			f.Notes = append(f.Notes, "visibility: "+p.Visibility)
		}
		if p.AISensitivity != "" {
			f.Notes = append(f.Notes, "aiSensitivity: "+p.AISensitivity)
		}
		if p.DynamicValues != nil {
			f.Notes = append(f.Notes, "values are looked up dynamically from the connector")
		}
		for i, v := range p.EnumValues {
			ev := fmt.Sprint(v)
			if i < len(p.EnumDisplayNames) && p.EnumDisplayNames[i] != "" && p.EnumDisplayNames[i] != ev {
				ev += " (" + p.EnumDisplayNames[i] + ")"
			}
			f.Enum = append(f.Enum, ev)
		}
		n.Fields = append(n.Fields, f)
	}
	return n
}

// ValidKey reports whether an overlay key names a connector or connector.operation.
func (s *Source) ValidKey(key string) bool {
	name, op, _ := strings.Cut(key, ".")
	n, ok := s.Lookup(name)
	if !ok {
		return false
	}
	if op == "" {
		return true
	}
	for _, f := range n.Fields {
		if f.Name == op {
			return true
		}
	}
	return false
}
