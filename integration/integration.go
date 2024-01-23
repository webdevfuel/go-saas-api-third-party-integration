package integration

import "net/http"

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type App interface {
	// HTTP
	Authenticate(request *http.Request)
	URL() string
	// Tags
	GetTagsPath() string
	UnmarshalTags(body []byte) ([]Tag, error)
}
