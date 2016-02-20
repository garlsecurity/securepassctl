package securepass

// APIResponse gives access to the response details
type APIResponse interface {
	ErrorCode() int
	ErrorMessage() string
}

// ApplicationDescriptor describes the basic attributes of
// a Securepass application
type ApplicationDescriptor struct {
	Label            string `json:"label"`
	Realm            string `json:"realm"`
	Group            string `json:"group"`
	Write            bool   `json:"write"`
	AllowNetworkIPv4 string `json:"allow_network_ipv4"`
	AllowNetworkIPv6 string `json:"allow_network_ipv6"`
	Privacy          bool   `json:"privacy"`
}

// Response is the base type for API calls responses
type Response struct {
	APIResponse
	ErrorMsg string
	RC       int
}

// ErrorCode returns the API call's numeric return code
func (r *Response) ErrorCode() int {
	return r.RC
}

// ErrorMessage returns the API call's text message
func (r *Response) ErrorMessage() string {
	return r.ErrorMsg
}

// PingResponse represents the /api/v1/ping call's HTTP response
type PingResponse struct {
	IP        string
	IPVersion int `json:"ip_version"`
	Response
}

// AppAddResponse describes the expected response from the
// /api/v1/apps/add
type AppAddResponse struct {
	AppID     string `ini:"APP_ID"`
	AppSecret string `ini:"APP_SECRET"`
	Response
}

// AppInfoResponse encapsulates the /api/v1/apps/info call's HTTP response
type AppInfoResponse struct {
	ApplicationDescriptor
	Response
}
