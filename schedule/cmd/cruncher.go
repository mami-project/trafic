package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mami-project/trafic/cruncher"
	"github.com/spf13/cobra"
)

var crunchCmd = &cobra.Command{
	Use:   "crunch <file>",
	Short: "Run the Telegraf cruncher on an iperf3 JSON output",
	Run:   crunch,
	Args:  cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(crunchCmd)
}

func crunch(cmd *cobra.Command, args []string) {
	raw, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	c := cruncher.NewTelegrafCruncher()

	out, err := c.Crunch(raw)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
