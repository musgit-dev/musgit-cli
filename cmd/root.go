/*
Copyright © 2025 Andrei Markin

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
	"fmt"
	"os"

	"musgit"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dbUri string
var musgitService *musgit.MusgitService

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "musgit-cli",
	Short:   "Musgit CLI client.",
	Version: "0.0.0",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if dbUri == "" {
			return fmt.Errorf("Missing --db-uri flag.")
		}
		musgitService = musgit.NewMusgitService(dbUri)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().
		StringVarP(&dbUri, "db-uri", "d", "", "DB with data")
	err := viper.BindPFlag("db-uri", rootCmd.PersistentFlags().Lookup("db-uri"))
	if err != nil {
		fmt.Println("DB URI is not provided.")
	}
}
