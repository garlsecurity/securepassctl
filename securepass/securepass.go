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
	req.Header.Set("User-Agent", UserAgent)
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

// NewRequest initializes and issues an HTTP request to the SecurePass endpoint
func (s *SecurePass) NewRequest(method, path string, buf *bytes.Buffer) (*http.Request, error) {
	var err error
	var req *http.Request

	URL, err := s.makeRequestURL(path)
	if err != nil {
		return nil, err
	}
	if buf != nil {
		req, err = http.NewRequest(method, URL, buf)
	} else {
		req, err = http.NewRequest(method, URL, nil)
	}
	if err != nil {
		return nil, err
	}
	s.setupRequestFieds(req)
	return req, nil
}

// DoRequest issues an HTTP request
func (s *SecurePass) DoRequest(req *http.Request, obj APIResponse, expstatus int) error {
	client := NewClient(nil)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != expstatus {
		return fmt.Errorf("%s", resp.Status)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(obj)
	if err != nil {
		return err
	}
	if obj.ErrorCode() != 0 {
		return fmt.Errorf("%d: %s", obj.ErrorCode(), obj.ErrorMessage())
	}
	return nil
}

// NewClient initialize http.Client with a certain http.Transport
func NewClient(tr *http.Transport) *http.Client {
	// Skip SSL certificate verification
	if tr == nil {
		tr = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	}
	return &http.Client{Transport: tr}
}

// AppInfo retrieves information on a SecurePass application
func (s *SecurePass) AppInfo(app string) (*AppInfoResponse, error) {
	var obj AppInfoResponse

	data := url.Values{}
	if app != "" {
		data.Set("APP_ID", app)
	}

	req, err := s.NewRequest("POST", "/api/v1/apps/info", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// AppInfo represents /api/v1/apps/info
func (s *SecurePass) AppAdd(app *ApplicationDescriptor) (*AppdAddResponse, error) {
	var obj AppdAddResponse

	data := url.Values{}
	data.Set("LABEL", app.Label)
	data.Set("WRITE", fmt.Sprintf("%v", app.Write))
	data.Set("PRIVACY", fmt.Sprintf("%v", app.Privacy))
	if app.AllowNetworkIPv4 != "" {
		data.Set("ALLOW_NETWORK_IPv4", app.AllowNetworkIPv4)
	}
	if app.AllowNetworkIPv6 != "" {
		data.Set("ALLOW_NETWORK_IPv6", app.AllowNetworkIPv6)
	}
	if app.Group != "" {
		data.Set("GROUP", app.Group)
	}
	if app.Realm != "" {
		data.Set("REALM", app.Realm)
	}

	req, err := s.NewRequest("POST", "/api/v1/apps/add", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// Ping reprenets the /api/v1/ping API call
func (s *SecurePass) Ping() (*PingResponse, error) {
	var obj PingResponse

	req, err := s.NewRequest("GET", "/api/v1/ping", nil)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}
