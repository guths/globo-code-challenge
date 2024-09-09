package morse

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func MorseCommand(root *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "morse",
		Short: "Command to start the api",
		Run: func(cmd *cobra.Command, args []string) {
			filePath, _ := cmd.Flags().GetString("file")
			outputFile, _ := cmd.Flags().GetString("output")

			var f *os.File

			if outputFile != "" {
				f, _ = os.Create(outputFile)
			}

			if filePath != "" {
				HandleFile(filePath, f)
				return
			}

			if len(args) == 0 {
				fmt.Println("the command need at least one argument")
				return
			}

			if len(args[0]) == 0 {
				fmt.Println("no morse code to read")
				return
			}

			HandleStdin(args[0], f)
		},
	}

	cmd.Flags().StringP("file", "f", "", "file need to be a txt")
	cmd.Flags().StringP("output", "o", "", "output file")

	return cmd
}
