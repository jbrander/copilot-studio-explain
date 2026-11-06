// Package render prints resolved resources in kubectl-explain style.
package render

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"copilot-studio-explain/cli/internal/explain"
	"copilot-studio-explain/cli/internal/model"
)

// Options controls plaintext output.
type Options struct {
	Recursive bool
	MaxDepth  int // recursive mode only; <= 0 means unlimited
	Width     int // wrap width; <= 0 means 80
}

func (o Options) width() int {
	if o.Width <= 0 {
		return 80
	}
	return o.Width
}

// Plain writes the kubectl-style plaintext explanation.
func Plain(w io.Writer, r *explain.Result, version string, opts Options) error {
	p := &printer{w: w, opts: opts, expr: r.Expr()}
	p.header(r, version)
	p.description(r)
	if r.Field != nil && r.Member == "" {
		p.fieldFacts(r.Field, r.Node != nil && len(r.Node.EnumValues) > 0)
	}
	n := r.Node
	if n != nil {
		switch n.Shape {
		case model.Union:
			p.union(n)
		case model.Enum:
			p.enum(n.EnumValues)
		case model.Scalar:
			if len(n.EnumValues) > 0 {
				p.enum(n.EnumValues)
			}
		case model.Object:
			if len(n.Fields) == 0 {
				p.line("FIELDS:")
				p.line("  <none>")
			} else if opts.Recursive {
				p.line("FIELDS:")
				p.recursive(n, 1, map[*model.Node]bool{n: true})
			} else {
				p.fields(n)
			}
		}
	}
	return p.err
}

type printer struct {
	w    io.Writer
	opts Options
	expr string // the resolved resource expression, for hints
	err  error
}

func (p *printer) line(format string, a ...any) {
	if p.err != nil {
		return
	}
	_, p.err = fmt.Fprintf(p.w, format+"\n", a...)
}

func (p *printer) blank() { p.line("") }

func (p *printer) header(r *explain.Result, version string) {
	p.line("KIND:       %s", r.Root.Name)
	p.line("SCHEMA:     %s (%s)", r.Source, version)
	if len(r.Path) > 0 {
		p.blank()
		label := "?"
		switch {
		case r.Member != "" && r.Node != nil:
			label = r.Node.Label
		case r.Field != nil:
			label = r.Field.Label
		}
		suffix := ""
		if r.Field != nil && r.Field.Required && r.Member == "" {
			suffix = " -required-"
		}
		p.line("FIELD:      %s <%s>%s", r.PathString(), label, suffix)
	}
	p.blank()
}

func (p *printer) description(r *explain.Result) {
	p.line("DESCRIPTION:")
	var parts []string
	var docs []string
	n := r.Node
	if r.Field != nil && r.Member == "" {
		if r.Field.Description != "" {
			parts = append(parts, r.Field.Description)
		}
		// Also describe the field's named type, when it has its own text.
		if n != nil && n.Name != "" && n.Description != "" && n.Description != r.Field.Description {
			parts = append(parts, n.Name+": "+n.Description)
		}
	} else if n != nil && n.Description != "" {
		parts = append(parts, n.Description)
	}
	if n != nil && n.Docs != "" {
		docs = append(docs, "See: "+n.Docs)
	}
	if len(parts) == 0 {
		parts = append(parts, "No description in schema.")
	}
	if n != nil && n.Shape == model.Object && n.DiscriminatorValue != "" && n.Discriminator != "" {
		parts = append(parts, fmt.Sprintf("YAML: `%s: %s`", n.Discriminator, n.DiscriminatorValue))
	}
	if n != nil && len(n.Shorthands) > 0 {
		parts = append(parts, "Shorthand: instead of an object you can also write "+strings.Join(n.Shorthands, " or ")+".")
	}
	p.wrapped(strings.Join(parts, "\n\n"), "    ")
	for _, d := range docs {
		p.line("    %s", d)
	}
	p.blank()
}

func (p *printer) fieldFacts(f *model.Field, enumShown bool) {
	g := *f
	if enumShown {
		g.Enum = nil
	}
	facts := factLines(&g)
	if len(facts) == 0 {
		return
	}
	for _, l := range facts {
		p.wrapped(l, "    ")
	}
	p.blank()
}

// factLines are the extra one-liners shown under a field.
func factLines(f *model.Field) []string {
	var out []string
	if f.Const != "" && !strings.Contains(f.Description, "Must be") {
		out = append(out, fmt.Sprintf("Must be %q.", f.Const))
	}
	if len(f.Enum) > 0 {
		out = append(out, "enum: "+strings.Join(f.Enum, ", "))
	}
	if f.HasDefault {
		out = append(out, "default: "+formatDefault(f.Default))
	}
	out = append(out, f.Notes...)
	return out
}

func formatDefault(v any) string {
	switch x := v.(type) {
	case string:
		return fmt.Sprintf("%q", x)
	case nil:
		return "null"
	default:
		return fmt.Sprint(x)
	}
}

