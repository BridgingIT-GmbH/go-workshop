package config

import "os"

var MongoUri = getParam("MONGO_URI", "mongodb://localhost:27017")
var ServerHostPort = getParam("SERVER_HOST_PORT", "localhost:8080")

func getParam(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists == false {
		return defaultValue
	}
	return value
}
