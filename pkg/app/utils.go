package app

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getEnv(varName, defaultValue string) string {
	env := os.Getenv(varName)
	if env == "" {
		return defaultValue
	}
	return varName
}

func getEnvAsInt(varName string, defaultValue int) int {
	env := getEnv(varName, "")
	if env == "" {
		return defaultValue
	}

	num, err := strconv.ParseInt(env, 10, 64)
	if err != nil {
		return defaultValue
	}

	return int(num)
}

func createIdentifierFromTitle(title string) string {
	identifier := strings.ReplaceAll(title, " ", "-")
	identifier = strings.ToLower(identifier)

	rawIdentifier := identifier
	counter := 1
	for {
		_, ok := feedStorage[identifier]
		if !ok {
			break
		}

		identifier = fmt.Sprintf("%s-%d", rawIdentifier, counter)
		counter++
	}

	return identifier
}
