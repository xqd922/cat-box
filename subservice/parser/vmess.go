package parser

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"
)

// ws
type ws struct {
	Type    string `json:"type"`
	Headers struct {
		Host string `json:"Host"`
	} `json:"headers"`
	Path string `json:"path"`
}

func Vmess(link string) (map[string]interface{}, error) {
	// Base64 decode
	linkSplit := strings.Split(link, "://")
	decodedLink, _ := base64.StdEncoding.DecodeString(linkSplit[1])
	mapLink := make(map[string]interface{})
	json.Unmarshal(decodedLink, &mapLink)

	// String to int
	portStr := mapLink["port"].(string)
	port, _ := strconv.Atoi(portStr)
	alter_idStr := mapLink["aid"].(string)
	alter_id, _ := strconv.Atoi(alter_idStr)

	// Generate configuration
	vmess := make(map[string]interface{})
	vmess["tag"] = mapLink["ps"]
	vmess["type"] = "vmess"
	vmess["server"] = mapLink["add"]
	vmess["server_port"] = port
	vmess["uuid"] = mapLink["id"]
	vmess["security"] = "auto"
	vmess["alter_id"] = alter_id
	vmess["packet_encoding"] = "xudp"

	// Handling plugins
	switch mapLink["net"] {
	case "ws":
		websocket := ws{
			Type: "ws",
			Headers: struct {
				Host string `json:"Host"`
			}{
				Host: mapLink["host"].(string),
			},
			Path: mapLink["path"].(string),
		}
		vmess["transport"] = websocket
	}
	return vmess, nil
}
