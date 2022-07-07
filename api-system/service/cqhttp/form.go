package cqhttp

var (
	PostTypeMessage   = "message"
	PostTypeRequest   = "request"
	PostTypeNotice    = "notice"
	PostTypeMetaEvent = "meta_event"
)

type CommonReportForm struct {
	Time     int64  `json:"time,omitempty"`
	SelfId   int64  `json:"self_id,omitempty"`
	PostType string `json:"post_type,omitempty"`
}

type MessageReportForm struct {
	CommonReportForm
	SubType   string `json:"sub_type,omitempty"`
	MessageId int32  `json:"message_id,omitempty"`
	UserId    int64  `json:"user_id,omitempty"`
	MessageForm
	RawMessage string        `json:"raw_message,omitempty"`
	Font       int           `json:"font,omitempty"`
	Sender     MessageSender `json:"sender"`
}

type MetaEventReportForm struct {
	CommonReportForm
	MetaEventType string `json:"meta_event_type,omitempty"`
}

type MessageForm struct {
}

type MessageSender struct {
	UserId   string `json:"user_id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Sex      string `json:"sex,omitempty"`
	Age      int32  `json:"age,omitempty"`
}

type MessageSenderFromGroup struct {
	MessageSender
	Card  string `json:"card,omitempty"`
	Area  string `json:"area,omitempty"`
	Level string `json:"level,omitempty"`
	Role  string `json:"role,omitempty"`
	Title string `json:"title,omitempty"`
}
