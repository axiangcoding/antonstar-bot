package service

import (
	"encoding/binary"
	"fmt"
	"golang.org/x/exp/rand"
	"time"
)

// DrawNumber 抽一个数字
func DrawNumber(id int) int32 {
	date := time.Now().Format("2006-01-02")
	sprintf := fmt.Sprintf("%d+%s", id, date)
	u := binary.BigEndian.Uint64([]byte(sprintf))
	return rand.New(rand.NewSource(u)).Int31n(100)
}
