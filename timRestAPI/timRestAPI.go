package timRestAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

const (
	sdkappid   = 1400048262
	identifier = "admin"
	usersig    = "eJxlj11LwzAYhe-7K0JvK5qkHxRhF3PMVo0Xad2Y3pTYpO5F*7E0rivD-*6sAwOe2*c5HM7RQQi5Tyy-FGXZfjamMGOnXHSNXOxe-MGuA1kIU-ha-oPq0IFWhaiM0hMkYRhSjG0HpGoMVHA2hKyhsXAv34tp47cfnMpBTCNqK-A2wcclX9ylPFpsdHWVb-eBZFmWP6hArEjcxnPZsdeP9bBpDymvPMrmcLNNa6JGM6xXCS13Owb3L2TwuP-MhefdJrzny3BMcJ9pPptZkwZqdT5EMI5iEoUW3SvdQ9tMAsUnhfr4J67z5XwD2mxccg__"
)

type TimRestAPI struct {
	sdkappid   int
	identifier string
	usersig    string
}

func CreateRestAPI() *TimRestAPI {
	return &TimRestAPI{
		sdkappid,
		identifier,
		usersig,
	}
}

func (timRestAPI *TimRestAPI) api(serviceName string, cmdName string, reqData []byte) string {
	urlPart := []string{"https://console.tim.qq.com/v4/", serviceName, "/", cmdName, "?usersig=",
		timRestAPI.usersig, "&identifier=", timRestAPI.identifier, "&sdkappid=", strconv.Itoa(timRestAPI.sdkappid),
		"&random=", strconv.Itoa(int(rand.Int31())), "&contenttype=json"}
	url := strings.Join(urlPart, "")

	bodyType := "application/json;charset=utf-8"

	req := bytes.NewBuffer(reqData)
	resp, _ := http.Post(url, bodyType, req)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	return string(body)
}

func (timRestAPI *TimRestAPI) AccountImport(identifier, nick, faceUrl string) {
	msg := struct{ Identifier, Nick, FaceUrl string }{identifier, nick, faceUrl}
	if reqData, err := json.Marshal(msg); err != nil {
		fmt.Println(err)
	} else {
		timRestAPI.api("im_open_login_svc", "account_import", reqData)
	}
}

func (timRestAPI *TimRestAPI) ProfilePortraitSet(accountId, newName string) {
	msg := struct {
		From_Account string
		ProfileItem  []struct{ Tag, Value string }
	}{accountId, []struct{ Tag, Value string }{{"Tag_Profile_IM_Nick", newName}}}

	if reqData, err := json.Marshal(msg); err != nil {
		fmt.Println(err)
	} else {
		timRestAPI.api("profile", "portrait_set", reqData)
	}
}

func (timRestAPI *TimRestAPI) GroupCreateGroup(groupType, groupName, ownerId string) {
	msg := struct{ Type, Name, Owner_Account string }{groupType, groupName, ownerId}
	if reqData, err := json.Marshal(msg); err != nil {
		fmt.Println(err)
	} else {
		timRestAPI.api("group_open_http_svc", "create_group", reqData)
	}
}

func (timRestAPI *TimRestAPI) GroupAddGroupMember(groupId, memberId string, silence int) {
	msg := struct {
		GroupId    string
		MemberList []struct{ Member_Account string }
		Silence    int
	}{groupId, []struct{ Member_Account string }{{memberId}}, silence}

	if reqData, err := json.Marshal(msg); err != nil {
		fmt.Println(err)
	} else {
		timRestAPI.api("group_open_http_svc", "add_group_member", reqData)
	}
}

func (timRestAPI *TimRestAPI) GroupDeleteGroupMember(groupId, memberId string, silence int) {
	msg := struct {
		GroupId             string
		MemberToDel_Account []string
		Silence             int
	}{groupId, []string{memberId}, silence}

	if reqData, err := json.Marshal(msg); err != nil {
		fmt.Println(err)
	} else {
		timRestAPI.api("group_open_http_svc", "delete_group_member", reqData)
	}
}

func (timRestAPI *TimRestAPI) GroupDestroyGroup(groupId string) {
	msg := struct{ GroupId string }{groupId}
	if reqData, err := json.Marshal(msg); err != nil {
		fmt.Println(err)
	} else {
		timRestAPI.api("group_open_http_svc", "destroy_group", reqData)
	}
}

func (timRestAPI *TimRestAPI) GroupSendGroupMsg(accountId, groupId, textContent string) {
	msg := struct {
		GroupId      string
		From_Account string
		Random       int32
		MsgBody      []struct {
			MsgType    string
			MsgContent struct{ Text string }
		}
	}{groupId, accountId, rand.Int31(), []struct {
		MsgType    string
		MsgContent struct{ Text string }
	}{{"TIMTextElem", struct{ Text string }{textContent}}}}

	if reqData, err := json.Marshal(msg); err != nil {
		fmt.Println(err)
	} else {
		timRestAPI.api("group_open_http_svc", "send_group_msg", reqData)
	}
}

func (timRestAPI *TimRestAPI) GroupSendGroupSystemNotification(groupId, textContent string) {
	msg := struct{ GroupId, Content string }{groupId, textContent}
	if reqData, err := json.Marshal(msg); err != nil {
		fmt.Println(err)
	} else {
		timRestAPI.api("group_open_http_svc", "send_group_system_notification", reqData)
	}
}
