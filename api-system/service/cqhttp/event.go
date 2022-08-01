package cqhttp

import "encoding/json"

var (
	PostTypeMessage    = "message"
	PostTypeRequest    = "request"
	PostTypeNotice     = "notice"
	PostTypeMetaEvent  = "meta_event"
	EventTypeHeartBeat = "heartbeat"
	MessageTypeGroup   = "group"
)

// MetaTypeHeartBeatEvent 心跳事件
type MetaTypeHeartBeatEvent struct {
	Interval      int    `json:"interval" mapstructure:"interval"`
	MetaEventType string `json:"meta_event_type" mapstructure:"meta_event_type"`
	PostType      string `json:"post_type" mapstructure:"post_type"`
	SelfId        int64  `json:"self_id" mapstructure:"self_id"`
	Status        struct {
		AppEnabled     bool        `json:"app_enabled" mapstructure:"app_enabled"`
		AppGood        bool        `json:"app_good" mapstructure:"app_good"`
		AppInitialized bool        `json:"app_initialized" mapstructure:"app_initialized"`
		Good           bool        `json:"good" mapstructure:"good"`
		Online         bool        `json:"online" mapstructure:"online"`
		PluginsGood    interface{} `json:"plugins_good" mapstructure:"plugins_good"`
		Stat           struct {
			DisconnectTimes int `json:"disconnect_times" mapstructure:"disconnect_times"`
			LastMessageTime int `json:"last_message_time" mapstructure:"last_message_time"`
			LostTimes       int `json:"lost_times" mapstructure:"lost_times"`
			MessageReceived int `json:"message_received" mapstructure:"message_received"`
			MessageSent     int `json:"message_sent" mapstructure:"message_sent"`
			PacketLost      int `json:"packet_lost" mapstructure:"packet_lost"`
			PacketReceived  int `json:"packet_received" mapstructure:"packet_received"`
			PacketSent      int `json:"packet_sent" mapstructure:"packet_sent"`
		} `json:"stat" mapstructure:"stat"`
	} `json:"status" mapstructure:"status"`
	Time int `json:"time" mapstructure:"time"`
}

func (m *MetaTypeHeartBeatEvent) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

func (m *MetaTypeHeartBeatEvent) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &m)
}

// MessageGroupEvent 群聊事件
type MessageGroupEvent struct {
	Anonymous   interface{} `json:"anonymous" mapstructure:"anonymous"`
	Font        int         `json:"font" mapstructure:"font"`
	GroupId     int64       `json:"group_id" mapstructure:"group_id"`
	Message     string      `json:"message" mapstructure:"message"`
	MessageId   int         `json:"message_id" mapstructure:"message_id"`
	MessageSeq  int         `json:"message_seq" mapstructure:"message_seq"`
	MessageType string      `json:"message_type" mapstructure:"message_type"`
	PostType    string      `json:"post_type" mapstructure:"post_type"`
	RawMessage  string      `json:"raw_message" mapstructure:"raw_message"`
	SelfId      int64       `json:"self_id" mapstructure:"self_id"`
	Sender      struct {
		Age      int    `json:"age" mapstructure:"age"`
		Area     string `json:"area" mapstructure:"area"`
		Card     string `json:"card" mapstructure:"card"`
		Level    string `json:"level" mapstructure:"level"`
		Nickname string `json:"nickname" mapstructure:"nickname"`
		Role     string `json:"role" mapstructure:"role"`
		Sex      string `json:"sex" mapstructure:"sex"`
		Title    string `json:"title" mapstructure:"title"`
		UserId   int64  `json:"user_id" mapstructure:"user_id"`
	} `json:"sender" mapstructure:"sender"`
	SubType string `json:"sub_type" mapstructure:"sub_type"`
	Time    int    `json:"time" mapstructure:"time"`
	UserId  int64  `json:"user_id" mapstructure:"user_id"`
}

func (m *MessageGroupEvent) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

func (m *MessageGroupEvent) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &m)
}
