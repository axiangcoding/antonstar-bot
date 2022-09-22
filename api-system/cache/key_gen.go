package cache

import "fmt"

const (
	CqHttpPrefix         = "CQHTTP"
	GameUserPrefix       = "GameUser"
	BiliRoomLivingPrefix = "BiliRoom"
)

func GenerateCQHTTPCacheKey(postType string, eventType string, selfId int64) string {
	return fmt.Sprintf("%s:%s;%s;%d", CqHttpPrefix, postType, eventType, selfId)
}

func GenerateGameUserCacheKey(nickname string) string {
	return fmt.Sprintf("%s:%s", GameUserPrefix, nickname)
}

func GenerateBiliRoomLivingCacheKey(groupId, roomId int64) string {
	return fmt.Sprintf("%s:%d;%d", BiliRoomLivingPrefix, groupId, roomId)

}
