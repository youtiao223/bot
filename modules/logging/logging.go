package logging

import (
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"qqBot/bot"
	"qqBot/utils"
)

var logger = utils.GetModuleLogger("internal.logging")

func logGroupMessage(msg *message.GroupMessage) {
	logger.
		WithField("from", "GroupMessage").
		WithField("MessageID", msg.Id).
		WithField("MessageIID", msg.InternalId).
		WithField("GroupCode", msg.GroupCode).
		WithField("SenderID", msg.Sender.Uin).
		Info(msg.ToString())
}

func logPrivateMessage(msg *message.PrivateMessage) {
	logger.
		WithField("from", "PrivateMessage").
		WithField("MessageID", msg.Id).
		WithField("MessageIID", msg.InternalId).
		WithField("SenderID", msg.Sender.Uin).
		WithField("Target", msg.Target).
		Info(msg.ToString())
}

func logFriendMessageRecallEvent(event *client.FriendMessageRecalledEvent) {
	logger.
		WithField("from", "FriendsMessageRecall").
		WithField("MessageID", event.MessageId).
		WithField("SenderID", event.FriendUin).
		Info("friend message recall")
}

func logGroupMessageRecallEvent(event *client.GroupMessageRecalledEvent) {
	logger.
		WithField("from", "GroupMessageRecall").
		WithField("MessageID", event.MessageId).
		WithField("GroupCode", event.GroupCode).
		WithField("SenderID", event.AuthorUin).
		WithField("OperatorID", event.OperatorUin).
		Info("group message recall")
}

func logGroupMuteEvent(event *client.GroupMuteEvent) {
	logger.
		WithField("from", "GroupMute").
		WithField("GroupCode", event.GroupCode).
		WithField("OperatorID", event.OperatorUin).
		WithField("TargetID", event.TargetUin).
		WithField("MuteTime", event.Time).
		Info("group mute")
}

func logDisconnect(event *client.ClientDisconnectedEvent) {
	logger.
		WithField("from", "Disconnected").
		WithField("reason", event.Message).
		Warn("bot disconnected")
}

func registerLog(b *bot.Bot) {
	b.OnGroupMessageRecalled(func(qqClient *client.QQClient, event *client.GroupMessageRecalledEvent) {
		logGroupMessageRecallEvent(event)
	})

	b.OnGroupMessage(func(qqClient *client.QQClient, groupMessage *message.GroupMessage) {
		logGroupMessage(groupMessage)
	})

	b.OnGroupMuted(func(qqClient *client.QQClient, event *client.GroupMuteEvent) {
		logGroupMuteEvent(event)
	})

	b.OnPrivateMessage(func(qqClient *client.QQClient, privateMessage *message.PrivateMessage) {
		logPrivateMessage(privateMessage)
	})

	b.OnFriendMessageRecalled(func(qqClient *client.QQClient, event *client.FriendMessageRecalledEvent) {
		logFriendMessageRecallEvent(event)
	})

	b.OnDisconnected(func(qqClient *client.QQClient, event *client.ClientDisconnectedEvent) {
		logDisconnect(event)
	})
}
