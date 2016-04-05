package securepassctl

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

// DebugLogger collects all debug messages
var DebugLogger = log.New(ioutil.Discard, "", log.LstdFlags)

// SecurePass main object type
type SecurePass struct {
	AppID     string `ini:"app_id"`
	AppSecret string `ini:"app_secret"`
	Endpoint  string `ini:"endpoint"`
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
func (s *SecurePass) NewRequest(method, path string, data *url.Values) (*http.Request, error) {
	var err error
	var req *http.Request

	URL, err := s.makeRequestURL(path)
	if err != nil {
		return nil, err
	}
	if data != nil {
		req, err = http.NewRequest(method, URL, bytes.NewBufferString(data.Encode()))
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
	DebugLogger.Printf("Sending %s request to %s", req.Method, req.URL)
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

	req, err := s.NewRequest("POST", "/api/v1/apps/info", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// AppAdd represents /api/v1/apps/add
func (s *SecurePass) AppAdd(app *ApplicationDescriptor) (*AppAddResponse, error) {
	var obj AppAddResponse

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

	req, err := s.NewRequest("POST", "/api/v1/apps/add", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// AppDel deletes an application from SecurePass
func (s *SecurePass) AppDel(app string) (*Response, error) {
	var obj Response

	data := url.Values{}
	if app != "" {
		data.Set("APP_ID", app)
	}

	req, err := s.NewRequest("POST", "/api/v1/apps/delete", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// AppMod represents /api/v1/apps/modify
func (s *SecurePass) AppMod(appID string, app *ApplicationDescriptor) (*Response, error) {
	var obj Response

	data := url.Values{}
	data.Set("APP_ID", appID)
	data.Set("WRITE", fmt.Sprintf("%v", app.Write))
	data.Set("PRIVACY", fmt.Sprintf("%v", app.Privacy))
	if app.Label != "" {
		data.Set("LABEL", app.Label)
	}
	if app.AllowNetworkIPv4 != "" {
		data.Set("ALLOW_NETWORK_IPv4", app.AllowNetworkIPv4)
	}
	if app.AllowNetworkIPv6 != "" {
		data.Set("ALLOW_NETWORK_IPv6", app.AllowNetworkIPv6)
	}
	if app.Group != "" {
		data.Set("GROUP", app.Group)
	}

	req, err := s.NewRequest("POST", "/api/v1/apps/modify", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// AppList retrieves the list of applications available in SecurePass
func (s *SecurePass) AppList(realm string) (*AppListResponse, error) {
	var obj AppListResponse

	data := url.Values{}
	if realm != "" {
		data.Set("REALM", realm)
	}

	req, err := s.NewRequest("POST", "/api/v1/apps/list", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// Logs retrieves application logs
func (s *SecurePass) Logs(realm, start, end string) (*LogsResponse, error) {
	var obj LogsResponse

	data := url.Values{}
	if realm != "" {
		data.Set("REALM", realm)
	}
	if start != "" {
		data.Set("START", start)
	}
	if end != "" {
		data.Set("END", end)
	}

	req, err := s.NewRequest("POST", "/api/v1/logs/get", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// GroupMember issues requests to /api/v1/groups/member
func (s *SecurePass) GroupMember(user, group string) (*GroupMemberResponse, error) {
	var obj GroupMemberResponse

	data := url.Values{}
	data.Set("USERNAME", user)
	data.Set("GROUP", group)

	req, err := s.NewRequest("POST", "/api/v1/groups/member", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// UserInfo issues requests to /api/v1/users/info
func (s *SecurePass) UserInfo(username string) (*UserInfoResponse, error) {
	var obj UserInfoResponse

	data := url.Values{}
	data.Set("USERNAME", username)

	req, err := s.NewRequest("POST", "/api/v1/users/info", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// UserList issues requests to /api/v1/users/list
func (s *SecurePass) UserList(realm string) (*UserListResponse, error) {
	var obj UserListResponse

	data := url.Values{}
	if realm != "" {
		data.Set("REALM", realm)
	}

	req, err := s.NewRequest("POST", "/api/v1/users/list", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// UserAuth issues requests to /api/v1/users/auth
func (s *SecurePass) UserAuth(username, secret string) (*UserAuthResponse, error) {
	var obj UserAuthResponse

	data := url.Values{}
	data.Set("USERNAME", username)
	data.Set("SECRET", secret)

	req, err := s.NewRequest("POST", "/api/v1/users/auth", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// UserAdd issues requests to /api/v1/users/add
func (s *SecurePass) UserAdd(user *UserDescriptor) (*UserAddResponse, error) {
	var obj UserAddResponse

	data := url.Values{}
	data.Set("USERNAME", user.Username)
	data.Set("NAME", user.Name)
	data.Set("SURNAME", user.Surname)
	data.Set("EMAIL", user.Email)
	data.Set("MOBILE", user.Mobile)
	data.Set("NIN", user.Nin)
	data.Set("RFID", user.Rfid)
	data.Set("MANAGER", user.Manager)

	req, err := s.NewRequest("POST", "/api/v1/users/add", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// UserDel deletes a user from SecurePass
func (s *SecurePass) UserDel(username string) (*Response, error) {
	var obj Response

	data := url.Values{}
	data.Set("USERNAME", username)

	req, err := s.NewRequest("POST", "/api/v1/users/delete", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// UserPasswordChange change user password
func (s *SecurePass) UserPasswordChange(username, password string) (*Response, error) {
	var obj Response

	data := url.Values{}
	if username != "" && password != "" {
		data.Set("USERNAME", username)
		data.Set("PASSWORD", password)
	}

	req, err := s.NewRequest("POST", "/api/v1/users/password/change", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// UserPasswordDisable disable a user's password
func (s *SecurePass) UserPasswordDisable(username string) (*Response, error) {
	var obj Response

	data := url.Values{}
	data.Set("USERNAME", username)

	req, err := s.NewRequest("POST", "/api/v1/users/password/disable", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// UserProvision provisions a user with a token
func (s *SecurePass) UserProvision(username, token string) (*Response, error) {
	var obj Response

	data := url.Values{}
	data.Set("USERNAME", username)
	data.Set("SWTOKEN", token)

	req, err := s.NewRequest("POST", "/api/v1/users/token/provision", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// UserXattrsDelete deletes an attribute from user's extended attributes
func (s *SecurePass) UserXattrsDelete(username, attribute string) (*Response, error) {
	var obj Response

	data := url.Values{}
	data.Set("USERNAME", username)
	data.Set("ATTRIBUTE", attribute)

	req, err := s.NewRequest("POST", "/api/v1/users/xattrs/delete", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// UserXattrsList lists user's extended attributes
func (s *SecurePass) UserXattrsList(username string) (*XattrsListResponse, error) {
	var obj XattrsListResponse

	data := url.Values{}
	data.Set("USERNAME", username)

	req, err := s.NewRequest("POST", "/api/v1/users/xattrs/list", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// UserXattrsSet set user's extended attributes
func (s *SecurePass) UserXattrsSet(username, attribute, value string) (*Response, error) {
	var obj Response

	data := url.Values{}
	data.Set("USERNAME", username)
	data.Set("ATTRIBUTE", attribute)
	data.Set("VALUE", value)

	req, err := s.NewRequest("POST", "/api/v1/users/xattrs/set", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// RadiusAdd adds a RADIUS to SecurePass RADIUS
func (s *SecurePass) RadiusAdd(radius *RadiusDescriptor) (*Response, error) {
	var obj Response

	data := url.Values{}
	data.Set("RADIUS", radius.Radius)
	data.Set("NAME", radius.Name)
	data.Set("SECRET", radius.Secret)
	data.Set("GROUP", radius.Group)
	data.Set("RFID", fmt.Sprintf("%v", radius.Rfid))
	data.Set("REALM", radius.Realm)

	req, err := s.NewRequest("POST", "/api/v1/radius/add", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// RadiusInfo retrieves information on a SecurePass RADIUS device
func (s *SecurePass) RadiusInfo(ipaddr string) (*RadiusInfoResponse, error) {
	var obj RadiusInfoResponse

	data := url.Values{}
	data.Set("RADIUS", ipaddr)

	req, err := s.NewRequest("POST", "/api/v1/radius/info", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// RadiusDel deletes a RADIUS device from SecurePass
func (s *SecurePass) RadiusDel(ipaddr string) (*Response, error) {
	var obj Response

	data := url.Values{}
	data.Set("RADIUS", ipaddr)

	req, err := s.NewRequest("POST", "/api/v1/radius/delete", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// RadiusList retrieves the list of RADIUS devices available in SecurePass
func (s *SecurePass) RadiusList(realm string) (*RadiusListResponse, error) {
	var obj RadiusListResponse

	data := url.Values{}
	if realm != "" {
		data.Set("REALM", realm)
	}

	req, err := s.NewRequest("POST", "/api/v1/radius/list", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// RadiusMod modify a RADIUS device available in SecurePass
func (s *SecurePass) RadiusMod(radiusID string, radius *RadiusDescriptor) (*Response, error) {
	var obj Response

	data := url.Values{}
	data.Set("RADIUS", radiusID)
	data.Set("RFID", fmt.Sprintf("%v", radius.Rfid))
	if radius.Name != "" {
		data.Set("NAME", radius.Name)
	}
	if radius.Secret != "" {
		data.Set("SECRET", radius.Secret)
	}
	if radius.Realm != "" {
		data.Set("REALM", radius.Realm)
	}
	if radius.Group != "" {
		data.Set("GROUP", radius.Group)
	}

	req, err := s.NewRequest("POST", "/api/v1/radius/modify", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// RealmXattrsDelete deletes an attribute from realm's extended attributes
func (s *SecurePass) RealmXattrsDelete(realm, attribute string) (*Response, error) {
	var obj Response

	data := url.Values{}
	data.Set("REALM", realm)
	data.Set("ATTRIBUTE", attribute)

	req, err := s.NewRequest("POST", "/api/v1/realms/xattrs/delete", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// RealmXattrsList lists realm's extended attributes
func (s *SecurePass) RealmXattrsList(realm string) (*XattrsListResponse, error) {
	var obj XattrsListResponse

	data := url.Values{}
	data.Set("REALM", realm)

	req, err := s.NewRequest("POST", "/api/v1/realms/xattrs/list", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// RealmXattrsSet set realm's extended attributes
func (s *SecurePass) RealmXattrsSet(realm, attribute, value string) (*Response, error) {
	var obj Response

	data := url.Values{}
	data.Set("REALM", realm)
	data.Set("ATTRIBUTE", attribute)
	data.Set("VALUE", value)

	req, err := s.NewRequest("POST", "/api/v1/realms/xattrs/set", &data)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}

// Ping issues requests to the /api/v1/ping API endpoint
func (s *SecurePass) Ping() (*PingResponse, error) {
	var obj PingResponse

	req, err := s.NewRequest("GET", "/api/v1/ping", nil)
	if err != nil {
		return nil, err
	}
	err = s.DoRequest(req, &obj, 200)
	return &obj, err
}
