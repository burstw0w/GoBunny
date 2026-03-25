package main

import (
	"fmt"
	"os"
	"GoBunny/api"
)

func main() {
	apiKey := os.Getenv("BUNNY_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: BUNNY_API_KEY env variable not set")
		os.Exit(1)
	}

	/*zones, err := api.GetPullZones(apiKey)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for _, zone := range zones {
		fmt.Printf("Zone: %s (ID: %d) | Enabled: %v\n", zone.Name, zone.Id, zone.Enabled)
		for _, host := range zone.Hostnames {
			fmt.Printf("  hostname: %s\n", host.Value)
		}
		for _, rule := range zone.EdgeRules {
			fmt.Printf("  edge rule: %s (enabled: %v)\n", rule.Description, rule.Enabled)
		}
	}*/
	err := api.CloneZone(apiKey, "burstx86", "burstx86-clone", "burstx86")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
