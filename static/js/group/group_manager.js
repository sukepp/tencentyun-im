//从我的群组列表中发群消息

function sendGroupMsg() {
    userID = document.getElementById("sgm_to_userid").value;
    userName = document.getElementById("sgm_to_user_name").value;
    toAccount = document.getElementById("sgm_to_groupid").value;
    toName = document.getElementById("sgm_to_group_name").value;
    msgtosend = document.getElementById("sgm_content").value;
    var groupHeadUrl = 'img/group.jpg'; 

    var msgLen = webim.Tool.getStrBytes(msgtosend);

    var maxLen, errInfo;

    maxLen = webim.MSG_MAX_LENGTH.GROUP;
    errInfo = "消息长度超出限制(最多" + Math.round(maxLen / 3) + "汉字)";

    if (msgtosend.length < 1) {
        alert("发送的消息不能为空!");
        return;
    }

    if (msgLen > maxLen) {
        alert(errInfo);
        return;
    }

    var sess = webim.MsgStore.sessByTypeId(webim.SESSION_TYPE.GROUP, toAccount);
    if (!sess) {
        sess = new webim.Session(webim.SESSION_TYPE.GROUP, toAccount, toName, groupHeadUrl, Math.round(new Date().getTime() / 1000));
    }
    var isSend = true; //是否为自己发送
    var seq = -1; //消息序列，-1表示sdk自动生成，用于去重
    var random = Math.round(Math.random() * 4294967296); //消息随机数，用于去重
    var msgTime = Math.round(new Date().getTime() / 1000); //消息时间戳
    var subType; //消息子类型

    subType = webim.GROUP_MSG_SUB_TYPE.COMMON;

    //var msg = new webim.Msg(sess, isSend, seq, random, msgTime, loginInfo.identifier, subType, loginInfo.identifierNick);
    var msg = new webim.Msg(sess, isSend, seq, random, msgTime, userID, subType, userName);

    var text_obj;

    text_obj = new webim.Msg.Elem.Text(msgtosend);
    msg.addText(text_obj);

    webim.sendMsg(msg, function(resp) {
        // $("#sgm_content").val('');
    }, function(err) {
        alert(err.ErrorInfo);
        // $("#sgm_content").val('');
    });
}
