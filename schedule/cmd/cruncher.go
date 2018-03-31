package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mami-project/trafic/cruncher"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var crunchCmd = &cobra.Command{
	Use:   "crunch",
	Short: "Run the requested cruncher on an iperf3 JSON output",
	Run:   crunch,
}

func init() {
	rootCmd.AddCommand(crunchCmd)

	flags := crunchCmd.Flags()

	flags.String("crunch-input", "", "iperf3 JSON file to process")
	viper.BindPFlag("crunch.input", flags.Lookup("crunch-input"))

	flags.String("crunch-outfmt", "csv", "output format")
	viper.BindPFlag("crunch.outfmt", flags.Lookup("crunch-outfmt"))
}

func crunch(cmd *cobra.Command, args []string) {
	raw, err := ioutil.ReadFile(viper.GetString("crunch.input"))
	if err != nil {
		log.Fatal(err)
	}

	var c cruncher.Cruncher

	outfmt := viper.GetString("crunch.outfmt")

	switch outfmt {
	case "csv":
		c = cruncher.NewCSVCruncher()
	case "influxdb":
		c = cruncher.NewInfluxDBCruncher("trafic_samples")
	case "telegraf":
		c = cruncher.NewTelegrafCruncher()
	default:
		log.Fatalf("unknown output format: %s", outfmt)
	}

	out, err := cruncher.Crunch(c, raw)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
