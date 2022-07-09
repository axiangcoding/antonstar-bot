package cqhttp

import "encoding/json"

var (
	PostTypeMessage    = "message"
	PostTypeRequest    = "request"
	PostTypeNotice     = "notice"
	PostTypeMetaEvent  = "meta_event"
	EventTypeHeartBeat = "heartbeat"
)

type MetaTypeHeartBeatEventMessage struct {
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

func (m *MetaTypeHeartBeatEventMessage) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

func (m *MetaTypeHeartBeatEventMessage) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &m)
}
