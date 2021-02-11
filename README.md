# CLI client for the Pager Project

##  Description

A simple CLI client which can push notifications

## Usage

### Locally

```bash
export DISCOVERY_URL=https://authprovider/.well-known # Your authentication provider's [OIDC discovery URL](https://auth0.com/docs/protocols/configure-applications-with-oidc-discovery)
export CLIENT_ID=cliclient # The client ID representing this client in your authentication provider
export CLIENT_SECRET=supersecret # The client secret of the aforementioned client
export EVENTS_SERVICE_URL=https://events.mydomain.io # URL to your instance of the [events service](https://github.com/tuuturu/pager-event-service)

pager-cli-client TITLE DESCRIPTION

# Example
pager-cli-client "New notification" "This notification enlightens"
```

### Docker

```bash
docker run \
    -e DISCOVERY_URL=https://authprovider/.well-known \
    -e CLIENT_ID=cliclient \
    -e CLIENT_SECRET=supersecret \
    -e EVENTS_SERVICE_URL=https://events.mydomain.io \
    docker.pkg.github.com/tuuturu/pager-email-client/pager-email-client \
    "Title" "Short Description"
```
