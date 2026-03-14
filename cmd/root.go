/*
Copyright © 2026 Zoom theoldzoom@proton.me

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"goip/internal/ipinfo"
	"goip/internal/output"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var jsonOutput bool

var rootCmd = &cobra.Command{
	Use:   "goip",
	Short: "A tool to see an IP address & get info about an IP address",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			ip, err := ipinfo.GetMyIP()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			if jsonOutput {
				json, err := json.Marshal(ip)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					return
				}
				fmt.Println(string(json))
				return
			}
			fmt.Println(output.FormatIPInfo(ip))
			return
		}

		ip := args[0]
		info, err := ipinfo.GetIPInfo(ip)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		if jsonOutput {
			json, err := json.Marshal(info)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			fmt.Println(string(json))
			return
		}
		fmt.Println(output.FormatIPInfo(info))

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/.goip)")

	rootCmd.Flags().BoolVar(&jsonOutput, "json", false, "output JSON")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		configDir, err := os.UserConfigDir()
		cobra.CheckErr(err)
		viper.SetConfigFile(filepath.Join(configDir, ".goip"))
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
