package main

import (
	"im/timRestAPI"
)

func main() {

	timRestAPI := timRestAPI.CreateRestAPI()
	timRestAPI.Account_import("sukepp", "test_sukepp", "www.baidu.com")
	timRestAPI.Profile_portrait_set("sukepp", "hanpeisong")
	timRestAPI.Group_create_group("Public", "myFirstGroup", "sukepp")
}
