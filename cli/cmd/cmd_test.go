package cmd

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

var update = flag.Bool("update", false, "rewrite golden files")

func runCLI(t *testing.T, args ...string) (string, string, int) {
	t.Helper()
	var out, errOut bytes.Buffer
	code := run(args, &out, &errOut)
	return out.String(), errOut.String(), code
}

func TestGolden(t *testing.T) {
	cases := map[string][]string{
		"sendactivity":             {"SendActivity"},
		"sendactivity-shorthand":   {"explain", "SendActivity"},
		"adaptivedialog-recursive": {"AdaptiveDialog", "--recursive", "--max-depth", "2"},
		"topic-actions":            {"AdaptiveDialog.beginDialog.OnRecognizedIntent.actions"},
		"enum-conversationoutcome": {"ConversationOutcome"},
		"leaf-field":               {"SendActivity.disabled"},
		"shorthand-entity":         {"BooleanPrebuiltEntity"},
		"card-choiceset":           {"card:Input.ChoiceSet"},
		"card-dotted-field":        {"card:Input.ChoiceSet.choices.data"},
		"connector":                {"connector:shared_office365users"},
		"connector-op-inputs":      {"connector:shared_office365users.UserProfile_V2.inputs"},
		"json-enum":                {"ConversationOutcome", "-o", "json"},
		"list":                     {"list"},
	}
	for name, args := range cases {
		t.Run(name, func(t *testing.T) {
			out, errOut, code := runCLI(t, args...)
			if code != 0 {
				t.Fatalf("exit %d: %s", code, errOut)
			}
			path := filepath.Join("..", "testdata", "golden", name+".txt")
			if *update {
				if err := os.WriteFile(path, []byte(out), 0o644); err != nil {
					t.Fatal(err)
				}
				return
			}
			want, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("%v (run `go test ./cmd -update` to create it)", err)
			}
			if out != string(want) {
				t.Errorf("output differs from %s (run `go test ./cmd -update` to accept):\n--- got\n%s", path, out)
			}
		})
	}
	// The shorthand and the explicit subcommand must agree.
	a, _, _ := runCLI(t, "SendActivity")
	b, _, _ := runCLI(t, "explain", "SendActivity")
	if a != b {
		t.Error("`cs-explain X` and `cs-explain explain X` differ")
	}
}

func TestErrors(t *testing.T) {
	cases := map[string]struct {
		args []string
		want string
	}{
		"typo":       {[]string{"SendActvity"}, "Did you mean: SendActivity?"},
		"bad field":  {[]string{"SendActivity.activty"}, "Did you mean: activity?"},
		"union":      {[]string{"AdaptiveDialog.beginDialog.actions"}, "is a union"},
		"bad output": {[]string{"SendActivity", "-o", "yaml"}, "unknown output format"},
		"bad schema": {[]string{"list", "nope"}, "unknown schema"},
		"two args":   {[]string{"explain", "a", "b"}, "accepts 1 arg"},
		"schema dir": {[]string{"--schema-dir", "/nonexistent", "SendActivity"}, "schema dir"},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			_, errOut, code := runCLI(t, tc.args...)
			if code == 0 {
				t.Fatal("expected non-zero exit")
			}
			if !strings.Contains(errOut, tc.want) {
				t.Errorf("stderr %q does not contain %q", errOut, tc.want)
			}
		})
	}
}

func TestList(t *testing.T) {
	out, errOut, code := runCLI(t, "list", "--kinds-only", "--filter", "sendactivity")
	if code != 0 {
		t.Fatal(errOut)
	}
	if !strings.Contains(out, "SendActivity ") || !strings.Contains(errOut, "1 definitions") {
		t.Errorf("unexpected list output:\n%s%s", out, errOut)
	}
	out, errOut, code = runCLI(t, "list")
	if code != 0 || !strings.Contains(errOut, "7 file kinds (of 744 definitions") || strings.Contains(out, "SendActivity") {
		t.Errorf("default list should show only the file kinds:\n%s%s", out, errOut)
	}
	for _, args := range [][]string{{"list", "--all"}, {"list", "-A"}, {"list", "bot", "--all"}} {
		if _, errOut, _ := runCLI(t, args...); !strings.Contains(errOut, "744 definitions") {
			t.Errorf("%v: want all 744 definitions, got %q", args, errOut)
		}
	}
	for _, src := range []string{"card", "connector", "connectors"} {
		if _, errOut, code := runCLI(t, "list", src); code != 0 {
			t.Errorf("list %s: %s", src, errOut)
		}
	}
}

func TestCoverage(t *testing.T) {
	out, errOut, code := runCLI(t, "coverage")
	if code != 0 {
		t.Fatalf("coverage failed (stale keys?):\n%s%s", out, errOut)
	}
	for _, s := range []string{"bot", "card", "connector"} {
		if !strings.Contains(out, s) {
			t.Errorf("coverage output lacks %s:\n%s", s, out)
		}
	}
}

func TestSchemaDir(t *testing.T) {
	dir := filepath.Join("..", "internal", "schemas", "data")
	out, errOut, code := runCLI(t, "--schema-dir", dir, "SendActivity")
	if code != 0 {
		t.Fatal(errOut)
	}
	embedded, _, _ := runCLI(t, "SendActivity")
	if out != embedded {
		t.Error("--schema-dir output differs from embedded output for the same files")
	}
}

func TestComplete(t *testing.T) {
	a := &app{}
	cases := map[string]string{
		"SendAct":                     "SendActivity",
		"SendActivity.act":            "SendActivity.activity",
		"SendActivity.activity.":      "SendActivity.activity.EventActivity",
		"car":                         "card:",
		"card:Input.Ch":               "card:Input.ChoiceSet",
		"card:Input.ChoiceSet.ch":     "card:Input.ChoiceSet.choices.data",
		"connector:shared_office365u": "connector:shared_office365users",
	}
	for in, want := range cases {
		got, _ := complete(a, in)
		if !slices.Contains(got, want) {
			t.Errorf("complete(%q) = %v, want it to contain %q", in, got, want)
		}
	}
}
