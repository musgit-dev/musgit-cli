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
	"strconv"

	"github.com/musgit-dev/musgit/models"
	"github.com/spf13/cobra"
)

var pieceCmd = &cobra.Command{
	Use:   "piece",
	Short: "Working with pieces",
	Run: func(cmd *cobra.Command, args []string) {
		for i, piece := range musgitService.GetPieces() {
			fmt.Printf("%d: \t%s\n", i, piece.Name)
		}
	},
}
var showPieceCmd = &cobra.Command{
	Use:   "show <piece_id>",
	Short: "Shows information about the piece",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pieceId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		piece, err := musgitService.GetPiece(pieceId)
		if err != nil {
			fmt.Println("Unknown Piece Id:", pieceId)
		}
		fmt.Println(piece)
	},
}
var addPieceCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds new piece",
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flag("name").Value.String()
		composer := cmd.Flag("name").Value.String()
		complexity, _ := strconv.Atoi(cmd.Flag("name").Value.String())
		piece, err := musgitService.AddPiece(
			name,
			composer,
			models.PieceComplexity(complexity),
		)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Added new piece: ", piece.ID)
	},
}
var practicePieceCmd = &cobra.Command{
	Use:   "practice <piece_id>",
	Short: "Starts practice for a selected piece",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pieceId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		lessonId, _ := strconv.ParseInt(
			cmd.Flag("lesson").Value.String(),
			10,
			64,
		)
		practice, err := musgitService.PracticePiece(pieceId, lessonId)
		if err != nil {
			fmt.Println("Failed to start practice :", err)
		}
		fmt.Println("Started practice: ", practice.ID)
	},
}

func init() {
	rootCmd.AddCommand(pieceCmd)
	pieceCmd.AddCommand(showPieceCmd)
	pieceCmd.AddCommand(addPieceCmd)
	pieceCmd.AddCommand(practicePieceCmd)

	addPieceCmd.Flags().String("name", "", "Name of a piece")
	addPieceCmd.Flags().String("composer", "", "Composer")
	addPieceCmd.Flags().Int64("complexity", 0, "Complexity of a piece")

	practicePieceCmd.Flags().Int64("lesson", 1, "Lesson ID")
}
