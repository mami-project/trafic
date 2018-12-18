package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mami-project/trafic/flowsim/udp"
)

var sinkIp string
var sinkPort int
var sinkVerbose bool

// sinkCmd represents the sink command
var sinkCmd = &cobra.Command{
	Use:   "sink",
	Short: "Start a flowsim UDP sink",
	Long: `Will run flowsim as a UDP CBR sink
and print stats like mean delay and mean jitter for the CBR flow at the end`,
	Run: func(cmd *cobra.Command, args []string) {
		udp.Sink(sinkIp, sinkPort, sinkVerbose)
	},
}

func init() {
	rootCmd.AddCommand(sinkCmd)

	sinkCmd.PersistentFlags().StringVarP(&sinkIp, "ip", "I", "127.0.0.1", "IP address or host name to listen on for the flowsim UDP sink")
	sinkCmd.PersistentFlags().IntVarP(&sinkPort, "port", "p", 8081, "UDP port of the flowsim UDP sink")
	sinkCmd.PersistentFlags().BoolVarP(&sinkVerbose, "verbose", "v", false, "Print per packet info")
}
