package explain

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"copilot-studio-explain/cli/internal/model"
)

// Result is a resolved resource path.
type Result struct {
	Source string
	Root   *model.Node
	// Path holds the canonical segments after the root (field names and union member names).
	Path []string
	// Field is the last field traversed, nil when the path is just the root.
	Field *model.Field
	// Member is set when the last segment selected a union member.
	Member string
	// Node is what to describe; nil when the path ends on a leaf field.
	Node *model.Node
}

// Expr is the canonical resource expression, e.g. "card:Input.ChoiceSet.choices".
func (r *Result) Expr() string {
	e := r.Root.Name
	if len(r.Path) > 0 {
		e += "." + r.PathString()
	}
	if r.Source != Bot {
		e = r.Source + ":" + e
	}
	return e
}

// PathString is the dotted path after the root.
func (r *Result) PathString() string { return strings.Join(r.Path, ".") }

// NotFoundError is a path segment that matched nothing.
type NotFoundError struct {
	Msg         string
	Suggestions []string
	Choices     []string
}

func (e *NotFoundError) Error() string {
	var b strings.Builder
	b.WriteString(e.Msg)
	if len(e.Suggestions) > 0 {
		fmt.Fprintf(&b, "\nDid you mean: %s?", strings.Join(e.Suggestions, ", "))
	}
	if n := len(e.Choices); n > 0 && n <= 60 {
		fmt.Fprintf(&b, "\nValid choices: %s", strings.Join(e.Choices, ", "))
	}
	return b.String()
}

// Resolve walks a resource expression such as "SendActivity.activity" or
// "card:Input.ChoiceSet.choices".
func (c *Catalog) Resolve(expr string) (*Result, error) {
	srcName, rest := SplitSource(strings.TrimSpace(expr))
	src, err := c.Source(srcName)
	if err != nil {
		return nil, err
	}
	rest = strings.Trim(rest, ".")
	if rest == "" {
		return nil, fmt.Errorf("you must name a resource to explain; try `cse list %s`", srcName)
	}
	segs := strings.Split(rest, ".")
	res := &Result{Source: srcName}

	// Root: the longest dotted prefix that names a definition (card types
	// such as Input.ChoiceSet contain dots).
	for k := len(segs); k >= 1; k-- {
		if n, ok := src.Lookup(strings.Join(segs[:k], ".")); ok {
			res.Root, segs = n, segs[k:]
			break
		}
	}
	if res.Root == nil {
		// Suggest against every dotted prefix: card names contain dots.
		var sugg []string
		for k := len(segs); k >= 1 && len(sugg) < 5; k-- {
			for _, s := range Suggest(strings.Join(segs[:k], "."), src.Names(), 5) {
				if !slices.Contains(sugg, s) && len(sugg) < 5 {
					sugg = append(sugg, s)
				}
			}
		}
		return nil, &NotFoundError{
			Msg:         fmt.Sprintf("%s: no %s definition matches %q (see `cse list %s --all`)", srcName, srcName, rest, srcName),
			Suggestions: sugg,
		}
	}
	cur := res.Root
	res.Node = cur
	for len(segs) > 0 {
		if cur == nil {
			return nil, fmt.Errorf("%s is a leaf of type <%s>; it has no field %q",
				res.describePos(), res.Field.Label, segs[0])
		}
		switch cur.Shape {
		case model.Union:
			m, used := matchMember(cur, segs)
			if m == nil {
				names := memberNames(cur)
				return nil, &NotFoundError{
					Msg: fmt.Sprintf("%s <%s> is a union: the next segment must be one of its %d kinds (set by `%s:`), not %q",
						res.describePos(), cur.Label, len(cur.Members), cur.Discriminator, segs[0]),
					Suggestions: Suggest(segs[0], names, 5),
					Choices:     names,
				}
			}
			segs = segs[used:]
			res.Path = append(res.Path, m.Name)
			res.Member = m.Name
			cur = m.Node()
		case model.Object:
			f, used := matchField(cur, segs)
			if f == nil {
				names := fieldNames(cur)
				return nil, &NotFoundError{
					Msg:         fmt.Sprintf("%s <%s> has no field %q", res.describePos(), cur.Label, segs[0]),
					Suggestions: Suggest(segs[0], names, 5),
					Choices:     names,
				}
			}
			segs = segs[used:]
			res.Path = append(res.Path, f.Name)
			res.Field, res.Member = f, ""
			cur = f.Child()
		default:
			return nil, fmt.Errorf("%s <%s> is %s and has no fields (got %q)",
				res.describePos(), cur.Label, article(cur.Shape), segs[0])
		}
		res.Node = cur
	}
	return res, nil
}

