package config

import "os"

var ServerHostPort = getParam("SERVER_HOST_PORT", "localhost:8090")

func getParam(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists == false {
		return defaultValue
	}
	return value
}
