package render

import (
	"encoding/json"
	"io"
	"regexp"
	"sort"

	"copilot-studio-explain/cli/internal/explain"
	"copilot-studio-explain/cli/internal/model"
)

var refRE = regexp.MustCompile(`"\$ref"\s*:\s*"#/definitions/([^"]+)"`)

type jsonDoc struct {
	Schema     string                     `json:"schema"`
	Version    string                     `json:"version"`
	Resource   string                     `json:"resource"`
	Path       string                     `json:"path,omitempty"`
	Type       string                     `json:"type"`
	Shape      string                     `json:"shape,omitempty"`
	Required   bool                       `json:"required,omitempty"`
	Definition json.RawMessage            `json:"definition"`
	Referenced map[string]json.RawMessage `json:"referenced,omitempty"`
}

// JSON writes the raw source fragment for the resolved node, plus the
// definitions it references directly (one level of $ref resolution).
func JSON(w io.Writer, r *explain.Result, version string, src explain.Source) error {
	doc := jsonDoc{
		Schema:   r.Source,
		Version:  version,
		Resource: r.Root.Name,
		Path:     r.PathString(),
	}
	switch {
	case r.Field != nil && r.Member == "":
		doc.Type = r.Field.Label
		doc.Required = r.Field.Required
		doc.Definition = model.RawJSON(r.Field.Raw)
	case r.Node != nil:
		doc.Type = r.Node.Label
		doc.Definition = model.RawJSON(r.Node.Raw)
	}
	if r.Node != nil {
		doc.Shape = r.Node.Shape.String()
	}
	if rs, ok := src.(explain.RawSource); ok {
		var names []string
		seen := map[string]bool{r.Root.Name: len(r.Path) == 0}
		for _, m := range refRE.FindAllSubmatch(doc.Definition, -1) {
			if n := string(m[1]); !seen[n] {
				seen[n] = true
				names = append(names, n)
			}
		}
		sort.Strings(names)
		for _, n := range names {
			if raw, ok := rs.Definition(n); ok {
				if doc.Referenced == nil {
					doc.Referenced = map[string]json.RawMessage{}
				}
				doc.Referenced[n] = raw
			}
		}
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	return enc.Encode(doc)
}
