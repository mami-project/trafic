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
	"github.com/spf13/cobra"
	"github.com/mami-project/trafic/abrsim/abr"
)

var ip string
var port int
var interval int
var iter int
var burstStr string

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start abrsim in client mode",
	Long: `Will run abrsim in client mode
and try to talk to an abrsim in server mode.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("client called with params")
		// fmt.Printf("ip =   %s\n", ip)
		// fmt.Printf("port = %d\n", port)
		abr.Client(ip, port, iter, interval, int(iperf3_atof(burstStr)))
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	clientCmd.PersistentFlags().StringVarP(&ip, "ip", "I", "127.0.0.1", "IP address or host name of the abrsim server to talk to")
	clientCmd.PersistentFlags().IntVarP(&port, "port", "p", 8081, "TCP port of the abrsim server")
	clientCmd.PersistentFlags().IntVarP(&iter, "iter", "n", 6, "Number of bursts")
	clientCmd.PersistentFlags().IntVarP(&interval, "interval", "t", 10, "Interval in secs between bursts")
	clientCmd.PersistentFlags().StringVarP(&burstStr, "burst", "N", "1000000", "Size of each burst (as x(.xxx)?[kmgtKMGT]?)")
}
