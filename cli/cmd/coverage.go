package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"copilot-studio-explain/cli/internal/explain"
	"copilot-studio-explain/cli/internal/jsonschema"
)

func newCoverageCmd(a *app) *cobra.Command {
	var showMissing bool
	c := &cobra.Command{
		Use:   "coverage",
		Short: "Report how much of each schema has curated descriptions, and stale overlay keys",
		Long: `Report how much of each schema has a description (from the curated overlay,
internal/describe/descriptions.yaml, or from the schema itself), and list
overlay keys that no longer match anything in the schema (for example after
the vendored schema is updated). Exits non-zero if any key is stale.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cat, err := a.cat()
			if err != nil {
				return err
			}
			out := cmd.OutOrStdout()
			fmt.Fprintf(out, "Schemas: %s\n\n", cat.Version())
			stale := 0
			for _, name := range explain.SourceNames {
				src, err := cat.Source(name)
				if err != nil {
					return err
				}
				sec := cat.Overlay().Section(name)
				var keys []string
				for _, k := range sec.Keys() {
					if !src.ValidKey(k) {
						keys = append(keys, k)
					}
				}
				if eng, ok := src.(*jsonschema.Engine); ok {
					cov := eng.Coverage()
					fmt.Fprintf(out, "%s\n", cov)
					if showMissing && len(cov.UndescribedKinds) > 0 {
						fmt.Fprintf(out, "  undescribed kinds: %s\n", strings.Join(cov.UndescribedKinds, ", "))
					}
				} else {
					fmt.Fprintf(out, "%-6s %d connectors (descriptions come from the connector files)\n", name, len(src.Names()))
				}
				_, missing, err := cat.Roots(name)
				if err != nil {
					return err
				}
				for _, k := range missing {
					keys = append(keys, "roots: "+k)
				}
				fmt.Fprintf(out, "       overlay entries: %d, stale: %d\n", sec.Len(), len(keys))
				for _, k := range keys {
					fmt.Fprintf(out, "         stale: %s\n", k)
				}
				stale += len(keys)
			}
			if stale > 0 {
				return fmt.Errorf("%d stale overlay key(s)", stale)
			}
			return nil
		},
	}
	c.Flags().BoolVar(&showMissing, "missing", false, "list kinds that have no description")
	return c
}
