package tfl

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`(?m)^([A-Za-z-'&\.\s]+)`)

func GetStationShortname(longName string) string {
	shortName := longName

	if strings.HasSuffix(shortName, " Underground Station") {
		shortName = strings.TrimSuffix(shortName, " Underground Station")
	}

	if strings.HasSuffix(shortName, " Rail Station") {
		shortName = strings.TrimSuffix(shortName, " Rail Station")
	}

	if strings.HasSuffix(shortName, " DLR Station") {
		shortName = strings.TrimSuffix(shortName, " DLR Station")
	}

	matches := re.FindStringSubmatch(shortName)

	return matches[1]
}
