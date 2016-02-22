package securepass

import (
	"fmt"
)

var SecurePassInst SecurePass

func init() {
	SecurePassInst = SecurePass{
		AppID:     "ce64dc90d88b11e5b001de2f4665c1f2@ci.secure-pass.net",
		AppSecret: "E2m6HawI743as61Kv0OhyPb6wAewXnwVkLLcF82rKOWe1SJ0Wd",
		Endpoint:  DefaultRemote,
	}
}

func ExampleSecurePass() {
	fmt.Println(SecurePassInst.AppID)
	fmt.Println(SecurePassInst.AppSecret)
	fmt.Println(SecurePassInst.Endpoint)
	// Output:
	// ce64dc90d88b11e5b001de2f4665c1f2@ci.secure-pass.net
	// E2m6HawI743as61Kv0OhyPb6wAewXnwVkLLcF82rKOWe1SJ0Wd
	// https://beta.secure-pass.net
}

func ExampleSecurePass_Ping() {
	resp, err := SecurePassInst.Ping()
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
