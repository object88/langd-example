package cmd

import (
	"fmt"

	uuidgen "github.com/object88/langd-example"
	"github.com/spf13/cobra"
)

func createGenerateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generates one or more UUIDs",
		RunE: func(_ *cobra.Command, _ []string) error {

			fmt.Println(uuidgen.GenerateUUID())

			return nil
		},
	}
	return cmd
}
