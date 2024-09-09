package main

import (
	"github.com/guths/globo-code-challenge/pkg/morse"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(morse.MorseCommand(rootCmd))
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
