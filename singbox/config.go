package singbox

import (
	"encoding/json"
	"os"
)

type AutoGenerated struct {
	AutoRoute                bool   `json:"auto_route"`
	DomainStrategy           string `json:"domain_strategy"`
	Inet4Address             string `json:"inet4_address"`
	Inet6Address             string `json:"inet6_address"`
	Mtu                      int    `json:"mtu"`
	Sniff                    bool   `json:"sniff"`
	SniffOverrideDestination bool   `json:"sniff_override_destination"`
	Stack                    string `json:"stack"`
	StrictRoute              bool   `json:"strict_route"`
	Type                     string `json:"type"`
}

func HandleProxyMode() error {
	config, err := os.ReadFile("./resources/core/config.json")
	if err != nil {
		return err
	}
	mapConfig := make(map[string]interface{})
	json.Unmarshal(config, &mapConfig)
	var newInbounds []interface{}
	for _, inbound := range mapConfig["inbounds"].([]interface{}) {
		inbound, _ := inbound.(map[string]interface{})
		if inbound["type"] == "mixed" {
			inbound["set_system_proxy"] = true
		}
		if inbound["type"] != "tun" {
			newInbounds = append(newInbounds, inbound)
		}
	}
	mapConfig["inbounds"] = newInbounds

	config, err = json.MarshalIndent(mapConfig, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("./resources/core/config.json", []byte(config), 0666)
	if err != nil {
		return err
	}

	return nil
}

func HandleTunMode() error {
	config, err := os.ReadFile("./resources/core/config.json")
	if err != nil {
		return err
	}
	TunConfig := AutoGenerated{
		AutoRoute:                true,
		DomainStrategy:           "prefer_ipv4",
		Inet4Address:             "172.19.0.1/30",
		Inet6Address:             "fdfe:dcba:9876::1/126",
		Mtu:                      9000,
		Sniff:                    true,
		SniffOverrideDestination: true,
		Stack:                    "mixed",
		StrictRoute:              true,
		Type:                     "tun",
	}
	mapConfig := make(map[string]interface{})
	json.Unmarshal(config, &mapConfig)
	isTunExist := false
	for _, inbound := range mapConfig["inbounds"].([]interface{}) {
		inbound, _ := inbound.(map[string]interface{})
		if inbound["type"] == "tun" {
			isTunExist = true
		}
		if inbound["type"] == "mixed" {
			inbound["set_system_proxy"] = false
		}
	}
	if !isTunExist {
		mapConfig["inbounds"] = append(mapConfig["inbounds"].([]interface{}), TunConfig)
	}

	config, err = json.MarshalIndent(mapConfig, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("./resources/core/config.json", []byte(config), 0666)
	if err != nil {
		return err
	}

	return nil
}