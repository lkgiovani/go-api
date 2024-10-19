package env

import (
	"fmt"
	"os"
	"strconv"
)

func GetEnvOrDie(key string) string {
	value := os.Getenv(key)

	if value == "" {
		err := fmt.Errorf("Missing environment variable %s", key)
		panic(err)
	}

	return value
}

func GetEnvOrDieAsInt(key string) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		fmt.Printf("Environment variable %s not set\n", key)
		os.Exit(1)
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Printf("Error converting %s to int: %v\n", key, err)
		os.Exit(1)
	}

	return value
}
