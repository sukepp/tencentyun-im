package main

import (
	"tencentyun_im/timRestAPI"
)

func main() {
	// 1.Create a RestAPI
	timRestAPI := timRestAPI.CreateRestAPI()

	// 2.Import a account
	timRestAPI.AccountImport("user3", "test_sukepp", "www.baidu.com")

	// 3.Set a profile portrait
	//timRestAPI.ProfilePortraitSet("sukepp", "hanpeisong")

	// 4.Create a group
	//timRestAPI.GroupCreateGroup("Public", "myFirstGroup", "sukepp")

	// 5.Add a group member
	//timRestAPI.GroupAddGroupMember("@TGS#2RKJBP6ET", "user2", 1)

	// 6.Delete a group member
	//timRestAPI.GroupDeleteGroupMember("@TGS#2AQDDO6ED", "user2", 1)

	// 7.Destroy a group
	//timRestAPI.GroupDestroyGroup("@TGS#2AQDDO6ED")

	// 8.Send a group message
	//timRestAPI.GroupSendGroupMsg("user2", "@TGS#2RKJBP6ET", "Hello")

	// 9.Send a system notification
	//timRestAPI.GroupSendGroupSystemNotification("@TGS#2RKJBP6ET", "SUKEPP")
}
