package env

import "os"

// LookupWithDefault returns the environment variable for the key if found, otherwise returns the defaultVal passed.
func LookupWithDefault(key, defaultVal string) string {
	val, found := os.LookupEnv(key)
	if !found {
		return defaultVal
	}

	return val
}
