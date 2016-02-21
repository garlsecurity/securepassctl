package securepass

import (
	"fmt"
)

func ExampleSecurePass() {
	appid, appsecret, remote := "appid", "appsecret", "https://test-remote.securepass.net"
	s := SecurePass{AppID: appid, AppSecret: appsecret, Endpoint: remote}
	fmt.Println(s.AppID)
	fmt.Println(s.AppSecret)
	fmt.Println(s.Endpoint)
	// Output:
	// appid
	// appsecret
	// https://test-remote.securepass.net
}

func ExampleSecurePass_Ping() {
	s := SecurePass{
		AppID:     "ce64dc90d88b11e5b001de2f4665c1f2@ci.secure-pass.net",
		AppSecret: "E2m6HawI743as61Kv0OhyPb6wAewXnwVkLLcF82rKOWe1SJ0Wd",
		Endpoint:  DefaultRemote,
	}
	resp, err := s.Ping()
	fmt.Println(err)
	fmt.Println(resp.IPVersion)
	fmt.Println(resp.RC)
	fmt.Println(resp.ErrorMsg)
	// Output:
	// <nil>
	// 4
	// 0
	//
}
