package integration

import (
	"os"
)

type ConvertKitIntegration struct {
	APIKey    string
	APISecret string
	APIURL    string
	TagsPath  string
}

func NewConvertKitIntegration() *ConvertKitIntegration {
	return &ConvertKitIntegration{
		APIKey:    os.Getenv("CONVERTKIT_API_KEY"),
		APISecret: os.Getenv("CONVERTKIT_API_SECRET"),
		APIURL:    "https://api.convertkit.com",
		TagsPath:  "/v3/tags",
	}
}
