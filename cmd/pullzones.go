package cmd

import (
	"GoBunny/api"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var forceFlag bool

var pullzonesCmd = &cobra.Command{
	Use:   "pullzones",
	Short: "Manage BunnyCDN Pull Zones",
}

var cloneCmd = &cobra.Command{
	Use:   "clone [source] [target] [hostname]",
	Short: "Clone a pull zone by name",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("BUNNY_API_KEY")
		if apiKey == "" {
			fmt.Println("Error: BUNNY_API_KEY env variable not set")
			os.Exit(1)
		}

		source := args[0]
		target := args[1]
		//host := args[2]

		err := api.CloneZone(apiKey, source, target)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Successfully cloned %s to %s\n", source, target)
	},
}

var rulesCmd = &cobra.Command{
	Use:   "rules [source]",
	Short: "Print Edge Rules of source zone",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("BUNNY_API_KEY")
		if apiKey == "" {
			fmt.Println("Error: BUNNY_API_KEY env variable not set")
			os.Exit(1)
		}

		source := args[0]

		zones, _ := api.GetPullZonesBasic(apiKey)
		var targetId int
		for _, z := range zones {
			if z.Name == source {
				targetId = z.Id
				break
			}
		}

		rules, err := api.GetRules(apiKey, targetId)
		if err != nil {
			fmt.Printf("Error %v\n", err)
			os.Exit(1)
		}
		if len(rules) == 0 {
			fmt.Println("No edge rules found.")
			os.Exit(1)
		}
		printRulesJSON(rules)
	},
}

var copyRulesCmd = &cobra.Command{
	Use:   "copy [source] [target]",
	Short: "Copies Edge rules from source pullzone to target pullzone",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("BUNNY_API_KEY")
		if apiKey == "" {
			fmt.Println("Error: BUNNY_API_KEY env variable not set")
			os.Exit(1)
		}

		httpClient := &http.Client{}

		source := args[0]
		target := args[1]

		zones, err := api.GetPullZonesBasic(apiKey)
		if err != nil {
			fmt.Printf("Error fetching zones: %v\n", err)
			os.Exit(1)
		}

		var sourceId, targetId int
		for _, z := range zones {
			if z.Name == source {
				sourceId = z.Id
			}
			if z.Name == target {
				targetId = z.Id
			}
		}

		if sourceId == 0 || targetId == 0 {
			fmt.Println("Error: Could not find source or target zone.")
			os.Exit(1)
		}

		rules, err := api.GetRules(apiKey, sourceId)
		if err != nil {
			fmt.Printf("Error fetching source rules: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Found %d rules in %s. Copying to %s...\n", len(rules), source, target)

		for _, rule := range rules {
			rule.Guid = ""
			err := api.AddEdgeRule(apiKey, targetId, rule, httpClient)
			if err != nil {
				fmt.Printf("  [!] Failed to copy rule '%s': %v\n", rule.Description, err)
			} else {
				fmt.Printf("  [✔] Copied: %s\n", rule.Description)
			}
		}

		fmt.Println("Done.")

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

var purgeCmd = &cobra.Command{
	Use:   "purge [zone]",
	Short: "Purge all cached files for a zone",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("BUNNY_API_KEY")
		if apiKey == "" {
			fmt.Println("Error: BUNNY_API_KEY not set")
			os.Exit(1)
		}

		zones, _ := api.GetPullZonesBasic(apiKey)
		var zoneId int
		for _, z := range zones {
			if z.Name == args[0] {
				zoneId = z.Id
				break
			}
		}

		if zoneId == 0 {
			fmt.Println("Zone not found")
			os.Exit(1)
		}

		if err := api.PurgeZone(apiKey, zoneId); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Cache purged for %s\n", args[0])
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [zone]",
	Short: "Deletes a pullzone",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("BUNNY_API_KEY")
		if apiKey == "" {
			fmt.Println("Error: BUNNY_API_KEY not set")
			os.Exit(1)
		}

		zones, _ := api.GetPullZonesBasic(apiKey)
		var zoneId int
		for _, z := range zones {
			if z.Name == args[0] {
				zoneId = z.Id
				break
			}
		}

		if zoneId == 0 {
			fmt.Println("Zone not found")
			os.Exit(1)
		}

		if !forceFlag {
			fmt.Printf("This will PERMANENTLY delete '%s'. Type '%s' to confirm: ", args[0], args[0])
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != args[0] {
				fmt.Println("Aborted.")
				os.Exit(0)
			}
		}

		if err := api.DeleteZone(apiKey, zoneId); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Zone '%s' deleted\n", args[0])
	},
}

var createCmd = &cobra.Command{
	Use:   "create [zone] [origin]",
	Short: "Creates a pullzone",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("BUNNY_API_KEY")
		if apiKey == "" {
			fmt.Println("Error: BUNNY_API_KEY not set")
			os.Exit(1)
		}

		name := args[0]
		origin := args[1]
		if !strings.HasPrefix(origin, "http://") && !strings.HasPrefix(origin, "https://") {
			origin = "https://" + origin
		}

		zones, err := api.GetPullZonesBasic(apiKey)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		for _, z := range zones {
			if z.Name == name {
				fmt.Printf("Zone '%s' already exists\n", name)
				os.Exit(1)
			}
		}

		_, err = api.CreateZone(apiKey, name, origin)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Zone '%s' created\n", name)
	},
}

var hostnameCmd = &cobra.Command{
	Use:   "hostname",
	Short: "Manage hostnames for a pull zone",
}

