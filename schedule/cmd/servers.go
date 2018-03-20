package cmd

import (
	"github.com/mami-project/trafic/runner"
	"github.com/spf13/cobra"
)

var serversCmd = &cobra.Command{
	Use:   "servers",
	Short: "Run the server side of a traffic mix",
	Run:   servers,
}

func init() {
	rootCmd.AddCommand(serversCmd)
}

func servers(cmd *cobra.Command, args []string) {
	run(runner.RoleServer)
}
