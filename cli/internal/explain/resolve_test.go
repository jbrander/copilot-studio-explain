package explain_test

import (
	"errors"
	"io"
	"slices"
	"strings"
	"testing"

	"copilot-studio-explain/cli/internal/explain"
	"copilot-studio-explain/cli/internal/model"
	"copilot-studio-explain/cli/internal/render"
)

func catalog(t *testing.T) *explain.Catalog {
	t.Helper()
	c, err := explain.Open("")
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestResolve(t *testing.T) {
	c := catalog(t)
	cases := []struct {
		expr, wantExpr, wantLabel string
		wantShape                 model.Shape
	}{
		{"SendActivity", "SendActivity", "SendActivity", model.Object},
		{"sendactivity.ACTIVITY", "SendActivity.activity", "ActivityTemplateBase", model.Union},
		{"AdaptiveDialog.beginDialog.OnRecognizedIntent.actions", "AdaptiveDialog.beginDialog.OnRecognizedIntent.actions", "DialogAction", model.Union},
		{"AdaptiveDialog.beginDialog.OnRecognizedIntent.actions.SendActivity.activity", "AdaptiveDialog.beginDialog.OnRecognizedIntent.actions.SendActivity.activity", "ActivityTemplateBase", model.Union},
		{"SendActivity.activity.Message", "SendActivity.activity.MessageOptionalKind", "MessageOptionalKind", model.Object}, // by kind value
		{"ConversationOutcome", "ConversationOutcome", "ConversationOutcome", model.Enum},
		{"BooleanPrebuiltEntity", "BooleanPrebuiltEntity", "BooleanPrebuiltEntity", model.Object},
		{"card:Input.ChoiceSet", "card:Input.ChoiceSet", "Input.ChoiceSet", model.Object},
		{"card:Input.ChoiceSet.choices.data", "card:Input.ChoiceSet.choices.data", "Data.Query", model.Object},
		{"card:AdaptiveCard.body.Input.Text", "card:AdaptiveCard.body.Input.Text", "Input.Text", model.Object},
		{"cards:TextBlock.size", "card:TextBlock.size", "FontSize", model.Enum},
		{"connector:office365users", "connector:shared_office365users", "ConnectorDefinition", model.Object},
		{"connector:shared_office365users.MyProfile_V2.outputs", "connector:shared_office365users.MyProfile_V2.outputs", "Record", model.Object},
	}
	for _, tc := range cases {
		t.Run(tc.expr, func(t *testing.T) {
			r, err := c.Resolve(tc.expr)
			if err != nil {
				t.Fatal(err)
			}
			if got := r.Expr(); got != tc.wantExpr {
				t.Errorf("Expr = %q, want %q", got, tc.wantExpr)
			}
			if r.Node == nil {
				t.Fatalf("Node is nil")
			}
			if r.Node.Label != tc.wantLabel || r.Node.Shape != tc.wantShape {
				t.Errorf("node = <%s> %s, want <%s> %s", r.Node.Label, r.Node.Shape, tc.wantLabel, tc.wantShape)
			}
		})
	}
}

func TestResolveLeaf(t *testing.T) {
	r, err := catalog(t).Resolve("SendActivity.disabled")
	if err != nil {
		t.Fatal(err)
	}
	if r.Node != nil || r.Field == nil || r.Field.Label != "boolean" || !r.Field.HasDefault {
		t.Fatalf("unexpected leaf result: node=%v field=%+v", r.Node, r.Field)
	}
}

func TestResolveErrors(t *testing.T) {
	c := catalog(t)
	cases := []struct {
		expr, wantMsg, wantSuggestion string
	}{
		{"SendActvity", "no bot definition matches", "SendActivity"},
		{"SendActivity.activty", `has no field "activty"`, "activity"},
		{"AdaptiveDialog.beginDialog.OnRecognisedIntent", "is a union", "OnRecognizedIntent"},
		{"SendActivity.id.foo", "is a leaf", ""},
		{"ConversationOutcome.foo", "has no fields", ""},
		{"card:Input.ChoiseSet", "no card definition matches", "Input.ChoiceSet"},
		{"", "must name a resource", ""},
		{"connector:", "must name a resource", ""},
	}
	for _, tc := range cases {
		t.Run(tc.expr, func(t *testing.T) {
			_, err := c.Resolve(tc.expr)
			if err == nil {
				t.Fatal("expected an error")
			}
			if !strings.Contains(err.Error(), tc.wantMsg) {
				t.Errorf("error %q does not contain %q", err, tc.wantMsg)
			}
			if tc.wantSuggestion != "" {
				var nf *explain.NotFoundError
				if !errors.As(err, &nf) || !slices.Contains(nf.Suggestions, tc.wantSuggestion) {
					t.Errorf("error %q: want suggestion %q", err, tc.wantSuggestion)
				}
			}
		})
	}
}

// TestEverythingRenders explains every definition (plain, recursive and JSON)
// and every first-level field, in all three sources.
func TestEverythingRenders(t *testing.T) {
	c := catalog(t)
	for _, name := range explain.SourceNames {
		src, err := c.Source(name)
		if err != nil {
			t.Fatal(err)
		}
		if len(src.Names()) == 0 {
			t.Fatalf("%s: no definitions", name)
		}
		prefix := name + ":"
		for _, def := range src.Names() {
			r, err := c.Resolve(prefix + def)
			if err != nil {
				t.Fatalf("%s%s: %v", prefix, def, err)
			}
			for _, opts := range []render.Options{{}, {Recursive: true, MaxDepth: 4}} {
				if err := render.Plain(io.Discard, r, c.Version(), opts); err != nil {
					t.Fatalf("%s%s: %v", prefix, def, err)
				}
			}
			if err := render.JSON(io.Discard, r, c.Version(), src); err != nil {
				t.Fatalf("%s%s json: %v", prefix, def, err)
			}
			for _, child := range explain.Children(r.Node) {
				expr := r.Expr() + "." + child
				cr, err := c.Resolve(expr)
				if err != nil {
					t.Fatalf("%s: %v", expr, err)
				}
				if err := render.Plain(io.Discard, cr, c.Version(), render.Options{}); err != nil {
					t.Fatalf("%s: %v", expr, err)
				}
			}
		}
	}
}

func TestCounts(t *testing.T) {
	c := catalog(t)
	want := map[string]int{explain.Bot: 744, explain.Card: 73, explain.Connector: 4}
	for name, n := range want {
		src, err := c.Source(name)
		if err != nil {
			t.Fatal(err)
		}
		if got := len(src.Names()); got != n {
			t.Errorf("%s: %d definitions, want %d (vendored schema changed? update this test)", name, got, n)
		}
	}
}

// TestNoStaleOverlayKeys fails when descriptions.yaml has a key that matches
// nothing in the schemas (a typo, or a kind removed upstream).
func TestNoStaleOverlayKeys(t *testing.T) {
	c := catalog(t)
	for _, name := range explain.SourceNames {
		src, err := c.Source(name)
		if err != nil {
			t.Fatal(err)
		}
		for _, k := range c.Overlay().Section(name).Keys() {
			if !src.ValidKey(k) {
				t.Errorf("%s: stale overlay key %q", name, k)
			}
		}
		_, missing, err := c.Roots(name)
		if err != nil {
			t.Fatal(err)
		}
		for _, k := range missing {
			t.Errorf("%s: stale roots kind %q", name, k)
		}
	}
}

func TestSplitSource(t *testing.T) {
	cases := map[string][2]string{
		"SendActivity":           {"bot", "SendActivity"},
		"bot:SendActivity":       {"bot", "SendActivity"},
		"card:Input.Text":        {"card", "Input.Text"},
		"Connector:shared_x.Op":  {"connector", "shared_x.Op"},
		"connectors:shared_x.Op": {"connector", "shared_x.Op"},
	}
	for in, want := range cases {
		s, rest := explain.SplitSource(in)
		if s != want[0] || rest != want[1] {
			t.Errorf("SplitSource(%q) = %q, %q; want %q, %q", in, s, rest, want[0], want[1])
		}
	}
}
