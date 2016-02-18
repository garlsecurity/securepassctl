package securepass

import (
	"fmt"
)

func ExampleNewSecurePass() {
	appid, appsecret, remote := "appid", "appsecret", "https://test-remote.securepass.net"
	s, err := NewSecurePass(appid, appsecret, remote)
	fmt.Println(err)
	fmt.Println(s.AppID)
	fmt.Println(s.AppSecret)
	fmt.Println(s.Endpoint)
	// Output:
	// <nil>
	// appid
	// appsecret
	// https://test-remote.securepass.net
}
