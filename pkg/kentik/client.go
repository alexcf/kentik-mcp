package kentik

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Config holds the credentials and region for authenticating with Kentik.
type Config struct {
	Email    string
	APIToken string
	Region   string // "US" (default) or "EU"
}

// Profile is a named set of Kentik credentials.
type Profile struct {
	Name     string
	Email    string
	APIToken string
	Region   string
}

// Client is an HTTP client for the Kentik API.
// It supports runtime account switching via SwitchProfile.
type Client struct {
	mu       sync.RWMutex
	email    string
	apiToken string
	v5Base   string
	v6Base   string
	current  string // profile name, or "default"
	profiles map[string]Profile
	http     *http.Client
}

func baseURLs(region string) (v5Base, v6Base string) {
	if strings.ToUpper(region) == "EU" {
		return "https://api.kentik.eu/api/v5", "https://grpc.api.kentik.eu"
	}
	return "https://api.kentik.com/api/v5", "https://grpc.api.kentik.com"
}

// NewClient creates a new Kentik API client.
func NewClient(cfg Config) *Client {
	v5, v6 := baseURLs(cfg.Region)
	return &Client{
		email:    cfg.Email,
		apiToken: cfg.APIToken,
		v5Base:   v5,
		v6Base:   v6,
		current:  "default",
		profiles: map[string]Profile{},
		http: &http.Client{
			Timeout: 120 * time.Second,
		},
	}
}

// AddProfiles registers named profiles for runtime switching.
func (c *Client) AddProfiles(profiles []Profile) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, p := range profiles {
		c.profiles[strings.ToLower(p.Name)] = p
	}
}

// SwitchProfile switches to a named profile. Returns an error if the profile is not found.
func (c *Client) SwitchProfile(name string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	key := strings.ToLower(name)
	p, ok := c.profiles[key]
	if !ok {
		return fmt.Errorf("profile %q not found; use kentik_list_accounts to see available profiles", name)
	}
	c.email = p.Email
	c.apiToken = p.APIToken
	c.v5Base, c.v6Base = baseURLs(p.Region)
	c.current = p.Name
	return nil
}

// ListProfiles returns all registered profile names and their email/region (no tokens).
func (c *Client) ListProfiles() []Profile {
	c.mu.RLock()
	defer c.mu.RUnlock()
	out := make([]Profile, 0, len(c.profiles))
	for _, p := range c.profiles {
		out = append(out, Profile{Name: p.Name, Email: p.Email, Region: p.Region})
	}
	return out
}

// CurrentProfile returns the name of the active profile.
func (c *Client) CurrentProfile() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.current
}

// CurrentEmail returns the email of the active credentials (for display).
func (c *Client) CurrentEmail() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.email
}

func (c *Client) headers() map[string]string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return map[string]string{
		"X-CH-Auth-Email":     c.email,
		"X-CH-Auth-API-Token": c.apiToken,
		"Content-Type":        "application/json",
	}
}

func (c *Client) v5URL(path string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.v5Base + path
}

func (c *Client) v6URL(path string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.v6Base + path
}

func (c *Client) doRequest(method, url string, body interface{}) (json.RawMessage, error) {
	var reqBody io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal request body: %w", err)
		}
		reqBody = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	for k, v := range c.headers() {
		req.Header.Set(k, v)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(respBody))
	}

	return json.RawMessage(respBody), nil
}

// V5 makes a request to the Kentik V5 REST API.
// path should start with "/" e.g. "/devices".
func (c *Client) V5(method, path string, body interface{}) (json.RawMessage, error) {
	return c.doRequest(method, c.v5URL(path), body)
}

// V6 makes a request to the Kentik V6 gRPC-gateway API.
// path should be the full path e.g. "/synthetics/v202309/tests".
func (c *Client) V6(method, path string, body interface{}) (json.RawMessage, error) {
	return c.doRequest(method, c.v6URL(path), body)
}
