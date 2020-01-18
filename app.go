// 空插件
package main

import (
	"github.com/Tnze/CoolQ-Golang-SDK/v2/cqp"
	log "github.com/sirupsen/logrus"
)

func main() { cqp.Main() }

func init() {
	// AppID 需要修改为你的插件的AppID
	cqp.AppID = "your.app.id"
	cqp.PrivateMsg = onPrivateMsg
}

// TODO: 恢复空插件
func onPrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {
	//cqp.SendPrivateMsg(fromQQ, msg) //复读机
	log.WithFields(log.Fields{
		"subType": subType,
		"msgID":   msgID,
		"fromQQ":  fromQQ,
		"msg":     msg,
		"font":    font,
	})
	return 0
}
