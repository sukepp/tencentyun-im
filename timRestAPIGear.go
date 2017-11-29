package main

import (
	"fmt"
	"tencentyun_im/timRestAPI"
)

/*
type GroupApply struct {
	CallbackCommand   string
	GroupId           string
	Type              string
	Requestor_Account string
}
*/

type GroupIdList []struct {
	GroupId string
	Name    string
}

func main() {
	// 1.Create a RestAPI
	timRestAPI := timRestAPI.CreateRestAPI()

	// 2.Import a account
	//timRestAPI.AccountImport("user3", "test_sukepp", "www.baidu.com")

	// 3.Set a profile portrait
	//timRestAPI.ProfilePortraitSet("sukepp", "hanpeisong")

	// 4.Create a group
	//timRestAPI.GroupCreateGroup("Public", "myThirdGroup", "sukepp")

	// 5.Add a group member
	//timRestAPI.GroupAddGroupMember("@TGS#2RKJBP6ET", "icecut", 1)

	// 6.Delete a group member
	//timRestAPI.GroupDeleteGroupMember("@TGS#2NKSFY6EG", "bob", 1)

	// 7.Destroy a group
	//timRestAPI.GroupDestroyGroup("@TGS#2AQDDO6ED")

	// 8.Send a group message
	//timRestAPI.GroupSendGroupMsg("sukepp", "@TGS#2RKJBP6ET", "Hello, I am sukepp")

	// 9.Send a system notification
	//timRestAPI.GroupSendGroupSystemNotification("@TGS#2RKJBP6ET", "SUKEPP")

	// 10.Get a joined group list
	// if groups := timRestAPI.GroupGetJoinedGroupList("sukepp"); groups != nil {
	// 	for _, group := range groups {
	// 		fmt.Println(group.GroupId + ":" + group.Name)
	// 	}
	// }

	// 11.Get a all-groups List
	groups := timRestAPI.GroupGetAppidGroupList(50)
	for _, group := range groups {
		fmt.Println(group)
	}

}
