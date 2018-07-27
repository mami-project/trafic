// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/mami-project/trafic/flowsim/tcp"
)

var ip string
var port int
var interval int
var iter int
var clientTos string
var burstStr string

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start flowsim in client mode",
	Long: `Will run flowsim in client mode
and try to talk to an flowsim in server mode.`,
	Run: func(cmd *cobra.Command, args []string) {
		tos, err := flow.Dscp(clientTos)
		if err != nil {
			fmt.Printf("Error decoding DSCP (%s): %v\n", clientTos, err)
		} else {
			val, err := utoi(burstStr)
			if err != nil {
				fmt.Printf("Warning: %v, generating %d byte bursts", err, val)
			}
			flow.Client(ip, port, iter, interval, val, tos * 4)
		}
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	clientCmd.PersistentFlags().StringVarP(&ip, "ip", "I", "127.0.0.1", "IP address or host name of the flowsim server to talk to")
	clientCmd.PersistentFlags().IntVarP(&port, "port", "p", 8081, "TCP port of the flowsim server")
	clientCmd.PersistentFlags().IntVarP(&iter, "iter", "n", 6, "Number of bursts")
	clientCmd.PersistentFlags().IntVarP(&interval, "interval", "t", 10, "Interval in secs between bursts")
	clientCmd.PersistentFlags().StringVarP(&burstStr, "burst", "N", "1M", "Size of each burst (as x(.xxx)?[kmgtKMGT]?)")
	clientCmd.PersistentFlags().StringVarP(&clientTos, "TOS", "T", "CS0", "Value of the DSCP field in the IP packets (valid int or DSCP-Id)")
}
