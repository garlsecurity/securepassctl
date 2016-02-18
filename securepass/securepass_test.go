package securepass

import (
	"fmt"
)

func ExampleNewSecurepass() {
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

func ExamplePing() {
	appid := "fe2564e8d65d11e5972fde2f4665c1f2@login.farm"
	appsecret := "dGFmR4RoWMnRxf1137fkXcVgvRnd1G1BAFzZS5sJJKUpKXQBFy"
	s, err := NewSecurePass(appid, appsecret, DefaultRemote)
	fmt.Println(err)
	resp, err := s.Ping()
	fmt.Println(err)
	fmt.Println(resp.IPVersion)
	fmt.Println(resp.RC)
	fmt.Println(resp.ErrorMsg)
	// Output:
	// <nil>
	// <nil>
	// 4
	// 0
	//
}
