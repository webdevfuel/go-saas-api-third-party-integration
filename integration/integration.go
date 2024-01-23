package integration

import (
	"fmt"
	"io"
	"net/http"

	"github.com/jmoiron/sqlx"
)

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

func GetIntegrationTags(app App) ([]Tag, error) {
	req, err := NewRequest("GET", app.URL())
	if err != nil {
		return []Tag{}, err
	}
	app.Authenticate(req)
	req.URL.Path = fmt.Sprintf("%s/%s", req.URL.Path, app.GetTagsPath())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []Tag{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Tag{}, err
	}
	tags, err := app.UnmarshalTags(body)
	if err != nil {
		return []Tag{}, err
	}
	return tags, nil
}

func NewRequest(method, url string) (*http.Request, error) {
	resp, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetIntegrationApp(id int, conn *sqlx.DB) (string, error) {
	var app string
	err := conn.QueryRow(`
		SELECT
		    apps.slug
		FROM
		    integrations
		    LEFT JOIN apps ON integrations.app_id = apps.id
		WHERE
		    integrations.id = $1`, id).Scan(&app)
	if err != nil {
		return "", err
	}
	return app, nil
}
