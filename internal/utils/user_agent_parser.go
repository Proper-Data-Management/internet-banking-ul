package utils

import "strings"

func ParseUserAgent(userAgent string) (deviceCategory string, os string, appVersion string) {
	fields := strings.Split(userAgent, "/")
	if len(fields) > 0 {
		deviceCategory = fields[0]
	}
	if len(fields) > 2 {
		os = fields[2]
	}
	if len(fields) > 3 {
		appVersion = fields[3]
	}
	return
}
