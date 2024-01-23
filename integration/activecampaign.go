package integration

import "os"

type ActiveCampaignIntegration struct {
	APIKey   string
	APIURL   string
	TagsPath string
}

func NewActiveCampaignIntegration() *ActiveCampaignIntegration {
	return &ActiveCampaignIntegration{
		APIURL:   os.Getenv("ACTIVECAMPAIGN_API_URL"),
		APIKey:   os.Getenv("ACTIVECAMPAIGN_API_KEY"),
		TagsPath: "/api/3/tags",
	}
}
