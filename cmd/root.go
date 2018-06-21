package cmd

import "github.com/spf13/cobra"

// InitializeCommands sets up the cobra commands
func InitializeCommands() *cobra.Command {
	rootCmd := createRootCommand()

	rootCmd.AddCommand(
		createGenerateCommand(),
	)

	return rootCmd
}

func createRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uuidgen",
		Short: "uuidgen creates UUIDs",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	return cmd
}
