package integration

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type ActiveCampaignTags struct {
	Tags []struct {
		ID  string `json:"id"`
		Tag string `json:"tag"`
	} `json:"tags"`
}

type ActiveCampaignIntegration struct {
	APIKey   string
	APIURL   string
	TagsPath string
}

func (integration ActiveCampaignIntegration) URL() string {
	return integration.APIURL
}

func (integration ActiveCampaignIntegration) Authenticate(request *http.Request) {
	request.Header.Add("Api-Token", integration.APIKey)
}

func (integration ActiveCampaignIntegration) GetTagsPath() string {
	return integration.TagsPath
}

func (integration ActiveCampaignIntegration) UnmarshalTags(data []byte) ([]Tag, error) {
	var activeCampaignTags ActiveCampaignTags
	err := json.Unmarshal(data, &activeCampaignTags)
	if err != nil {
		return []Tag{}, err
	}
	var tags []Tag
	for _, tag := range activeCampaignTags.Tags {
		tags = append(tags, Tag{ID: tag.ID, Name: tag.Tag})
	}
	return tags, nil
}

func NewActiveCampaignIntegration(id int, conn *sqlx.DB) (*ActiveCampaignIntegration, error) {
	var apiKey string
	var apiURL string

	err := getFieldValue(conn, id, "api_key", &apiKey)
	if err != nil {
		return nil, err
	}

	err = getFieldValue(conn, id, "api_url", &apiURL)
	if err != nil {
		return nil, err
	}

	return &ActiveCampaignIntegration{
		APIURL:   apiURL,
		APIKey:   apiKey,
		TagsPath: "/api/3/tags",
	}, nil
}
