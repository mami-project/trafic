package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "schedule",
	Short: "schedule",
	Long:  "schedule",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	pflags := rootCmd.PersistentFlags()
	pflags.StringVar(&cfgFile, "config", "/etc/trafic.yaml", "configuration file")

	pflags.String("log-tag", "", "a tag that is prepended to all log lines")
	viper.BindPFlag("log.tag", pflags.Lookup("log-tag"))

	pflags.String("flows-dirs", "", "comma separated folder(s) with flow configuration files")
	viper.BindPFlag("flows.dirs", pflags.Lookup("flows-dirs"))

	pflags.String("scheduler-tick", "250ms", "scheduler granularity")
	viper.BindPFlag("scheduler.tick", pflags.Lookup("scheduler-tick"))

	pflags.String("stats-dir", "/var/trafic/stats", "Folder where to stash the collected samples")
	viper.BindPFlag("stats.dir", pflags.Lookup("stats-dir"))

	pflags.Bool("influxdb-enabled", false, "(also) forward the collected samples to an InfluxDB instance")
	viper.BindPFlag("influxdb.enabled", pflags.Lookup("influxdb-enabled"))

	pflags.String("influxdb-endpoint", "http://localhost:8086", "InfluxDB endpoint")
	viper.BindPFlag("influxdb.endpoint", pflags.Lookup("influxdb-endpoint"))

	pflags.String("influxdb-db", "trafic", "name of the InfluxDB database where to stash our samples")
	viper.BindPFlag("influxdb.db", pflags.Lookup("influxdb-db"))

	pflags.String("influxdb-measurements", "samples", "name of the InfluxDB measurements")
	viper.BindPFlag("influxdb.measurements", pflags.Lookup("influxdb-measurements"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(filepath.Join("/etc", "/trafic", ".trafic"))
		viper.SetConfigName("trafic.yaml")
	}

	viper.SetEnvPrefix("TRAFIC")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
