package bilibili

type RootResp struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Msg     string `json:"msg,omitempty"`
}

type RoomInfoResp struct {
	RootResp
	Data RoomInfoData `json:"data"`
}

type RoomInfoData struct {
	UId         int64  `json:"uid,omitempty"`
	RoomId      int64  `json:"room_id,omitempty"`
	Attention   int    `json:"attention,omitempty"`
	Online      int    `json:"online,omitempty"`
	Description string `json:"description,omitempty"`
	LiveStatus  int    `json:"live_status,omitempty"`
	Title       string `json:"title,omitempty"`
	LiveTime    string `json:"live_time,omitempty"`
}
