package cmd

import (
	"github.com/spf13/cobra"
)

// Version is the CLI version. Override with `-ldflags "-X .../cmd.Version=..."` at build time.
var Version = "0.0.1-dev"

// NewRoot wires the top-level command and all subcommands.
func NewRoot() *cobra.Command {
	root := &cobra.Command{
		Use:   "blueprint",
		Short: "Inspire Blueprint — pull and run open LLMs locally",
		Long: `Blueprint downloads open-model GGUF weights and runs them through
llama.cpp on your machine, exposing an OpenAI-compatible endpoint
on 127.0.0.1:8080.

Free, no telemetry, no account. Made by Inspire AI Lab.`,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	root.AddCommand(newVersionCmd())
	root.AddCommand(newPullCmd())
	root.AddCommand(newServeCmd())
	root.AddCommand(newStatusCmd())
	root.AddCommand(newStopCmd())
	root.AddCommand(newRuntimeCmd())

	return root
}
