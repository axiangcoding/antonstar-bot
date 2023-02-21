package table

import (
	"github.com/axiangcoding/antonstar-bot/internal/data/display"
	"gorm.io/gorm"
)

type GameNew struct {
	gorm.Model
	Link      string `gorm:"uniqueIndex;size:255"`
	PosterUrl string
	Title     string
	Comment   string
	DateStr   string
}

func (n GameNew) ToDisplayGameUser() display.GameNew {
	return display.GameNew{
		CreatedAt: n.CreatedAt,
		UpdatedAt: n.UpdatedAt,
		Link:      n.Link,
		PosterUrl: n.PosterUrl,
		Title:     n.Title,
		Comment:   n.Comment,
		DateStr:   n.DateStr,
	}
}
