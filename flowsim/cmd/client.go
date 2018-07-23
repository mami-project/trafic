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
var TOS int
var burstStr string

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start flowsim in client mode",
	Long: `Will run flowsim in client mode
and try to talk to an flowsim in server mode.`,
	Run: func(cmd *cobra.Command, args []string) {
		if TOS < 0 {
			fmt.Println("TOS needs to be >= 0")
			return
		}
		if TOS > 63 {
			fmt.Println("TOS needs to be < 64")
		}
		flow.Client(ip, port, iter, interval, iperf3_atoi(burstStr), TOS * 4)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	clientCmd.PersistentFlags().StringVarP(&ip, "ip", "I", "127.0.0.1", "IP address or host name of the flowsim server to talk to")
	clientCmd.PersistentFlags().IntVarP(&port, "port", "p", 8081, "TCP port of the flowsim server")
	clientCmd.PersistentFlags().IntVarP(&iter, "iter", "n", 6, "Number of bursts")
	clientCmd.PersistentFlags().IntVarP(&interval, "interval", "t", 10, "Interval in secs between bursts")
	clientCmd.PersistentFlags().StringVarP(&burstStr, "burst", "N", "1M", "Size of each burst (as x(.xxx)?[kmgtKMGT]?)")
	clientCmd.PersistentFlags().IntVarP(&TOS, "TOS", "T", 0, "Value of the TOS field in the IP packets (0 <= TOS < 64)")
}
