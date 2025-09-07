package envy

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Get returns the value of an environment variable or a default value if not found.
func Get(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

// Require returns the value of an environment variable or panics if not found.
func Require(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("Required environment variable %s not set", key))
	}
	return val
}

// GetInt returns an integer environment variable or default if not found.
func GetInt(key string, defaultVal int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	num, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return num
}

// GetBool returns a boolean environment variable or default if not found.
func GetBool(key string, defaultVal bool) bool {
	val := strings.ToLower(os.Getenv(key))
	if val == "" {
		return defaultVal
	}
	if val == "true" || val == "1" || val == "yes" {
		return true
	}
	if val == "false" || val == "0" || val == "no" {
		return false
	}
	return defaultVal
}

// Load loads environment variables from a .env file.
func Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// skip comments and empty lines
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}
	return scanner.Err()
}
