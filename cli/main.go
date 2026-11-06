// Command cs-explain describes the YAML schema of Microsoft Copilot Studio
// agents, in the style of `kubectl explain`.
package main

import (
	"os"

	"copilot-studio-explain/cli/cmd"
)

func main() {
	os.Exit(cmd.Execute())
}
