package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"flowsim/tcp"
	"flowsim/quic"
)

<<<<<<< HEAD
var clientIp string
var clientPort int
var clientInterval int
var clientIter int
var clientTos string
var clientBurst string
var clientQuic bool
=======
var ip string
var port int
var interval int
var iter int
var clientTos string
var burstStr string
>>>>>>> fbee7a0b5f883550c24032fbc22b98cb434236a0

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start flowsim in client mode",
	Long: `Will run flowsim in client mode
and try to talk to an flowsim in server mode.`,
	Run: func(cmd *cobra.Command, args []string) {
<<<<<<< HEAD
		val, err := utoi(clientBurst)
		if err != nil {
			fmt.Printf("Warning: %v, generating %d byte bursts", err, val)
		}
		if clientQuic {
			quic.Client(clientIp, clientPort, clientIter, clientInterval, val)
		} else {
			tos, err := tcp.Dscp(clientTos)
			if err != nil {
				fmt.Printf("Error decoding DSCP (%s): %v\n", clientTos, err)
			} else {
				tcp.Client(clientIp, clientPort, clientIter, clientInterval, val, tos * 4)
			}
		}
=======
		Tos, err := flow.Dscp(clientTos)
		if err != nil {
			fmt.Printf("Error decoding DSCP (%s): %v\n", clientTos, err)
			return
		}

		flow.Client(ip, port, iter, interval, iperf3_atoi(burstStr), Tos * 4)
>>>>>>> fbee7a0b5f883550c24032fbc22b98cb434236a0
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

<<<<<<< HEAD
	clientCmd.PersistentFlags().StringVarP(&clientIp, "ip", "I", "127.0.0.1", "IP address or host name of the flowsim server to talk to")
	clientCmd.PersistentFlags().IntVarP(&clientPort, "port", "p", 8081, "TCP port of the flowsim server")
	clientCmd.PersistentFlags().IntVarP(&clientIter, "iter", "n", 6, "Number of bursts")
	clientCmd.PersistentFlags().IntVarP(&clientInterval, "interval", "t", 10, "Interval in secs between bursts")
	clientCmd.PersistentFlags().StringVarP(&clientBurst, "burst", "N", "1M", "Size of each burst (as x(.xxx)?[kmgtKMGT]?)")
	clientCmd.PersistentFlags().StringVarP(&clientTos, "TOS", "T", "CS0", "Value of the DSCP field in the IP packets (valid int or DSCP-Id)")
	clientCmd.PersistentFlags().BoolVarP(&clientQuic,"quic", "Q", false, "Use QUIC (default is TCP)")
=======
	clientCmd.PersistentFlags().StringVarP(&ip, "ip", "I", "127.0.0.1", "IP address or host name of the flowsim server to talk to")
	clientCmd.PersistentFlags().IntVarP(&port, "port", "p", 8081, "TCP port of the flowsim server")
	clientCmd.PersistentFlags().IntVarP(&iter, "iter", "n", 6, "Number of bursts")
	clientCmd.PersistentFlags().IntVarP(&interval, "interval", "t", 10, "Interval in secs between bursts")
	clientCmd.PersistentFlags().StringVarP(&burstStr, "burst", "N", "1M", "Size of each burst (as x(.xxx)?[kmgtKMGT]?)")
	clientCmd.PersistentFlags().StringVarP(&clientTos, "TOS", "T", "CS0", "Value of the DSCP field in the IP packets (valid int or DSCP-Id)")
>>>>>>> fbee7a0b5f883550c24032fbc22b98cb434236a0
}
