// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"time"
)

type Friend struct {
	// 主键Id
	ID int64 `json:"id"`
	// 用户Id
	UserID int32 `json:"user_id"`
	// 好友Id
	FriendID int32 `json:"friend_id"`
	// 好友备注
	Note string `json:"note"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
}

type Group struct {
	// 主键Id
	ID int64 `json:"id"`
	// 创建者Id
	CreatorID int32 `json:"creator_id"`
	// 群组名
	GroupName string `json:"group_name"`
	// 群员人数
	CurrentQuantity int32 `json:"current_quantity"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
}

type GroupMemberShip struct {
	// 群组成员的唯一标识符
	ID int64 `json:"id"`
	// 用户Id
	UserID int32 `json:"user_id"`
	// 群组Id
	GroupID int32 `json:"group_id"`
	// 角色, 0: 管理员, 1: 普通成员
	Role int32 `json:"role"`
	// 加入时间
	JoinedAt time.Time `json:"joined_at"`
}

type Message struct {
	// 主键Id
	ID int64 `json:"id"`
	// 消息发送者Id
	SenderID int32 `json:"sender_id"`
	// 消息接收者Id
	ReceiverID int32 `json:"receiver_id"`
	// 消息类型, 1: 文字, 2: 图片, 3: 文件, 4: 语音, 5: 视频, 6: 语音通话, 7: 视频通话
	MessageType int16 `json:"message_type"`
	// 消息内容
	Content string `json:"content"`
	// 发送类型, 1: 私聊, 2: 群聊
	SendType int16 `json:"send_type"`
	// 消息状态, 1: 发送失败, 2: 发送成功, 3: 已读, 4: 未读
	MessageStatus int16 `json:"message_status"`
	// 消息发送时间
	SendingTime time.Time `json:"sending_time"`
	// 消息读取时间
	ReceivTime time.Time `json:"receiv_time"`
}

type User struct {
	// 用户Id
	ID int32 `json:"id"`
	// 用户名
	Username string `json:"username"`
	// 用户昵称
	Nickname string `json:"nickname"`
	// 密码
	Password string `json:"password"`
	// 邮箱
	Email string `json:"email"`
	// 性别, 1 男, 2 女, 3 未知
	Gender int16 `json:"gender"`
	// 头像图片路径或链接
	ProfilePictureUrl string `json:"profile_picture_url"`
	// 在线状态(在线/离线)
	OnlineStatus bool `json:"online_status"`
	// 密码更新时间
	PasswordChangedAt time.Time `json:"password_changed_at"`
	// 最后在线时间
	LastLoginAt time.Time `json:"last_login_at"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 更新时间
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLoginLog struct {
	// 日志Id
	ID int64 `json:"id"`
	// 用户id
	UserID int32 `json:"user_id"`
	// 登录时间
	LoginTime time.Time `json:"login_time"`
	// 登录ip
	LoginIp string `json:"login_ip"`
	// 登录ip所在地
	LoginIpRegion string `json:"login_ip_region"`
	// 是否登录异常, t: 异常, f: 非异常
	IsLoginExceptional bool `json:"is_login_exceptional"`
	// 登录平台
	Platform string `json:"platform"`
	// user_agent
	UserAgent string `json:"user_agent"`
}
