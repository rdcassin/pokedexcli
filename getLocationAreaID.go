package main

import (
	"strings"
	"errors"
)

func getLocationAreaID(url string) (string, error) {
	id := ""
	parts := strings.Split(url, "/")

	if len(parts) < 2 {
		return "", errors.New("invalid URL format")
	}

	nonBlanks := []string{}
	for _, part := range parts {
		if part != "" {
			nonBlanks = append(nonBlanks, part)
		}
	}

	if len(nonBlanks) == 0 {
		return "", errors.New("invalid URL format, no non-blank parts found")
	}
	if len(nonBlanks) < 2 {
		return "", errors.New("invalid URL format, not enough parts found")
	}

	id = nonBlanks[len(nonBlanks) - 1]

	return id, nil
}