func (p *printer) union(n *model.Node) {
	p.line("ONE OF (set `%s:`):", n.Discriminator)
	tw := tabwriter.NewWriter(p.w, 0, 0, 3, ' ', 0)
	for _, m := range n.Members {
		name := m.Name
		switch {
		case m.Value == "":
			name += " (no " + n.Discriminator + ")"
		case m.Value != m.Name:
			name += " (" + n.Discriminator + ": " + m.Value + ")"
		}
		row(tw, name, FirstSentence(m.Description, 90))
	}
	if p.err == nil {
		p.err = tw.Flush()
	}
	if len(n.MemberNotes) > 0 {
		p.line("  also accepts: %s", strings.Join(n.MemberNotes, ", "))
	}
	p.blank()
	p.line("Use `cs-explain %s.<%s>` to explain one of them.", p.expr, titleCase(n.Discriminator))
}

// row writes a two-column table row without trailing padding when the second column is empty.
func row(w io.Writer, a, b string) {
	if b == "" {
		fmt.Fprintf(w, "  %s\n", a)
		return
	}
	fmt.Fprintf(w, "  %s\t%s\n", a, b)
}

func titleCase(s string) string {
	if s == "" {
		return "Kind"
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func (p *printer) enum(vals []model.EnumValue) {
	p.line("ENUM:")
	tw := tabwriter.NewWriter(p.w, 0, 0, 3, ' ', 0)
	for _, v := range vals {
		row(tw, v.Value, FirstSentence(v.Description, 100))
	}
	if p.err == nil {
		p.err = tw.Flush()
	}
}

func typeSuffix(f *model.Field) string {
	s := ""
	if f.Required {
		s = " -required-"
	}
	return s
}

func (p *printer) fields(n *model.Node) {
	p.line("FIELDS:")
	for i, f := range n.Fields {
		if i > 0 {
			p.blank()
		}
		p.line("  %s\t<%s>%s", f.Name, f.Label, typeSuffix(f))
		if f.Description != "" {
			p.wrapped(f.Description, "    ")
		}
		for _, l := range factLines(f) {
			p.wrapped(l, "    ")
		}
		if c := f.Child(); c != nil && c.Shape == model.Union {
			p.line("    one of %d kinds: cs-explain %s.%s", len(c.Members), p.expr, f.Name)
		}
	}
}

func (p *printer) recursive(n *model.Node, depth int, seen map[*model.Node]bool) {
	indent := strings.Repeat("  ", depth)
	for _, f := range n.Fields {
		line := fmt.Sprintf("%s%s\t<%s>%s", indent, f.Name, f.Label, typeSuffix(f))
		c := f.Child()
		expand := false
		switch {
		case c == nil:
		case c.Shape == model.Union:
			line += fmt.Sprintf("  # one of %d kinds", len(c.Members))
		case c.Shape == model.Object && seen[c]:
			line += "  # (recursive)"
		case c.Shape == model.Object && p.opts.MaxDepth > 0 && depth >= p.opts.MaxDepth:
			if len(c.Fields) > 0 {
				line += "  # ..."
			}
		case c.Shape == model.Object:
			expand = true
		}
		p.line("%s", line)
		if expand {
			seen[c] = true
			p.recursive(c, depth+1, seen)
			delete(seen, c)
		}
	}
}

// FirstSentence trims a description to its first sentence (or max runes).
func FirstSentence(s string, max int) string {
	s = strings.TrimSpace(strings.SplitN(s, "\n", 2)[0])
	s = cutSentence(s)
	if r := []rune(s); len(r) > max {
		s = string(r[:max-1]) + "…"
	}
	return s
}

// cutSentence keeps text up to the first ". " that isn't an abbreviation.
func cutSentence(s string) string {
	from := 0
	for {
		i := strings.Index(s[from:], ". ")
		if i < 0 {
			return s
		}
		i += from
		word := s[strings.LastIndexAny(s[:i], " (")+1 : i+1]
		switch strings.ToLower(word) {
		case "i.e.", "e.g.", "etc.", "vs.":
			from = i + 2
			continue
		}
		return s[:i+1]
	}
}

func (p *printer) wrapped(text, indent string) {
	for _, l := range wrap(text, indent, p.opts.width()) {
		p.line("%s", l)
	}
}

// wrap word-wraps text, keeping explicit newlines.
func wrap(text, indent string, width int) []string {
	var out []string
	for _, para := range strings.Split(text, "\n") {
		words := strings.Fields(para)
		if len(words) == 0 {
			out = append(out, "")
			continue
		}
		cur := indent + words[0]
		for _, w := range words[1:] {
			if len(cur)+1+len(w) > width {
				out = append(out, cur)
				cur = indent + w
			} else {
				cur += " " + w
			}
		}
		out = append(out, cur)
	}
	return out
}
