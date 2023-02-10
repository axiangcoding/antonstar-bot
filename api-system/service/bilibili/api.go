package bilibili

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"strconv"
	"time"
)

func GetBiliBiliRoomInfo(roomId int64) (*RoomInfoResp, error) {
	client := resty.New().SetTimeout(time.Second * 10)
	var roomInfo RoomInfoResp
	url := "https://api.live.bilibili.com/room/v1/Room/get_info"
	resp, err := client.R().SetQueryParam("room_id", strconv.FormatInt(roomId, 10)).
		SetResult(&roomInfo).
		Get(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New("response status code error")
	}
	return &roomInfo, err
}
