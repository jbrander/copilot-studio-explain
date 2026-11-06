package cmd

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"copilot-studio-explain/cli/internal/explain"
	"copilot-studio-explain/cli/internal/model"
	"copilot-studio-explain/cli/internal/render"
)

func newListCmd(a *app) *cobra.Command {
	var (
		all          bool
		filter       string
		kindsOnly    bool
		undocumented bool
	)
	c := &cobra.Command{
		Use:   "list [bot|card|connector]",
		Short: "List the definitions in a schema, like `kubectl api-resources`",
		Long: `List the definitions in a schema, like ` + "`kubectl api-resources`" + `.

For the bot schema the default is the kinds that can be the root of a
.mcs.yml file (agent, topics, tools, knowledge, ...). --all lists every
definition; --filter, --kinds-only and --undocumented imply --all. The card
and connector schemas always list everything.`,
		Example: `  cs-explain list
  cs-explain list --all
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
			if !all && filter == "" && !kindsOnly && !undocumented {
				roots, missing, err := cat.Roots(srcName)
				if err != nil {
					return err
				}
				if len(missing) > 0 {
					return fmt.Errorf("file kinds in the overlay are no longer in the %s schema: %s", srcName, strings.Join(missing, ", "))
				}
				if len(roots) > 0 {
					return listRoots(cmd.OutOrStdout(), cmd.ErrOrStderr(), roots, len(src.Names()))
				}
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
	c.Flags().BoolVarP(&all, "all", "A", false, "list every definition, not just the kinds that can be the root of a file")
	c.Flags().StringVar(&filter, "filter", "", "only names containing this substring (case-insensitive); implies --all")
	c.Flags().BoolVar(&kindsOnly, "kinds-only", false, "only objects with a `kind:`/`type:` value (what you can write in YAML); implies --all")
	c.Flags().BoolVar(&undocumented, "undocumented", false, "only definitions without a description; implies --all")
	return c
}

// listRoots prints the kinds that can be the root of a file.
func listRoots(out, errOut io.Writer, roots []explain.RootKind, total int) error {
	tw := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "KIND\tFILE\tDESCRIPTION\n")
	for _, r := range roots {
		fmt.Fprintf(tw, "%s\t%s\t%s\n", r.Node.Name, r.File, render.FirstSentence(r.Node.Description, 70))
	}
	if err := tw.Flush(); err != nil {
		return err
	}
	fmt.Fprintf(errOut, "%d file kinds (of %d definitions; use --all to list them all)\n", len(roots), total)
	return nil
}
