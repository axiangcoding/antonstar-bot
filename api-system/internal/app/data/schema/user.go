package schema

import (
	"axiangcoding/antonstar/api-system/pkg/logging"
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
)

const (
	UserRoleAdmin   = "admin"
	UserRoleManager = "manager"
	UserRoleNormal  = "normal"
)

type User struct {
	gorm.Model
	UserId int64 `gorm:"uniqueIndex"`
	//下面三种信息在库中都是不可重复的
	UserName string  `gorm:"uniqueIndex;size:255"`
	Email    *string `gorm:"uniqueIndex;size:255"`
	Phone    *string `gorm:"uniqueIndex;size:255"`
	Password string
	Roles    string
}

func (u *User) GenerateId() {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		logging.Error("generate snowflake id error", err)
		return
	}
	u.UserId = node.Generate().Int64()
}
