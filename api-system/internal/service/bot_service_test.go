package service

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

var tests = []struct {
	id     int64
	time   string
	number int32
}{
	{
		id:     123456789,
		time:   "2022-07-20",
		number: 70,
	}, {
		id:     123456789,
		time:   "2022-07-21",
		number: 32,
	},
	{
		id:     123456789,
		time:   "2022-07-22",
		number: 26,
	},
}

func TestDrawNumber(t *testing.T) {
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			now, _ := time.Parse("2006-01-02", tt.time)
			num := DrawNumber(tt.id, now)
			assert.Equal(t, tt.number, num)
		})
	}
}
