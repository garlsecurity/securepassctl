package securepassctl

// APIResponse gives access to the response details
type APIResponse interface {
	ErrorCode() int
	ErrorMessage() string
}

// Response is the base type for API calls responses
type Response struct {
	APIResponse
	ErrorMsg string
	RC       int
}

// ErrorCode returns the API call's numeric return code
func (r *Response) ErrorCode() int { return r.RC }

// ErrorMessage returns the API call's text message
func (r *Response) ErrorMessage() string { return r.ErrorMsg }

// PingResponse represents the /api/v1/ping call's HTTP response
type PingResponse struct {
	IP        string
	IPVersion int `json:"ip_version"`
	Response
}

// AppAddResponse describes the expected response from the
// /api/v1/apps/add
type AppAddResponse struct {
	AppID     string `json:"APP_ID"`
	AppSecret string `json:"APP_SECRET"`
	Response
}

// AppInfoResponse encapsulates the /api/v1/apps/info call's HTTP response
type AppInfoResponse struct {
	ApplicationDescriptor
	Response
}

// AppListResponse encapsulates the /api/v1/apps HTTP response
type AppListResponse struct {
	AppID []string `json:"APP_ID"`
	Response
}

// LogsResponse encapsulates SecurePass application's logs
type LogsResponse struct {
	Logs map[string]LogEntry
	Response
}

// GroupMemberResponse encapsulates whether a group belogs to a member
type GroupMemberResponse struct {
	Member bool
	Response
}

// UserInfoResponse encapsulates the /api/v1/users/info HTTP response
type UserInfoResponse struct {
	UserDescriptor
	Response
}

// UserListResponse encapsulates the /api/v1/users HTTP response
type UserListResponse struct {
	Username []string
	Response
}

// UserAddResponse encapsulates the /api/v1/users/add HTTP response
type UserAddResponse struct {
	Username string
	Response
}

// UserAuthResponse encapsulates the /api/v1/users/auth HTTP response
type UserAuthResponse struct {
	Authenticated bool
	Response
}

// UserXattrsListResponse encapsulates the /api/v1/users/xattrs/list HTTP response
type UserXattrsListResponse UserXattrsDescriptor

// ErrorCode returns the API call's numeric return code
func (r *UserXattrsListResponse) ErrorCode() int {
	rc := map[string]interface{}(UserXattrsDescriptor(*r))["rc"]
	return int(rc.(float64))
}

// ErrorMessage returns the API call's text message
func (r *UserXattrsListResponse) ErrorMessage() string {
	v := map[string]interface{}(UserXattrsDescriptor(*r))["errorMsg"]
	u := v.(string)
	return u
}