func article(s model.Shape) string {
	if s == model.Enum {
		return "an enum"
	}
	return "a scalar"
}

func (r *Result) describePos() string {
	if len(r.Path) == 0 {
		return r.Root.Name
	}
	return r.Root.Name + "." + r.PathString()
}

// matchField finds the field named by the longest prefix of segs (card
// property names such as "choices.data" contain dots). Exact matches win over
// case-insensitive ones.
func matchField(n *model.Node, segs []string) (*model.Field, int) {
	for k := len(segs); k >= 1; k-- {
		name := strings.Join(segs[:k], ".")
		for _, f := range n.Fields {
			if f.Name == name {
				return f, k
			}
		}
		for _, f := range n.Fields {
			if strings.EqualFold(f.Name, name) {
				return f, k
			}
		}
	}
	return nil, 0
}

func matchMember(n *model.Node, segs []string) (*model.Member, int) {
	for k := len(segs); k >= 1; k-- {
		name := strings.Join(segs[:k], ".")
		for i := range n.Members {
			m := &n.Members[i]
			if strings.EqualFold(m.Name, name) || (m.Value != "" && strings.EqualFold(m.Value, name)) {
				return m, k
			}
		}
	}
	return nil, 0
}

func fieldNames(n *model.Node) []string {
	out := make([]string, len(n.Fields))
	for i, f := range n.Fields {
		out[i] = f.Name
	}
	return out
}

func memberNames(n *model.Node) []string {
	out := make([]string, len(n.Members))
	for i, m := range n.Members {
		out[i] = m.Name
	}
	return out
}

// Children lists the next path segments available below a node: field names
// for objects, member names for unions.
func Children(n *model.Node) []string {
	if n == nil {
		return nil
	}
	switch n.Shape {
	case model.Object:
		return fieldNames(n)
	case model.Union:
		return memberNames(n)
	}
	return nil
}

// Suggest returns up to max candidates close to s: case-insensitive prefix or
// substring matches first, then small edit distances.
func Suggest(s string, candidates []string, max int) []string {
	ls := strings.ToLower(s)
	type scored struct {
		name  string
		score int
	}
	var out []scored
	for _, c := range candidates {
		lc := strings.ToLower(c)
		switch {
		case lc == ls:
			out = append(out, scored{c, -2})
		case strings.HasPrefix(lc, ls) || strings.HasPrefix(ls, lc):
			out = append(out, scored{c, -1})
		case len(ls) >= 4 && strings.Contains(lc, ls):
			out = append(out, scored{c, 0})
		default:
			if d := levenshtein(ls, lc); d <= maxInt(2, len(ls)/3) {
				out = append(out, scored{c, d})
			}
		}
	}
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].score != out[j].score {
			return out[i].score < out[j].score
		}
		if len(out[i].name) != len(out[j].name) {
			return len(out[i].name) < len(out[j].name)
		}
		return out[i].name < out[j].name
	})
	if len(out) > max {
		out = out[:max]
	}
	names := make([]string, len(out))
	for i, o := range out {
		names[i] = o.name
	}
	return names
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func levenshtein(a, b string) int {
	ra, rb := []rune(a), []rune(b)
	prev := make([]int, len(rb)+1)
	cur := make([]int, len(rb)+1)
	for j := range prev {
		prev[j] = j
	}
	for i := 1; i <= len(ra); i++ {
		cur[0] = i
		for j := 1; j <= len(rb); j++ {
			cost := 1
			if ra[i-1] == rb[j-1] {
				cost = 0
			}
			cur[j] = min(prev[j]+1, cur[j-1]+1, prev[j-1]+cost)
		}
		prev, cur = cur, prev
	}
	return prev[len(rb)]
}
