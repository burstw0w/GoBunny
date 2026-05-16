package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPullZonesBasic(apiKey string) ([]PullZoneBasic, error) {
	req, err := http.NewRequest("GET", "https://api.bunny.net/pullzone", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("AccessKey", apiKey)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("bunny API returned %d", resp.StatusCode)
	}

	var zones []PullZoneBasic
	err = json.NewDecoder(resp.Body).Decode(&zones)
	return zones, err
}

func GetPullZonesFull(apiKey string) ([]PullZoneFull, error) {
	req, err := http.NewRequest("GET", "https://api.bunny.net/pullzone", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("AccessKey", apiKey)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("bunny API returned %d", resp.StatusCode)
	}

	var zones []PullZoneFull
	err = json.NewDecoder(resp.Body).Decode(&zones)
	return zones, err
}

func CloneZone(apiKey string, sourceName string, newName string) error {
	zones, err := GetPullZonesFull(apiKey)
	if err != nil {
		return err
	}

	var source *PullZoneFull
	for i, z := range zones {
		if z.Name == sourceName {
			source = &zones[i]
			break
		}
	}
	if source == nil {
		return fmt.Errorf("zone '%s' not found", sourceName)
	}

	source.Name = newName
	source.Id = 0
	source.EdgeScriptId = 0
	source.MiddlewareScriptId = nil
	source.Hostnames = nil
	body, err := json.Marshal(source)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://api.bunny.net/pullzone", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("AccessKey", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create zone, API returned %d", resp.StatusCode, string(bodyBytes))
	}

	var newZone PullZoneFull
	err = json.NewDecoder(resp.Body).Decode(&newZone)
	if err != nil {
		return err
	}

	fmt.Printf("Created zone '%s' with ID %d\n", newZone.Name, newZone.Id)

	//err = AddHostname(apiKey, newZone.Id, hostname)
	//if err != nil {
	//    return err
	//}

	for _, rule := range source.EdgeRules {
		rule.Guid = ""
		err = AddEdgeRule(apiKey, newZone.Id, rule)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Done! Zone '%s' cloned successfully\n", newName)
	return nil
}

func AddHostname(apiKey string, zoneId int, hostname string) error {
	body, _ := json.Marshal(map[string]string{"Hostname": hostname})
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.bunny.net/pullzone/%d/addHostname", zoneId), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("AccessKey", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to add hostname, API returned %d", resp.StatusCode)
	}
	fmt.Printf("Added hostname '%s'\n", hostname)
	return nil
}

func LoadFreeCertificate(apiKey string, hostname string) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.bunny.net/pullzone/loadFreeCertificate?hostname=%s", hostname), nil)
	if err != nil {
		return err
	}
	req.Header.Set("AccessKey", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to load SSL certificate, API returned %d", resp.StatusCode)
	}
	fmt.Printf("SSL certificate provisioned for '%s'\n", hostname)
	return nil
}

func SetForceSSL(apiKey string, zoneId int, hostname string, force bool) error {
	body, _ := json.Marshal(map[string]interface{}{"Hostname": hostname, "ForceSSL": force})
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.bunny.net/pullzone/%d/setForceSSL", zoneId), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("AccessKey", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to set ForceSSL, API returned %d", resp.StatusCode)
	}
	state := "disabled"
	if force {
		state = "enabled"
	}
	fmt.Printf("Force SSL %s for '%s'\n", state, hostname)
	return nil
}

func RemoveHostname(apiKey string, zoneId int, hostname string) error {
	body, _ := json.Marshal(map[string]string{"Hostname": hostname})
	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://api.bunny.net/pullzone/%d/removeHostname", zoneId), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("AccessKey", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to remove hostname, API returned %d", resp.StatusCode)
	}
	fmt.Printf("Removed hostname '%s'\n", hostname)
	return nil
}

func AddEdgeRule(apiKey string, zoneId int, rule EdgeRuleFull) error {
	body, _ := json.Marshal(rule)
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.bunny.net/pullzone/%d/edgerules/addOrUpdate", zoneId), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("AccessKey", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to add edge rule '%s', API returned %d: %s", rule.Description, resp.StatusCode, string(bodyBytes))
	}
	fmt.Printf("Added edge rule '%s'\n", rule.Description)
	return nil
}

func DeleteEdgeRule(apiKey string, pullZoneId int, edgeRuleId string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://api.bunny.net/pullzone/%d/edgerules/%s", pullZoneId, edgeRuleId), nil)
	if err != nil {
		return err
	}
	req.Header.Set("AccessKey", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete edge rule, API returned %d", resp.StatusCode)
	}
	fmt.Println("Edge rule deleted")
	return nil
}

func SetEdgeRuleEnabled(apiKey string, pullZoneId int, edgeRuleId string, enabled bool) error {
	body, _ := json.Marshal(map[string]interface{}{"Id": pullZoneId, "Value": enabled})
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.bunny.net/pullzone/%d/edgerules/%s/setEdgeRuleEnabled", pullZoneId, edgeRuleId), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("AccessKey", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to set edge rule enabled, API returned %d", resp.StatusCode)
	}
	state := "disabled"
	if enabled {
		state = "enabled"
	}
	fmt.Printf("Edge rule %s\n", state)
	return nil
}

func GetRules(apiKey string, zoneId int) ([]EdgeRuleFull, error) {
	url := fmt.Sprintf("https://api.bunny.net/pullzone/%d", zoneId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("AccessKey", apiKey)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return nil, fmt.Errorf("Failed to read edge rules, error: %d", resp.StatusCode)
	}

	var zone PullZoneFull
	if err := json.NewDecoder(resp.Body).Decode(&zone); err != nil {
		return nil, err
	}

	return zone.EdgeRules, nil
}

func PurgeZone(apiKey string, zoneId int) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.bunny.net/pullzone/%d/purgeCache", zoneId), nil)
	if err != nil {
		return err
	}

	req.Header.Set("AccessKey", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return fmt.Errorf("API returned %d", resp.StatusCode)
	}
	return nil
}

func CreateZone(apiKey string, name string, originUrl string) (*PullZoneFull, error) {
	zone := PullZoneFull{
		Name:      name,
		OriginUrl: originUrl,
	}

	body, err := json.Marshal(zone)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.bunny.net/pullzone", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("AccessKey", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to create zone, API returned %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var newZone PullZoneFull
	err = json.NewDecoder(resp.Body).Decode(&newZone)
	if err != nil {
		return nil, err
	}

	return &newZone, nil
}

func DeleteZone(apiKey string, zoneId int) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://api.bunny.net/pullzone/%d", zoneId), nil)
	if err != nil {
		return err
	}

	req.Header.Set("AccessKey", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return fmt.Errorf("API returned %d", resp.StatusCode)
	}

	return nil
}

func UpdateZone(apiKey string, zoneId int, updates map[string]interface{}) error {
	body, err := json.Marshal(updates)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.bunny.net/pullzone/%d", zoneId), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("AccessKey", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to update zone, API returned %d: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}
