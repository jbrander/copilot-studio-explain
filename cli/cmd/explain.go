package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"copilot-studio-explain/cli/internal/explain"
	"copilot-studio-explain/cli/internal/render"
)

type explainOpts struct {
	recursive bool
	maxDepth  int
	output    string
}

func newExplainCmd(a *app) *cobra.Command {
	o := &explainOpts{}
	c := &cobra.Command{
		Use:   "explain RESOURCE[.field...]",
		Short: "Describe a definition or one of its fields",
		Long: `Describe a definition or one of its fields.

RESOURCE is a definition from one of three schemas:
  bot (default)  Copilot Studio YAML: SendActivity, AdaptiveDialog, GptComponentMetadata, ...
  card:          Adaptive Cards: card:AdaptiveCard, card:Input.ChoiceSet, ...
  connector:     connector operations: connector:shared_office365.<operationId>

Follow it with a dotted field path. Arrays and maps are stepped into
automatically. When a field is a union (DialogAction, TriggerBase, ...), the
next segment picks the member kind.`,
		Example: `  cs-explain explain SendActivity
  cs-explain explain AdaptiveDialog.beginDialog
  cs-explain explain AdaptiveDialog.beginDialog.OnRecognizedIntent.actions.Question.entity
  cs-explain explain GptComponentMetadata --recursive --max-depth 2
  cs-explain explain card:Input.ChoiceSet.choices.data
  cs-explain explain connector:shared_office365users
  cs-explain explain ConversationOutcome -o json`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cat, err := a.cat()
			if err != nil {
				return err
			}
			res, err := cat.Resolve(args[0])
			if err != nil {
				return err
			}
			switch o.output {
			case "", "plaintext", "text":
				return render.Plain(cmd.OutOrStdout(), res, cat.Version(),
					render.Options{Recursive: o.recursive, MaxDepth: o.maxDepth})
			case "json":
				src, err := cat.Source(res.Source)
				if err != nil {
					return err
				}
				return render.JSON(cmd.OutOrStdout(), res, cat.Version(), src)
			default:
				return fmt.Errorf("unknown output format %q (want plaintext or json)", o.output)
			}
		},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) > 0 {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
			return complete(a, toComplete)
		},
	}
	c.Flags().BoolVar(&o.recursive, "recursive", false, "print the field tree (union members are not expanded)")
	c.Flags().IntVar(&o.maxDepth, "max-depth", 10, "with --recursive, how many levels to expand (0 = unlimited)")
	c.Flags().StringVarP(&o.output, "output", "o", "plaintext", "output format: plaintext or json")
	return c
}

// complete offers definition names, then field / member names under the
// already-typed path.
func complete(a *app, toComplete string) ([]string, cobra.ShellCompDirective) {
	cat, err := a.cat()
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}
	srcName, rest := explain.SplitSource(toComplete)
	prefix := ""
	if srcName != explain.Bot || strings.HasPrefix(strings.ToLower(toComplete), "bot:") {
		prefix = toComplete[:len(toComplete)-len(rest)]
	}
	src, err := cat.Source(srcName)
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}
	var out []string
	// Definition names (they can contain dots, e.g. card:Input.Text).
	for _, n := range src.Names() {
		if strings.HasPrefix(strings.ToLower(n), strings.ToLower(rest)) {
			out = append(out, prefix+n)
		}
	}
	if srcName == explain.Bot && prefix == "" && !strings.Contains(rest, ".") {
		for _, s := range []string{explain.Card + ":", explain.Connector + ":"} {
			if strings.HasPrefix(s, strings.ToLower(rest)) {
				out = append(out, s)
			}
		}
	}
	// Children of the deepest resolvable parent path.
	if i := strings.LastIndex(rest, "."); i >= 0 {
		parent, partial := rest[:i], rest[i+1:]
		if res, err := cat.Resolve(prefix + parent); err == nil {
			for _, ch := range explain.Children(res.Node) {
				if strings.HasPrefix(strings.ToLower(ch), strings.ToLower(partial)) {
					out = append(out, prefix+parent+"."+ch)
				}
			}
		}
	}
	directive := cobra.ShellCompDirectiveNoFileComp | cobra.ShellCompDirectiveNoSpace
	return out, directive
}
