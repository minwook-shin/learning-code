package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	StringVar := ""
	StringVarP := ""
	rootCmd.PersistentFlags().StringVar(&StringVar, "stringVar", "", "StringVar defines a string flag")
	rootCmd.PersistentFlags().StringVarP(&StringVarP, "stringVarP", "v", "", "like StringVar, but accepts a shorthand letter")
	rootCmd.PersistentFlags().String("String", "", "String defines a string flag")
	rootCmd.PersistentFlags().StringP("stringP", "p", "", "like String, but accepts a shorthand letter")
}

var rootCmd = &cobra.Command{
	Use:   "main.go",
	Short: "A 'Hello, World!' program generally is a computer program that outputs or displays the message 'Hello, World!'.",
	Long: `A "Hello, World!" program is 
traditionally used to introduce 
novice programmers to a programming language.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("hello world!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
