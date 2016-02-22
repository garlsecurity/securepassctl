package securepass

import (
	"fmt"
	"math/rand"
	"time"
)

var testInstance SecurePass

func init() {
	testInstance = SecurePass{
		AppID:     "ce64dc90d88b11e5b001de2f4665c1f2@ci.secure-pass.net",
		AppSecret: "E2m6HawI743as61Kv0OhyPb6wAewXnwVkLLcF82rKOWe1SJ0Wd",
		Endpoint:  DefaultRemote,
	}
}

func ExampleSecurePass() {
	fmt.Println(testInstance.AppID)
	fmt.Println(testInstance.AppSecret)
	fmt.Println(testInstance.Endpoint)
	// Output:
	// ce64dc90d88b11e5b001de2f4665c1f2@ci.secure-pass.net
	// E2m6HawI743as61Kv0OhyPb6wAewXnwVkLLcF82rKOWe1SJ0Wd
	// https://beta.secure-pass.net
}

func ExampleSecurePass_Ping() {
	resp, err := testInstance.Ping()
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

func ExampleSecurePass_AppAdd() {
	var (
		resp         APIResponse
		addResponse  *AppAddResponse
		fixtureAppID string
	)

	app := fmt.Sprintf("test_fixture_%d_%d", time.Now().Unix(), rand.Int())
	// Create a new app
	addResponse, _ = testInstance.AppAdd(&ApplicationDescriptor{
		Label: app,
	})
	fixtureAppID = addResponse.AppID
	fmt.Println(addResponse.ErrorCode())
	fmt.Println(addResponse.ErrorMessage() == "")
	// Check for its existence
	resp, _ = testInstance.AppInfo(fixtureAppID)
	fmt.Println(resp.ErrorCode())
	// Remove it
	resp, _ = testInstance.AppDel(fixtureAppID)
	fmt.Println(resp.ErrorCode())
	// Check whether it does not longer exist
	resp, _ = testInstance.AppInfo(fixtureAppID)
	fmt.Println(resp.ErrorCode())
	// Output:
	// 0
	// true
	// 0
	// 0
	// 10
}
