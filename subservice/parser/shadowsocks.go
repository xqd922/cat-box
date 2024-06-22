package parser

import (
	"encoding/base64"
	"errors"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func Shadowsocks(link string) (map[string]interface{}, error) {
	reg := regexp.MustCompile(`ss://([A-Za-z0-9_-]+)@([^:#]+):(\d+).*?#(.*?)$`)
	match := reg.FindStringSubmatch(link)
	if match == nil || len(match) < 5 {
		return nil, errors.New("invalid shadowsocks link format")
	}

	// Parse tag, server, port, method, password
	port, _ := strconv.Atoi(match[3])
	tag, _ := url.QueryUnescape(match[4])
	server := match[2]
	authPassByte, _ := base64.StdEncoding.DecodeString(match[1])
	authPass := strings.Split(string(authPassByte), ":")
	method := authPass[0]
	password := authPass[1]

	// Generate configuration
	ss := make(map[string]interface{})
	ss["tag"] = tag
	ss["type"] = "shadowsocks"
	ss["server"] = server
	ss["server_port"] = port
	ss["method"] = method
	ss["password"] = password
	return ss, nil
}
