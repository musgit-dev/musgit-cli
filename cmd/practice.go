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

	"github.com/spf13/cobra"
)

var practiceCmd = &cobra.Command{
	Use:   "practice",
	Short: "Shows information on practices",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("practice called")
	},
}
var stopPracticeCmd = &cobra.Command{
	Use:   "stop <practice_id>",
	Short: "Stops selected practice",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("practice stopped")
	},
}
var evalutePracticeCmd = &cobra.Command{
	Use:   "evalute <practice_id> <evaluation>",
	Short: "Provide evaluation for the practice",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("practice evaluated")
	},
}

func init() {
	rootCmd.AddCommand(practiceCmd)
	practiceCmd.AddCommand(stopPracticeCmd)
	practiceCmd.AddCommand(evalutePracticeCmd)
}
