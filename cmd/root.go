/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/octacian/timecalc/time"
	"github.com/peterh/liner"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "timecalc [expression?]",
	Short: "timecalc is a calculator built to work with quantities of time",
	Long: `timecalc is a calculator built to work with quantities of time.
warning: order of operations is not complete outside of basic groups.

Example usage:
>>> 10 + (((7 * 81) - 42) / 20)
36.25
>> 10 + ((7 * 81) - 42) / 20
26.75 # Demonstrates lack of order of operations
>>> .5 - 1.7
-1.2
>>> ::30.5 * 3
00:01:31.5
>>> :30 % (8 * 1000 * 60)
00:06
>>> : + 1
00:00:00.001 # Ones place represents milliseconds`,
	Run: func(cmd *cobra.Command, args []string) {
		evalMode, _ := cmd.Flags().GetBool("eval")

		if len(args) > 0 {
			input := strings.Join(args, " ")
			if !evalMode {
				fmt.Printf(">>> %s\n", input)
			}

			result, err := time.FromString(input)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(result)
		}

		if evalMode {
			return
		}

		line := liner.NewLiner()
		defer line.Close()

		line.SetCtrlCAborts(true)

		for {
			input, err := line.Prompt(">>> ")
			if err != nil {
				if err == liner.ErrPromptAborted {
					return
				}
				log.Fatalf("error reading line: %s", err)
			}

			if input == "exit" {
				return
			}

			line.AppendHistory(input)

			result, err := time.FromString(input)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(result)
		}
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
	rootCmd.Flags().BoolP("eval", "e", false, "disables repl and expects query as command arguments")
}
