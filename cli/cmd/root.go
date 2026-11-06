// Package cmd wires the cs-explain command line.
package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"copilot-studio-explain/cli/internal/explain"
)

type app struct {
	schemaDir string
	catalog   *explain.Catalog
}

func (a *app) cat() (*explain.Catalog, error) {
	if a.catalog != nil {
		return a.catalog, nil
	}
	dir := a.schemaDir
	if dir == "" {
		dir = os.Getenv("CS_EXPLAIN_SCHEMA_DIR")
	}
	c, err := explain.Open(dir)
	if err != nil {
		return nil, err
	}
	a.catalog = c
	return c, nil
}

// NewRoot builds the command tree. out/errOut are used for tests.
func NewRoot(out, errOut io.Writer) *cobra.Command {
	a := &app{}
	ex := newExplainCmd(a)
	root := &cobra.Command{
		Use:   "cs-explain [RESOURCE[.field...]]",
		Short: "Describe the YAML schema of Copilot Studio agents, like `kubectl explain`",
		Long: `cs-explain describes the fields of the YAML you write for Microsoft Copilot Studio
agents (.mcs.yml topics, actions, triggers, agent and knowledge files), Adaptive
Cards, and connector operations. It reads schemas vendored from
github.com/microsoft/skills-for-copilot-studio, which are embedded in the binary.

A RESOURCE is a schema definition, optionally followed by a dotted field path.
When a field holds a union (for example a topic's actions), the next segment is
the kind you want:

  cs-explain SendActivity
  cs-explain AdaptiveDialog.beginDialog.OnRecognizedIntent.actions
  cs-explain card:Input.ChoiceSet.choices
  cs-explain connector:shared_office365.<operationId>.inputs

"cs-explain RESOURCE" is short for "cs-explain explain RESOURCE".`,
		Args:              cobra.ArbitraryArgs,
		ValidArgsFunction: ex.ValidArgsFunction,
		SilenceUsage:      true,
		SilenceErrors:     true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			return ex.RunE(cmd, args)
		},
	}
	root.SetOut(out)
	root.SetErr(errOut)
	root.PersistentFlags().StringVar(&a.schemaDir, "schema-dir", "",
		"read schemas from this upstream reference/ directory instead of the embedded copy (env CS_EXPLAIN_SCHEMA_DIR)")
	// The root accepts explain's flags so the shorthand form works.
	root.Flags().AddFlagSet(ex.Flags())
	root.AddCommand(ex, newListCmd(a), newCoverageCmd(a))
	return root
}

// Execute runs the CLI and returns the process exit code.
func Execute() int {
	return run(os.Args[1:], os.Stdout, os.Stderr)
}

func run(args []string, out, errOut io.Writer) int {
	root := NewRoot(out, errOut)
	root.SetArgs(args)
	if err := root.Execute(); err != nil {
		var nf *explain.NotFoundError
		if errors.As(err, &nf) {
			fmt.Fprintln(errOut, "error:", nf.Error())
			return 1
		}
		fmt.Fprintln(errOut, "error:", err)
		return 1
	}
	return 0
}
