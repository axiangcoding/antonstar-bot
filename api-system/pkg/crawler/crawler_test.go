package crawler

import (
	"fmt"
	table2 "github.com/axiangcoding/antonstar-bot/internal/data/table"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProfileFromWTOfficial(t *testing.T) {
	if err := GetProfileFromWTOfficial("OnTheRocks", func(status int, user *table2.GameUser) {
		if status != StatusFound {
			t.Failed()
		}
		assert.Equal(t, "OnTheRocks", user.Nick)
		assert.Equal(t, 100, user.Level)
	}); err != nil {
		t.Failed()
	}
}

func TestGetProfileFromThunderskill(t *testing.T) {
	if err := GetProfileFromThunderskill("OnTheRocks", func(status int, skill *ThunderSkillResp) {
		if status != StatusFound {
			t.Failed()
		}
		assert.Equal(t, 84.72, skill.Stats.R.Kpd)
		assert.Equal(t, 88.57, skill.Stats.S.Kpd)
		assert.Equal(t, 82.48, skill.Stats.A.Kpd)
	}); err != nil {
		t.Failed()
	}
}

func TestGetFirstPageNewsFromWTOfficial(t *testing.T) {
	if err := GetFirstPageNewsFromWTOfficial("en", func(news []table2.GameNew) {
		fmt.Println(len(news))
		for _, i := range news {
			fmt.Println(i)
		}
	}); err != nil {
		t.Failed()
	}
}
