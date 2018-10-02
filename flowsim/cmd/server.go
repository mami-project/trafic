package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/mami-project/trafic/flowsim/tcp"
	"github.com/mami-project/trafic/flowsim/quic"
)

var serverIp string
var serverPort int
var serverSingle bool
var serverTos string
var serverQuic bool

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start an flow server",
	Long: `Start an ABR server.
It will basically sit there and wait for the client to request bunches of data
over a TCP connection`,
	Run: func(cmd *cobra.Command, args []string) {
		if serverQuic {
			quic.Server(serverIp, serverPort, serverSingle)
		} else {
			tos, err := tcp.Dscp(serverTos)
			if err != nil {
				fmt.Printf("Error decoding DSCP (%s): %v\n", serverTos, err)
			} else {
				tcp.Server(serverIp, serverPort, serverSingle, tos * 4)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.PersistentFlags().StringVarP(&serverIp, "ip", "I", "127.0.0.1", "IP address or host name bound to the flowsim server")
	serverCmd.PersistentFlags().IntVarP(&serverPort, "port", "p", 8081, "TCP port bound to the flowsim server")
	serverCmd.PersistentFlags().BoolVarP(&serverSingle,"one-off", "1", false, "Just accept one connection and quit (default is run until killed)")
	serverCmd.PersistentFlags().StringVarP(&serverTos, "TOS", "T", "CS0", "Value of the DSCP field in the IP layer (number or DSCP id)")
	serverCmd.PersistentFlags().BoolVarP(&serverQuic,"quic", "Q", false, "Use QUIC (default is TCP)")
}
