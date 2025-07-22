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

var lessonCmd = &cobra.Command{
	Use:   "lesson",
	Short: "Works with lessons",
	RunE: func(cmd *cobra.Command, args []string) error {
		lessons := musgitService.GetLessons()
		for i, lesson := range lessons {
			fmt.Printf("%d: \t%d\n", i, lesson.ID)
		}
		return nil

	},
}

var startLessonCmd = &cobra.Command{
	Use:   "start",
	Short: "Stars new lesson",
	Run: func(cmd *cobra.Command, args []string) {
		lesson, _ := musgitService.StartLesson()
		fmt.Println("lesson started:", lesson.ID)
	},
}

var pauseLessonCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pauses current lesson",
	Run: func(cmd *cobra.Command, args []string) {
		err := musgitService.PauseCurrentLesson()
		if err != nil {
			fmt.Println("Failed to pause lesson:", err)
		}

		fmt.Println("lesson paused")
	},
}

var resumeLessonCmd = &cobra.Command{
	Use:   "resume",
	Short: "Resumes current lesson",
	Run: func(cmd *cobra.Command, args []string) {
		err := musgitService.ResumeCurrentLesson()
		if err != nil {
			fmt.Println("Failed to resume lesson:", err)
		}

		fmt.Println("lesson resumed")
	},
}
var stopLessonCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops current lesson",
	Run: func(cmd *cobra.Command, args []string) {
		err := musgitService.StopCurrentLesson()
		if err != nil {
			fmt.Println("Failed to stop lesson:", err)
		}

		fmt.Println("lesson stopped")
	},
}

func init() {
	rootCmd.AddCommand(lessonCmd)
	lessonCmd.AddCommand(startLessonCmd)
	lessonCmd.AddCommand(resumeLessonCmd)
	lessonCmd.AddCommand(stopLessonCmd)
	lessonCmd.AddCommand(pauseLessonCmd)
}
