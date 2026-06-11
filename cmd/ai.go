/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var name string

// aiCmd represents the ai command
var aiCmd = &cobra.Command{
	Use:   "ai",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains.`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print("Enter your name (type exit to stop): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == "exit" {
				fmt.Println("Bye 👋")
				break
			}

			fmt.Println("Hello", input)

		}
	},
}

func init() {
	rootCmd.AddCommand(aiCmd)
	aiCmd.Flags().StringVarP(&name, "name", "n", "Abdul Alim", "user name")
}
