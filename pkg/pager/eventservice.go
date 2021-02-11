package pager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/tuuturu/pager-event-service/pkg/core/models"
)

func CreateEvent(baseURL *url.URL, token string, event models.Event) error {
	eventsURL := fmt.Sprintf("%s/events", baseURL.String())

	payload, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("marshalling payload: %w", err)
	}

	request, err := http.NewRequest(http.MethodPost, eventsURL, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("creating event request: %w", err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("posting event: %w", err)
	}

	if response.StatusCode != http.StatusCreated {
		return fmt.Errorf("response returned %d", response.StatusCode)
	}

	return nil
}
