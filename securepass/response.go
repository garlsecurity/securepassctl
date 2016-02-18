package securepass

// Response is the base type for API calls responses
type Response struct {
	ErrorMsg string
	RC       int
}

// PingResponse represents the /api/v1/ping call's HTTP response
type PingResponse struct {
	IP        string
	IPVersion int `json:"ip_version"`
	Response
}

// AppInfoResponse encapsulates the /api/v1/apps/info call's HTTP response
type AppInfoResponse struct {
	Label            string
	Realm            string
	Group            string
	Write            bool
	AllowNetworkIPv4 string `json:"allow_network_ipv4"`
	AllowNetworkIPv6 string `json:"allow_network_ipv6"`
	Privacy          bool
	Response
}
