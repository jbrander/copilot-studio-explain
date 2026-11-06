package cmd

import (
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"copilot-studio-explain/cli/internal/explain"
	"copilot-studio-explain/cli/internal/model"
	"copilot-studio-explain/cli/internal/render"
)

func newListCmd(a *app) *cobra.Command {
	var (
		filter       string
		kindsOnly    bool
		undocumented bool
	)
	c := &cobra.Command{
		Use:   "list [bot|card|connector]",
		Short: "List the definitions in a schema, like `kubectl api-resources`",
		Example: `  cs-explain list
  cs-explain list --kinds-only --filter search
  cs-explain list card
  cs-explain list connector`,
		Args:      cobra.MaximumNArgs(1),
		ValidArgs: explain.SourceNames,
		RunE: func(cmd *cobra.Command, args []string) error {
			cat, err := a.cat()
			if err != nil {
				return err
			}
			srcName := explain.Bot
			if len(args) == 1 {
				srcName, _ = explain.SplitSource(strings.TrimSuffix(args[0], ":") + ":")
				if !strings.EqualFold(strings.TrimSuffix(args[0], "s"), srcName) && !strings.EqualFold(args[0], srcName) {
					return fmt.Errorf("unknown schema %q (want one of: %s)", args[0], strings.Join(explain.SourceNames, ", "))
				}
			}
			src, err := cat.Source(srcName)
			if err != nil {
				return err
			}
			tw := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
			disc := map[string]string{explain.Bot: "KIND", explain.Card: "TYPE", explain.Connector: "FIELDS"}[srcName]
			fmt.Fprintf(tw, "NAME\tSHAPE\t%s\tDESCRIPTION\n", disc)
			count := 0
			for _, name := range src.Names() {
				if filter != "" && !strings.Contains(strings.ToLower(name), strings.ToLower(filter)) {
					continue
				}
				n, ok := src.Lookup(name)
				if !ok {
					continue
				}
				third := n.DiscriminatorValue
				if srcName == explain.Connector {
					third = fmt.Sprint(len(n.Fields))
				}
				if kindsOnly && (n.Shape != model.Object || n.DiscriminatorValue == "") {
					continue
				}
				if undocumented && n.Description != "" {
					continue
				}
				fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n", name, n.Shape, third, render.FirstSentence(n.Description, 70))
				count++
			}
			if err := tw.Flush(); err != nil {
				return err
			}
			fmt.Fprintf(cmd.ErrOrStderr(), "%d definitions\n", count)
			return nil
		},
	}
	c.Flags().StringVar(&filter, "filter", "", "only names containing this substring (case-insensitive)")
	c.Flags().BoolVar(&kindsOnly, "kinds-only", false, "only objects with a `kind:`/`type:` value (what you can write in YAML)")
	c.Flags().BoolVar(&undocumented, "undocumented", false, "only definitions without a description")
	return c
}
