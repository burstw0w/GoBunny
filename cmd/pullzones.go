package cmd

import (
	"fmt"
	"os"
	"GoBunny/api"

	"github.com/spf13/cobra"
)

var pullzonesCmd = &cobra.Command{
	Use:   "pullzones",
	Short: "Manage BunnyCDN Pull Zones",
}

var cloneCmd = &cobra.Command{
	Use:   "clone [source] [target] [hostname]",
	Short: "Clone a pull zone by name",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("BUNNY_API_KEY")
		if apiKey == "" {
			fmt.Println("Error: BUNNY_API_KEY env variable not set")
			os.Exit(1)
		}

		source := args[0]
		target := args[1]
		host := args[2]

		err := api.CloneZone(apiKey, source, target, host)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Successfully cloned %s to %s\n", source, target)
	},
}

func init() {
	pullzonesCmd.AddCommand(cloneCmd)
	rootCmd.AddCommand(pullzonesCmd)
}