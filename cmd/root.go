/*
Copyright © 2025 Andrei Markin<andrey.markin.ppc@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/musgit-dev/musgit"

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
