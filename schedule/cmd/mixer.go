package cmd

import (
	"log"

	"github.com/mami-project/trafic/mixer"
	"github.com/spf13/cobra"
)

var (
	DescFile string
	OutDir   string
)

var mixerCmd = &cobra.Command{
	Use:   "mixer",
	Short: "Creates a traffic mix from a description file",
	Run:   mix,
}

func init() {
	rootCmd.AddCommand(mixerCmd)

	flags := mixerCmd.Flags()

	flags.StringVar(&DescFile, "description-file", "~/mix.yaml", "A YAML file containing the description of the traffic mix")
	flags.StringVar(&OutDir, "out-dir", "/tmp/trafic/mix", "the base directory where to put the generated files")
}

func mix(cmd *cobra.Command, args []string) {
	c, err := mixer.NewDescriptionFromFile(DescFile)
	if err != nil {
		log.Fatal(err)
	}

	err = Generate(c, OutDir)
	if err != nil {
		log.Fatal(err)
	}
}

func Generate(desc *mixer.Description, baseDir string) error {
	log.Printf("saving generated configuration to: %s", baseDir)

	for i := range desc.Flows {
		kind := desc.Flows[i].Kind
		m, err := mixer.LookupMixer(kind)
		if err != nil {
			log.Fatal(err)
		}

		err = (*m).WriteConf(baseDir, desc.Global, desc.Flows[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
