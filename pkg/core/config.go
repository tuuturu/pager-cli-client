package core

import (
	"fmt"
	"net/url"
	"os"
)

type Config struct {
	DiscoveryURL *url.URL
	ClientID     string
	ClientSecret string

	EventsServiceURL *url.URL
}

func (c Config) Validate() error {
	return nil
}

func LoadConfig() (Config, error) {
	cfg := Config{}

	if rawDiscoveryURL := os.Getenv("DISCOVERY_URL"); rawDiscoveryURL != "" {
		discoveryURL, err := url.Parse(rawDiscoveryURL)
		if err != nil {
			return cfg, fmt.Errorf("error parsing DISCOVERY_URL: %w", err)
		}

		cfg.DiscoveryURL = discoveryURL
	}

	cfg.ClientID = os.Getenv("CLIENT_ID")
	cfg.ClientSecret = os.Getenv("CLIENT_SECRET")

	if rawEventsServiceURL := os.Getenv("EVENTS_SERVICE_URL"); rawEventsServiceURL != "" {
		eventsServiceURL, err := url.Parse(rawEventsServiceURL)
		if err != nil {
			return cfg, fmt.Errorf("error parsing EVENTS_SERVICE_URL: %w", err)
		}

		cfg.EventsServiceURL = eventsServiceURL
	}

	return cfg, nil
}
