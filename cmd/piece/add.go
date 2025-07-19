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
	"log"
	"musgit/internal/adapters/db"
	"musgit/internal/application/api"
	"musgit/internal/application/domain"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new piece to DB",
	RunE: func(cmd *cobra.Command, args []string) error {
		dbUri := viper.GetString("db-uri")
		name := viper.GetString("name")
		composer := viper.GetString("composer")
		complexity := viper.GetInt64("complexity")

		err := addPiece(
			dbUri,
			name,
			composer,
			domain.PieceComplexity(complexity),
		)
		return err
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("name", "n", "", "Name of the piece")
	addCmd.Flags().StringP("composer", "c", "", "Name of Composer")
	addCmd.Flags().Int64("complexity", 0, "Piece complexity")

	viper.BindPFlag("name", addCmd.Flags().Lookup("name"))
	viper.BindPFlag("composer", addCmd.Flags().Lookup("composer"))
	viper.BindPFlag("complexity", addCmd.Flags().Lookup("complexity"))
}

func addPiece(
	dbUri, name, composer string,
	complexity domain.PieceComplexity,
) error {
	dbAdapter, dbErr := db.NewAdapter(dbUri)
	if dbErr != nil {
		log.Fatalf("Failed to init db, err: %v", dbErr)
	}

	app := api.NewMusgitService(dbAdapter)
	piece, err := app.AddPiece(name, composer, complexity)
	log.Printf("Added new piece with id %d", piece.ID)
	return err
}
