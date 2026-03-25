package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"io"
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

func GetPullZonesFull(apiKey string) ([]PullZoneFull, error){
	req, err := http.NewRequest("GET", "https://api.bunny.net/pullzone", nil)
	if err != nil{
		return nil, err
	}

	req.Header.Set("AccessKey", apiKey)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil{
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("bunny API returned %d", resp.StatusCode)
	}

	var zones []PullZoneFull
	err = json.NewDecoder(resp.Body).Decode(&zones)
	return zones, err;
}

func CloneZone(apiKey string, sourceName string, newName string, hostname string) error {
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

    //err = addHostname(apiKey, newZone.Id, hostname, client)
    //if err != nil {
    //    return err
    //}

    for _, rule := range source.EdgeRules {
        rule.Guid = "" 
        err = addEdgeRule(apiKey, newZone.Id, rule, client)
        if err != nil {
            return err
        }
    }

    fmt.Printf("Done! Zone '%s' cloned successfully\n", newName)
    return nil
}

func addHostname(apiKey string, zoneId int, hostname string, client *http.Client) error {
    body, _ := json.Marshal(map[string]string{"Hostname": hostname})
    req, err := http.NewRequest("POST", fmt.Sprintf("https://api.bunny.net/pullzone/%d/addHostname", zoneId), bytes.NewBuffer(body))
    if err != nil {
        return err
    }
    req.Header.Set("AccessKey", apiKey)
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return fmt.Errorf("failed to add hostname, API returned %d", resp.StatusCode)
    }
    fmt.Printf("Added hostname '%s'\n", hostname)
    return nil
}

func addEdgeRule(apiKey string, zoneId int, rule EdgeRuleFull, client *http.Client) error {
	body, _ := json.Marshal(rule)
    req, err := http.NewRequest("POST", fmt.Sprintf("https://api.bunny.net/pullzone/%d/edgerules/addOrUpdate", zoneId), bytes.NewBuffer(body))
    if err != nil {
        return err
    }
    req.Header.Set("AccessKey", apiKey)
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 && resp.StatusCode != 201 {
		bodyBytes, _ := io.ReadAll(resp.Body)
        return fmt.Errorf("failed to add edge rule '%s', API returned %d", rule.Description, resp.StatusCode, string(bodyBytes))
    }
    fmt.Printf("Copied edge rule '%s'\n", rule.Description)
    return nil
}


