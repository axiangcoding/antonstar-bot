package service

import (
	"fmt"
	"golang.org/x/exp/rand"
	"hash/crc32"
	"time"
)

// DrawNumber 抽一个数字
func DrawNumber(id int, now time.Time) int32 {
	date := now.Format("2006-01-02")
	sprintf := fmt.Sprintf("%d+%s", id, date)
	hash := crc32.ChecksumIEEE([]byte(sprintf))
	return rand.New(rand.NewSource(uint64(hash))).Int31n(100)
}
