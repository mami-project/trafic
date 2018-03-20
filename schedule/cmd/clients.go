package cmd

import (
	"github.com/mami-project/trafic/runner"
	"github.com/spf13/cobra"
)

var clientsCmd = &cobra.Command{
	Use:   "clients",
	Short: "Run the client side of a traffic mix",
	Run:   clients,
}

func init() {
	rootCmd.AddCommand(clientsCmd)
}

func clients(cmd *cobra.Command, args []string) {
	run(runner.RoleClient)
}
