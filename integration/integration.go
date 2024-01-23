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

func NewRequest(method, url string) (*http.Request, error) {
	resp, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
