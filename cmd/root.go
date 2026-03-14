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
	"io"
	"os"

	"github.com/spf13/cobra"
)

var jsonOutput bool

type errorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func printJSON(w io.Writer, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, string(data))
	return nil
}

var rootCmd = &cobra.Command{
	Use:       "goip",
	Short:     "Show your public IP or look up details for any IP address or domain name",
	ValidArgs: []string{"<ip>"},
	Long: `goip shows information about your current public IP when run without arguments.
If you pass an IP address or domain name, it looks up details for that target instead.
Use --json for machine-readable output.
Supports both IPv4 and IPv6 addresses, and domain names.`,
	Example: `goip
goip 1.1.1.1
goip --json
goip 6.7.6.7 --json
goip example.com
goip example.com --json`,
	SilenceErrors: true,
	SilenceUsage:  true,

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			ip, err := ipinfo.GetMyIP()
			if err != nil {
				if jsonOutput && ip.Status != "" {
					if jsonErr := printJSON(os.Stderr, ip); jsonErr == nil {
						return err
					}
				} else if jsonOutput {
					if jsonErr := printJSON(os.Stderr, errorResponse{Status: "fail", Message: err.Error()}); jsonErr == nil {
						return err
					}
				}
				fmt.Fprintln(os.Stderr, err)
				return err
			}

			if jsonOutput {
				if err := printJSON(os.Stdout, ip); err != nil {
					fmt.Fprintln(os.Stderr, err)
					return err
				}
				return nil
			}
			fmt.Println(output.FormatIPInfo(ip))
			return nil
		}
		ip := args[0]
		info, err := ipinfo.GetIPInfo(ip)
		if err != nil {
			if jsonOutput && info.Status != "" {
				if jsonErr := printJSON(os.Stderr, info); jsonErr == nil {
					return err
				}
			} else if jsonOutput {
				if jsonErr := printJSON(os.Stderr, errorResponse{Status: "fail", Message: err.Error()}); jsonErr == nil {
					return err
				}
			}
			fmt.Fprintln(os.Stderr, err)
			return err
		}
		if jsonOutput {
			if err := printJSON(os.Stdout, info); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return err
			}
			return nil
		}
		fmt.Println(output.FormatIPInfo(info))

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVar(&jsonOutput, "json", false, "output JSON")
}
