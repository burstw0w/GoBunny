package cmd

import (
	"GoBunny/api"
	"encoding/json"
	"fmt"
	"os"
	"strings"

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

var infoAll = &cobra.Command{
	Use:   "infoAll",
	Short: "Returns a JSON of all your pullzones and their configurations",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("BUNNY_API_KEY")
		if apiKey == "" {
			fmt.Println("Error: BUNNY_API_KEY env variable not set")
			os.Exit(1)
		}

		zones, err := api.GetPullZonesFull(apiKey)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		if len(zones) == 0 {
			fmt.Println("No pull zones found.")
			os.Exit(1)
		}
		printZonesJSON(zones)
	},
}

func init() {
	pullzonesCmd.AddCommand(cloneCmd)
	pullzonesCmd.AddCommand(infoAll)
	rootCmd.AddCommand(pullzonesCmd)
}

/*func printZones(zones []api.PullZoneFull) {
	for _, z := range zones {
		fmt.Printf("● %s (ID: %d)\n", z.Name, z.Id)

		fmt.Printf("  ├─ Status:      %s\n", formatStatus(z.Enabled, z.Suspended))
		fmt.Printf("  ├─ Origin:      %s\n", z.OriginUrl)

		if len(z.Hostnames) > 0 {
			fmt.Printf("  ├─ Hostname:    %s\n", z.Hostnames[0].Value)
		}

		usageGB := float64(z.MonthlyBandwidthUsed) / (1024 * 1024 * 1024)
		fmt.Printf("  ├─ Bandwidth:   %.2f GB / %d GB\n", usageGB, z.MonthlyBandwidthLimit / (1024 * 1024 * 1024))

		fmt.Printf("  ├─ Smart Cache: %v\n", z.EnableSmartCache)
		fmt.Printf("  ├─ WAF:         %v\n", z.ZoneSecurityEnabled)
		fmt.Printf("  ├─ TLS 1.0/1.1: %v (Legacy)\n", z.EnableTLS1 || z.EnableTLS1_1)

		if z.EnableOriginShield {
			fmt.Printf("  └─ Shield:      ENABLED (%s)\n", z.OriginShieldZoneCode)
		} else {
			fmt.Printf("  └─ Shield:      DISABLED\n")
		}

		fmt.Println("")
	}
}*/

func printZonesJSON(zones []api.PullZoneFull) {
	for _, z := range zones {
		data, err := json.MarshalIndent(z, "", "  ")
		if err != nil {
			fmt.Printf("Error encoding zone %s: %v\n", z.Name, err)
			continue
		}

		fmt.Printf("--- FULL DATA FOR: %s ---\n", z.Name)
		fmt.Println(string(data))
		fmt.Println(strings.Repeat("=", 40))
	}
}

func formatStatus(enabled, suspended bool) string {
	if suspended {
		return "SUSPENDED"
	}
	if enabled {
		return "ACTIVE"
	}
	return "DISABLED"
}
