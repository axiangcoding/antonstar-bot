package crawler

import (
	"github.com/axiangcoding/antonstar-bot/data/table"
	"github.com/axiangcoding/antonstar-bot/logging"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetProfileFromWTOfficial(t *testing.T) {
	if err := GetProfileFromWTOfficial("OnTheRocks", func(status int, user *table.GameUser) {
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

func TestMain(m *testing.M) {
	logging.InitLogger()
	code := m.Run()
	os.Exit(code)
}
