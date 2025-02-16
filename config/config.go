package config

import (
	"encoding/json"
	"flag"
	"fmt"
	ratelimiter "github.com/lucaiatropulus/social/internal/rate_limiter"
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	APP         appConfig          `json:"app"`
	DB          dbConfig           `json:"db"`
	Mail        mailConfig         `json:"mail"`
	Auth        authConfig         `json:"auth"`
	Redis       redisConfig        `json:"redis"`
	RateLimiter ratelimiter.Config `json:"rateLimiter"`
}

func LoadConfig(environment *string) *Config {
	flag.Parse()

	fileName := fmt.Sprintf("env/env.%s.yaml", *environment)
	data, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal("Unable to read config file: ", err)
	}

	jsonData, err := convertYAMLToJSON(string(data))

	if err != nil {
		log.Fatal("Unable to parse config file: ", err)
	}

	var config Config
	err = json.Unmarshal([]byte(jsonData), &config)

	if err != nil {
		log.Fatal("Unable to read config file: ", err)
	}

	return &config
}

func convertYAMLToJSON(yamlData string) (string, error) {
	lines := strings.Split(yamlData, "\n")
	jsonMap := make(map[string]interface{})
	var currentSection map[string]interface{}
	var lastKey string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Skip empty lines and comments
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		// Detect key-value pairs
		if strings.Contains(trimmed, ":") {
			parts := strings.SplitN(trimmed, ":", 2)
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			// Check for nested objects (if value is empty)
			if value == "" {
				// Create a new section
				newSection := make(map[string]interface{})
				jsonMap[key] = newSection
				currentSection = newSection
			} else {
				// Try to parse as integer
				if intValue, err := strconv.Atoi(value); err == nil {
					currentSection[key] = intValue
				} else if value == "true" || value == "false" {
					// Handle boolean values
					currentSection[key] = (value == "true")
				} else {
					// Treat as string
					currentSection[key] = strings.Trim(value, `"`)
				}
			}
			lastKey = key
		} else {
			// If a line doesn't have ":", it could be an array item
			if lastKey != "" {
				if _, exists := currentSection[lastKey]; !exists {
					currentSection[lastKey] = []string{}
				}
				currentSection[lastKey] = append(currentSection[lastKey].([]string), trimmed)
			}
		}
	}

	// Convert map to JSON
	jsonData, err := json.Marshal(jsonMap)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}

	return string(jsonData), nil
}