var hostnameAddCmd = &cobra.Command{
	Use:   "add [zone] [hostname]",
	Short: "Add a hostname to a pull zone",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("BUNNY_API_KEY")
		if apiKey == "" {
			fmt.Println("Error: BUNNY_API_KEY not set")
			os.Exit(1)
		}

		zones, _ := api.GetPullZonesBasic(apiKey)
		var zoneId int
		for _, z := range zones {
			if z.Name == args[0] {
				zoneId = z.Id
				break
			}
		}
		if zoneId == 0 {
			fmt.Println("Zone not found")
			os.Exit(1)
		}

		if err := api.AddHostname(apiKey, zoneId, args[1]); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

var hostnameListCmd = &cobra.Command{
	Use:   "list [zone]",
	Short: "List hostnames for a pull zone",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("BUNNY_API_KEY")
		if apiKey == "" {
			fmt.Println("Error: BUNNY_API_KEY not set")
			os.Exit(1)
		}

		zones, _ := api.GetPullZonesFull(apiKey)
		for _, z := range zones {
			if z.Name == args[0] {
				if len(z.Hostnames) == 0 {
					fmt.Println("No hostnames found.")
					os.Exit(0)
				}
				fmt.Printf("%-4s %-45s %-5s %-6s %s\n", "ID", "Hostname", "SSL", "Force SSL", "Type")
				fmt.Println(strings.Repeat("-", 75))
				for _, h := range z.Hostnames {
					cert := "no"
					if h.HasCertificate {
						cert = "yes"
					}
					force := "no"
					if h.ForceSSL {
						force = "yes"
					}
					hType := "custom"
					if h.IsSystemHostname {
						hType = "system"
					} else if h.IsManagedHostname {
						hType = "managed"
					}
					fmt.Printf("%-4d %-45s %-5s %-6s %s\n", h.Id, h.Value, cert, force, hType)
				}
				return
			}
		}
		fmt.Println("Zone not found")
		os.Exit(1)
	},
}

var hostnameCertCmd = &cobra.Command{
	Use:   "cert [zone] [hostname]",
	Short: "Provision a free Let's Encrypt SSL certificate for a hostname",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("BUNNY_API_KEY")
		if apiKey == "" {
			fmt.Println("Error: BUNNY_API_KEY not set")
			os.Exit(1)
		}

		zones, _ := api.GetPullZonesBasic(apiKey)
		var found bool
		for _, z := range zones {
			if z.Name == args[0] {
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Zone not found")
			os.Exit(1)
		}

		if err := api.LoadFreeCertificate(apiKey, args[1]); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

var hostnameForceSSLCmd = &cobra.Command{
	Use:   "forcessl [zone] [hostname] [on|off]",
	Short: "Enable or disable Force SSL on a hostname",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("BUNNY_API_KEY")
		if apiKey == "" {
			fmt.Println("Error: BUNNY_API_KEY not set")
			os.Exit(1)
		}

		zones, _ := api.GetPullZonesBasic(apiKey)
		var zoneId int
		for _, z := range zones {
			if z.Name == args[0] {
				zoneId = z.Id
				break
			}
		}
		if zoneId == 0 {
			fmt.Println("Zone not found")
			os.Exit(1)
		}

		force := strings.ToLower(args[2]) == "on"
		if err := api.SetForceSSL(apiKey, zoneId, args[1], force); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

var hostnameRemoveCmd = &cobra.Command{
	Use:   "remove [zone] [hostname]",
	Short: "Remove a hostname from a pull zone",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("BUNNY_API_KEY")
		if apiKey == "" {
			fmt.Println("Error: BUNNY_API_KEY not set")
			os.Exit(1)
		}

		zones, _ := api.GetPullZonesBasic(apiKey)
		var zoneId int
		for _, z := range zones {
			if z.Name == args[0] {
				zoneId = z.Id
				break
			}
		}
		if zoneId == 0 {
			fmt.Println("Zone not found")
			os.Exit(1)
		}

		if err := api.RemoveHostname(apiKey, zoneId, args[1]); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	deleteCmd.Flags().BoolVarP(&forceFlag, "yes", "y", false, "Skip confirmation")
	pullzonesCmd.AddCommand(cloneCmd)
	pullzonesCmd.AddCommand(infoAll)
	pullzonesCmd.AddCommand(rulesCmd)
	pullzonesCmd.AddCommand(purgeCmd)
	pullzonesCmd.AddCommand(deleteCmd)
	pullzonesCmd.AddCommand(createCmd)
	pullzonesCmd.AddCommand(hostnameCmd)
	hostnameCmd.AddCommand(hostnameAddCmd)
	hostnameCmd.AddCommand(hostnameListCmd)
	hostnameCmd.AddCommand(hostnameCertCmd)
	hostnameCmd.AddCommand(hostnameForceSSLCmd)
	hostnameCmd.AddCommand(hostnameRemoveCmd)
	rulesCmd.AddCommand(copyRulesCmd)
	rootCmd.AddCommand(pullzonesCmd)
	rootCmd.AddCommand(rulesCmd)
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

func printRulesJSON(rules []api.EdgeRuleFull) {
	fmt.Printf("=== EDGE RULES ===\n")
	for _, z := range rules {
		data, err := json.MarshalIndent(z, "", "  ")
		if err != nil {
			fmt.Printf("Error encoding rules")
			continue
		}
		fmt.Printf("-- RULE NAME '%s' --\n", z.Description)
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
