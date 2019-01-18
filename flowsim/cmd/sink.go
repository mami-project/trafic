package cmd

import (
	// "fmt"
	common "github.com/mami-project/trafic/flowsim/common"
	"github.com/mami-project/trafic/flowsim/udp"
	"github.com/spf13/cobra"
)

var sinkIp string
var sinkPort int
var sinkMulti bool
var sinkVerbose bool
var sinkIpv6 bool

// sinkCmd represents the sink command
var sinkCmd = &cobra.Command{
	Use:   "sink",
	Short: "Start a flowsim UDP sink",
	Long: `Will run flowsim as a one-of UDP CBR sink
and print stats like mean delay and mean jitter for the CBR flow at the end.
This is a pure sink, set DSCP in the source`,
	Run: func(cmd *cobra.Command, args []string) {
		useIp, err := common.FirstIP(sinkIp, sinkIpv6)
		common.FatalError(err)
		udp.Sink(useIp, sinkPort, sinkMulti, sinkVerbose)
	},
}

func init() {
	rootCmd.AddCommand(sinkCmd)

	sinkCmd.PersistentFlags().StringVarP(&sinkIp, "ip", "I", "localhost", "IP address or host name to listen on for the flowsim UDP sink")
	sinkCmd.PersistentFlags().IntVarP(&sinkPort, "port", "p", 8081, "UDP port of the flowsim UDP sink")
	sinkCmd.PersistentFlags().BoolVarP(&sinkMulti, "multi", "m", false, "Stay in the sink forever and print stats for multiple incoming streams")
	sinkCmd.PersistentFlags().BoolVarP(&sinkVerbose, "verbose", "v", false, "Print per packet info")
	sinkCmd.PersistentFlags().BoolVarP(&sinkIpv6, "ipv6", "6", false, "Use IPv6 (default is IPv4)")
}
