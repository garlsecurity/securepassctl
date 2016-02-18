package securepass

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	// DefaultRemote is the default Content-Type header used in HTTP requests
	DefaultRemote = "https://beta.secure-pass.net"
	// ContentType is the default Content-Type header used in HTTP requests
	ContentType = "application/json"
	// UserAgent contains the default User-Agent value used in HTTP requests
	UserAgent = "SecurePass CLI"
)

// SecurePass main object type
type SecurePass struct {
	AppID     string `ini:"APP_ID"`
	AppSecret string `ini:"APP_SECRET"`
	Endpoint  string
}

// PingResponse represents the /api/v1/ping call's HTTP response
type PingResponse struct {
	ErrorMsg  string
	IP        string
	IPVersion int `json:"ip_version"`
	RC        int
}

// NewSecurePass makes and initialize a new SecurePass instance
func NewSecurePass(appid, appsecret, remote string) (*SecurePass, error) {
	u, err := url.Parse(remote)
	if err != nil {
		return nil, err
	}
	if !u.IsAbs() {
		return nil, fmt.Errorf("'%s' is not an absolute URL", remote)
	}
	if u.Scheme != "https" {
		return nil, fmt.Errorf("scheme of '%s' isn't 'https'", remote)
	}

	return &SecurePass{
		AppID:     appid,
		AppSecret: appsecret,
		Endpoint:  remote}, nil
}

func (s *SecurePass) setupRequestFieds(req *http.Request) {
	req.Header.Set("Accept", ContentType)
	req.Header.Set("Content-Type", ContentType)
	req.Header.Set("X-SecurePass-App-ID", s.AppID)
	req.Header.Set("X-SecurePass-App-Secret", s.AppSecret)
}

func (s *SecurePass) makeRequestURL(path string) (string, error) {
	baseURL, _ := url.Parse(s.Endpoint)
	URL, err := url.Parse(path)
	if err != nil {
		return "", err
	}
	return baseURL.ResolveReference(URL).String(), nil
}

func (s *SecurePass) makeRequest(method,
	path string, buf *bytes.Buffer) (*http.Request, error) {
	URL, err := s.makeRequestURL(path)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, URL, buf)
	if err != nil {
		return nil, err
	}
	s.setupRequestFieds(req)
	return req, nil
}

// NewClient initialize http.Client with a certain http.Transport
func NewClient(tr *http.Transport) *http.Client {
	// Skip SSL certificate verification
	if tr == nil {
		tr = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	}
	return &http.Client{Transport: tr}
}

// NewRequest initializes and issues an HTTP request to the SecurePass endpoint
func (s *SecurePass) NewRequest(method, path string,
	content *map[string]interface{}) (*http.Request, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(content); err != nil {
		return nil, err
	}
	return s.makeRequest(method, path, &buf)
}

// Ping reprenets the /api/v1/ping API call
func (s *SecurePass) Ping() (*PingResponse, error) {
	var obj PingResponse
	client := NewClient(nil)

	req, err := s.NewRequest("GET", "/api/v1/ping", nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", resp.Status)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&obj)
	if err != nil {
		return nil, err
	}
	if obj.RC != 0 {
		return &obj, fmt.Errorf("%v", obj.ErrorMsg)
	}
	return &obj, nil
}
