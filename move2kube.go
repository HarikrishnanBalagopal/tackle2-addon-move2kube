package main

import (
	"fmt"
	"os"

	"github.com/konveyor/tackle2-addon/command"
)

// Build index.html
func runMove2Kube(input, output string) error {
	if err := os.RemoveAll(output); err != nil {
		return fmt.Errorf("failed to remove the output directory. Error: %w", err)
	}
	addon.Activity("Running Move2Kube transform on the input directory.")
	cmd := command.Command{Path: "/usr/bin/move2kube"}
	cmd.Options.Add("transform")
	cmd.Options.Add("--source", input)
	cmd.Options.Add("--output", output)
	cmd.Options.Add("--log-level", "trace")
	cmd.Options.Add("--qa-skip")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run the Move2Kube transform command. Error: %w", err)
	}
	return nil
}
