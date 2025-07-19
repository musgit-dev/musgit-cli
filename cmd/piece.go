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

	"github.com/spf13/cobra"
)

var pieceCmd = &cobra.Command{
	Use:   "piece",
	Short: "Working with pieces",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("piece called")
	},
}
var showPieceCmd = &cobra.Command{
	Use:   "show <piece_id>",
	Short: "Shows information about the piece",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("piece shown")
	},
}
var addPieceCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds new piece",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("piece added")
	},
}
var practicePieceCmd = &cobra.Command{
	Use:   "practice <piece_id>",
	Short: "Starts practice for a selected piece",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("piece practiced")
	},
}

func init() {
	rootCmd.AddCommand(pieceCmd)
	pieceCmd.AddCommand(showPieceCmd)
	pieceCmd.AddCommand(addPieceCmd)
	pieceCmd.AddCommand(practicePieceCmd)

}
