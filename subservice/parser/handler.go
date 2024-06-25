package parser

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2/log"
	"github.com/valyala/fasthttp"
)

func fetchSubscribe(url string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func decodeSubscribe(body []byte) ([]string, error) {
	decodedBody, err := base64.StdEncoding.DecodeString(string(body))
	if err != nil {
		return nil, err
	}

	bodyStr := string(decodedBody)
	regex := regexp.MustCompile(`(ss|vmess|trojan|vless|hysteria|hysteria2|shadowtls)://([^\s]+)`)
	matches := regex.FindAllStringSubmatch(bodyStr, -1)

	var nodes []string
	for _, match := range matches {
		if len(match) >= 1 {
			nodes = append(nodes, match[0])
		}
	}

	return nodes, nil
}

func parseSubscribe(links []string) ([]map[string]interface{}, error) {
	var outbounds []map[string]interface{}
	for _, link := range links {
		prot := strings.Split(link, "://")

		switch prot[0] {
		case "ss":
			node, err := Shadowsocks(link)
			if err != nil {
				return nil, err
			}
			outbounds = append(outbounds, node)
		case "vmess":
			node, err := Vmess(link)
			if err != nil {
				return nil, err
			}
			outbounds = append(outbounds, node)
		}
	}
	return outbounds, nil
}

func Handler(url string) ([]byte, error) {
	req, err := fetchSubscribe(url)
	if err != nil {
		return nil, err
	}
	nodes, err := decodeSubscribe(req)
	if err != nil {
		return nil, err
	}
	outbounds, err := parseSubscribe(nodes)
	if err != nil {
		log.Error("parseSubscribe error:", err)
		return nil, err
	}
	template, err := os.ReadFile("./resources/template/template.json")
	if err != nil {
		return nil, err
	}
	mapTemplate := make(map[string]interface{})
	json.Unmarshal(template, &mapTemplate)

	// add tag
	for _, i := range mapTemplate["outbounds"].([]interface{}) {
		m, _ := i.(map[string]interface{})
		if m["tag"] == "proxy" || m["tag"] == "auto" {
			for _, f := range outbounds {
				m["outbounds"] = append(m["outbounds"].([]interface{}), f["tag"])
			}
		}
	}

	// add outbound
	for _, i := range outbounds {
		mapTemplate["outbounds"] = append(mapTemplate["outbounds"].([]interface{}), i)
	}

	config, err := json.MarshalIndent(mapTemplate, "", "  ")
	if err != nil {
		return nil, err
	}

	return config, nil
}
