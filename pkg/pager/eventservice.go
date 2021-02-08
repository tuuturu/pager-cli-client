package pager

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tuuturu/pager-event-service/pkg/core/models"
	"net/http"
	"net/url"
)

func CreateEvent(baseURL *url.URL, token string, event models.Event) error {
	eventsURL := fmt.Sprintf("%s/events", baseURL.String())

	payload, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("error marshalling payload: %w", err)
	}

	request, err := http.NewRequest(http.MethodPost, eventsURL, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("error creating event request: %w", err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("error posting event: %w", err)
	}

	if response.StatusCode != http.StatusCreated {
		return errors.New(fmt.Sprintf("response returned %d", response.StatusCode))
	}

	return nil
}
