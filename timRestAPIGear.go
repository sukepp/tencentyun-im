package main

import (
	"im/timRestAPI"
)

func main() {
	// 1.Create a RestAPI
	timRestAPI := timRestAPI.CreateRestAPI()

	// 2.Import a account
	timRestAPI.Account_import("user3", "test_sukepp", "www.baidu.com")

	// 3.Set a profile portrait
	//timRestAPI.Profile_portrait_set("sukepp", "hanpeisong")

	// 4.Create a group
	//timRestAPI.Group_create_group("Public", "myFirstGroup", "sukepp")

	// 5.Add a group member
	//timRestAPI.Group_add_group_member("@TGS#2RKJBP6ET", "user2", 1)

	// 6.Delete a group member
	//timRestAPI.Group_delete_group_member("@TGS#2AQDDO6ED", "user2", 1)

	// 7.Destroy a group
	//timRestAPI.Group_destroy_group("@TGS#2AQDDO6ED")

	// 8.Send a group message
	//timRestAPI.Group_send_group_msg("user2", "@TGS#2RKJBP6ET", "Hello")

	// 9.Send a system notification
	//timRestAPI.Group_send_group_system_notification("@TGS#2RKJBP6ET", "SUKEPP")
}
